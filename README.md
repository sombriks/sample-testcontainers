# [Sample Testcontainers][repo]

Samples on why and how to use [TestContainers][testcontainers]

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
// app/services/services.spec.js
import {resolve} from 'node:path';
import test from 'ava';
import {PostgreSqlContainer} from '@testcontainers/postgresql';
import {prepareDatabase} from '../configs/database.js';
import {boardServices} from './board-services.js';

test.before(async t => {
  // Testcontainer setup
  t.context.postgres = await new PostgreSqlContainer('postgres:16.3-alpine3.20')
    .withDatabase(process.env.PG_DATABASE)
    .withUsername(process.env.PG_USERNAME)
    .withPassword(process.env.PG_PASSWORD)
    .withBindMounts([{
      source: resolve(process.env.PG_INIT_SCRIPT),
      target: '/docker-entrypoint-initdb.d/init.sql',
    }])
    .start();

  // Application setup properly tailored for tests
  const database = prepareDatabase(t.context.postgres.getConnectionUri());
  const service = boardServices({db: database});

  // Context register for proper teardown
  t.context.db = database;
  t.context.service = service;
});

test.after.always(async t => {
  // teardown 
  await t.context.db.destroy();
  await t.context.postgres.stop({timeout: 500});
});

test('should list people', async t => {
  const people = await t.context.service.listUsers();
  t.is(people.length, 5); // we know there are 5 people
});

test('should list tasks', async t => {
  const tasks = await t.context.service.listTasks();
  t.is(tasks.length, 5);
});
```

Mind to write proper testable code: it's very tempting to just create and export
your objects directly from modules:

```javascript
import Pug from 'koa-pug';

export const pug = new Pug({
	viewPath: 'app/templates', // TODO use import.meta.url thing
});
```

But for proper testing you must provide inversion of control, dependency
injection (the **D** in *SOLID*):

```javascript
import Knex from 'knex'

// we can get a Knex instance either pointing to postgres from env or provide a
// custom connection string.
export const prepareDatabase = (connection = process.env.PG_CONNECTION_URL) => Knex({
	client: 'pg',
	connection,
})
```

Besides that implementation detail, everything else should work under test the
same way it works during development or in production. sampe code, no mocks,
same database engine, same dialect.

### Sample code - Echo/Goqu/Testify

Testify offers setup hooks where you can provision and later release the
database runtime.

_some sample code_

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
