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

import org.springframework.boot.test.context.TestConfiguration
import org.springframework.boot.testcontainers.service.connection.ServiceConnection
import org.springframework.context.annotation.Bean
import org.testcontainers.containers.PostgreSQLContainer
import org.testcontainers.utility.DockerImageName

@TestConfiguration(proxyBeanMethods = false)
class TestcontainersConfiguration {

    @Bean
    @ServiceConnection
    fun postgresContainer(): PostgreSQLContainer<*> {
        return PostgreSQLContainer(
            DockerImageName
                .parse("postgres:16.3-alpine3.20")
        ).withEnv(
            mapOf(
                "POSTGRES_DB" to "kanbandb",
                "POSTGRES_USER" to "kanbanusr",
                "POSTGRES_PASSWORD" to "kanbanpwd"
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

Ava has hooks where you can properly setup and teardown the database and then
update database configuration accordingly. Mind to write proper testable code.

_some sample code_

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
