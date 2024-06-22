package sample.testcontainer.kanban

import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.test.context.TestConfiguration
import org.springframework.boot.testcontainers.service.connection.ServiceConnection
import org.springframework.context.annotation.Bean
import org.springframework.test.context.ActiveProfiles
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
                "POSTGRES_DB" to "kanbandb",
                "POSTGRES_USER" to username,
                "POSTGRES_PASSWORD" to password
            )
        ).withInitScript("./initial-state.sql")
    }

}
