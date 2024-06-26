# [Sample Testcontainers][repo]

Samples on why and how to use [TestContainers][testcontainers]

[![GO CI](https://github.com/sombriks/sample-testcontainers/actions/workflows/go.yml/badge.svg)](https://github.com/sombriks/sample-testcontainers/actions/workflows/go.yml)
[![JVM CI](https://github.com/sombriks/sample-testcontainers/actions/workflows/jvm.yml/badge.svg)](https://github.com/sombriks/sample-testcontainers/actions/workflows/jvm.yml)
[![Node CI](https://github.com/sombriks/sample-testcontainers/actions/workflows/node.yml/badge.svg)](https://github.com/sombriks/sample-testcontainers/actions/workflows/node.yml)

<video width="320" height="240" controls>
  <source src="./docs/sample-kanban-2024-06-25-11-16-07.mp4" type="video/mp4">
</video>

## Test Boundaries

Untested code is a dark jungle filled with unknown bugs. 

We write tests to light up a fire to keep unexpected problems away.

But how far should a test suite should go?

It's clear that any business-specific code must be covered with tests, but does
a 3rd party API endpoint should be tested too? And the database?

There are frontiers. Anything out of our control can not be properly tested.

And this is the crossroads: expand our control or mock boundaries.

## The problem with too much mocks

Don't get me wrong, mocks at the boundaries works. But as advised by Mockito
front page project, _don't mock everything_.

For example, this mock looks perfectly reasonable:

```kotlin
// mock to list data - ok
@BeforeEach
fun setup() {
    _when(
        personRepository.findByNameContainingIgnoreCase(
            anyString(), anyOrNull()
        )
    ).thenReturn(personPage)
}

@Test
fun `should list people`() {
    val result = boardService.listPeople("", pageable)
    assertThat(result, notNullValue())
}
```

But then:

```kotlin
// mock to insert - fail
@Test
@Disabled("We can keep mocking but we don't trust the test anymore")
fun `should save people`() {
    val person = Person(name = "Ferdinando")
    boardService.savePerson(person)
    assertThat(person.id, notNullValue()) //new person should have an id now
}
```

In this situation you can simply keep growing the mock surface but there will be
a point when you will be testing nothing at all.

To really solve it, your boundaries must expand. And if the boundary to expand
is the database, here goes some samples.

## Introducing TestContainers

One way to test the database is to use some lightweight database runtime like h2
or sqlite, but that comes with a price: the dialect might be different from the
real deal and therefore you must be cautious about your queries.

To properly avoid that, it's ideal to use same RDBMS for development, staging
and for testing.

Using TestContainers makes this task a real easy breeze.

## Testing the database

Whenever we need to "test the database", what we're really testing is a known
database state. We expect a certain user/password to be accepted; we expect a
certain schema and a set of tables to exists. We expect some data to be present.

Therefore, when spinning up a test suite involving relational data, some setup
is needed. And TestContainers offers goodies to be used exactly in that phase.

### Sample code - Spring/Kotlin/JUnit

Spring tests has not only the setup phase but also The @TestConfiguration
stereotype, so the DI container does all the heavy-lifting for you:

```kotlin
package sample.testcontainer.kanban

import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.test.context.TestConfiguration
import org.springframework.boot.testcontainers.service.connection.ServiceConnection
import org.springframework.context.annotation.Bean
import org.testcontainers.containers.PostgreSQLContainer
import org.testcontainers.utility.DockerImageName

@TestConfiguration(proxyBeanMethods = false)
class TestcontainersConfiguration {

    @Value("\${database}")
    private lateinit var database: String
    @Value("\${spring.datasource.username}")
    private lateinit var username: String
    @Value("\${spring.datasource.password}")
    private lateinit var password: String

    @Bean
    @ServiceConnection
    fun postgresContainer(): PostgreSQLContainer<*> {
        return PostgreSQLContainer(
            DockerImageName
                .parse("postgres:16.3-alpine3.20")
        ).withEnv(
            mapOf(
                "POSTGRES_DB" to database,
                "POSTGRES_USER" to username,
                "POSTGRES_PASSWORD" to password
            )
        ).withInitScript("./initial-state.sql")
    }

}

```

This configuration should be "imported" into the test case so the default
database configuration, which probably won't be present in a CI workflow, can be
replaced in a transparent way. Someone at TestContainers team indeed made a fine
work on this craft:

```kotlin
package sample.testcontainer.kanban

import org.junit.jupiter.api.Test
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.Import

@SpringBootTest
// just add that and you have a full-featured, predictable, database for test!
@Import(TestcontainersConfiguration::class)
class SampleKanbanJvmApplicationTests {

	@Test
	fun contextLoads() {
	}

}
```

### Sample code - Koa/Knex/Ava

Ava has hooks where you can properly set up and tear down the database. Update
[database configuration][node-tc] accordingly:

```javascript
// in app/configs/hook-test-container.js
import {resolve} from 'node:path';
import {PostgreSqlContainer} from '@testcontainers/postgresql';

/**
 * Helper to provision a postgresql for testing purposes
 *
 * @returns {Promise<StartedPostgreSqlContainer>} database container
 */
export const preparePostgres = async () => new PostgreSqlContainer('postgres:16.3-alpine3.20')
  .withDatabase(process.env.PG_DATABASE)
  .withUsername(process.env.PG_USERNAME)
  .withPassword(process.env.PG_PASSWORD)
  .withBindMounts([{
    source: resolve(process.env.PG_INIT_SCRIPT),
    target: '/docker-entrypoint-initdb.d/init.sql',
  }])
  .start();
```

A quick note, but the node postgresql container has a distinct idiom for the
initial script when compared with jvm or golang versions. Those have a
`withInitScript` builder call, while node version offer a more generic
`withBindMounts` call.

You then integrate the test container provisioning into your ava test like this:

```javascript
// in app/app.spec.js
import request from 'supertest';
import test from 'ava';
import {prepareApp} from './main.js';
import {prepareDatabase} from './configs/database.js';
import {boardServices} from './services/board-services.js';
import {boardRoutes} from './routes/board-routes.js';
import {preparePostgres} from './configs/hook-test-container.js';

test.before(async t => {
	// TestContainer setup
	t.context.postgres = await preparePostgres();

	// Application setup properly tailored for tests
	const database = prepareDatabase(t.context.postgres.getConnectionUri());
	const service = boardServices({db: database});
	const controller = boardRoutes({service});

	const {app} = prepareApp({db: database, service, controller});

	// Context registering for proper teardown
	t.context.db = database;
	t.context.app = app;
});

test.after.always(async t => {
	await t.context.db.destroy();
	await t.context.postgres.stop({timeout: 500});
});

test('app should be ok', async t => {
	const result = await request(t.context.app.callback()).get('/');
	t.is(result.status, 302);
	t.is(result.headers.location, '/board');
});

test('db should be ok', async t => {
	const {rows: [{result}]} = await t.context.db.raw('SELECT 1 + 1 as result');
	t.truthy(result);
	t.is(result, 2);
});

test('should serve login and have users', async t => {
	const result = await request(t.context.app.callback()).get('/login');
	t.is(result.status, 200);
	t.regex(result.text, /Alice|Bob|Caesar|Davide|Edward/);
});
```

Mind to write proper testable code: it's very tempting to just create and export
your objects directly from modules:

```javascript
// in app/configs/views.js
import {resolve} from 'node:path';
import Pug from 'koa-pug';

export const pug = new Pug({
  viewPath: resolve('./app/templates'),
});
```

It's pretty fine most of the time, templates directory isn't likely to become a
configurable thing, so it's ok.

But for proper testing you must provide inversion of control, dependency
inversion, the **D** in *[SOLID][solid]*:

```javascript
// in app/configs/database.js
import Knex from 'knex';

export const prepareDatabase = (connection = process.env.PG_CONNECTION_URL) => Knex({
  client: 'pg',
  connection,
});
```

The `prepareDatabase` call let us send any connection string we want for the
database, quite useful when we are spinning up a postgres container, but if none
is provided it will rely on what we have configured in the environment under the
`PG_CONNECTION_URL` variable.

Besides that implementation detail, everything else should work under test the
same way it works during development or in production. same code, no mocks, same
database engine, same dialect, same thing.

### Sample code - Echo/Goqu/Testify

[Testify][testify] offers setup hooks where you can provision and later release the
database runtime.

```go
package services

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/joho/godotenv"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/configs"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

type ServiceTestSuit struct {
	suite.Suite
	ctx     context.Context
	tc      *postgres.PostgresContainer
	db      *goqu.Database
	service *BoardService
}

// TestRunSuite when writing suites this is needed as a 'suite entrypoint'
// see https://pkg.go.dev/github.com/stretchr/testify/suite
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuit))
}

func (s *ServiceTestSuit) SetupSuite() {
	var err error
	// Test execution point is inside the package, not in project root
	_ = godotenv.Load("../../.env")

	s.ctx = context.Background()

	props, err := configs.NewDbProps()
	if err != nil {
		s.Fail("Suite setup failed", err)
	}
	s.tc, err = postgres.RunContainer(s.ctx,
		testcontainers.WithImage("postgres:16.3-alpine3.20"),
		postgres.WithInitScripts(fmt.Sprint("../../", props.InitScript)), // path changes due test entrypoint
		postgres.WithUsername(props.Username),
		postgres.WithDatabase(props.Database),
		postgres.WithPassword(props.Password),
		testcontainers.WithWaitStrategy(wait.
			ForLog("database system is ready to accept connections").
			WithOccurrence(2).
			WithStartupTimeout(10*time.Second)))
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	dsn, err := s.tc.ConnectionString(s.ctx, fmt.Sprint("sslmode=", props.SslMode))
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	s.db, err = configs.NewGoquDb(nil, &dsn)
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	s.service, err = NewBoardService(s.db)
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

}

func (s *ServiceTestSuit) TearDownSuite() {
	err := s.tc.Terminate(s.ctx)
	if err != nil {
		s.Fail("Suite tear down failed", err)
	}
}

// the test cases
```

Similar to the advice givn on node version, mind the configuration phase! your
code is supposed to offer reasonable defaults and proper dependency injection so
you can provide test values or production values whenever needed:

```go
package configs

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"log"
)

// NewGoquDb - provision a query builder instance
func NewGoquDb(d *DbProps, dsn *string) (*goqu.Database, error) {
	var err error
	if d == nil {
		log.Println("[WARN] db props missing, creating a default one...")
		d, err = NewDbProps()
	}

	// configure the query builder
	if dsn == nil {
		newDsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=%s", //
			d.Username, d.Password, d.Hostname, d.Database, d.SslMode)
		dsn = &newDsn
	} else {
		log.Printf("[INFO] using provided dsn [%s]\n", *dsn)
	}
	con, err := sql.Open("postgres", *dsn)
	if err != nil {
		return nil, err
	}
	// https://doug-martin.github.io/goqu/docs/selecting.html#scan-struct
	goqu.SetIgnoreUntaggedFields(true)
	db := goqu.New("postgres", con)
	db.Logger(log.Default())

	return db, nil
}
```

The sample above is called during configuration phase to provision the query
builder instance; it receives, however, optional parameters that allow us to set
appropriate values for development, test or production.

## CI/CD integration

Now the best part: most CI/CD infrastructure available out there will offer
docker runtimes, so your tests will run smoothly.

_some sample code_

## Conclusion

Now that your boundaries got extended, your confidence on the code grows more
and more. It does what it's supposed to do. It saves and list the expected
content. It works*. As far as the tests can tell.

The complete source code can be found here.

Happy hacking!

[repo]: https://github.com/sombriks/sample-testcontainers
[testcontainers]: https://testcontainers.com/
[node-tc]: https://testcontainers.com/guides/getting-started-with-testcontainers-for-nodejs/
[solid]: https://en.wikipedia.org/wiki/Dependency_inversion_principle
[testify]: https://github.com/stretchr/testify
