package sample.testcontainer.kanban

import org.springframework.boot.fromApplication
import org.springframework.boot.with
import org.springframework.core.env.AbstractEnvironment

/**
 * this little helper is generated by start.spring.io when you select the
 * testcontainers starter. we just tweaked it a little since we fuzz with the
 * credentials when running under test profile.
 */
fun main(args: Array<String>) {
	System.setProperty(AbstractEnvironment.ACTIVE_PROFILES_PROPERTY_NAME, "test");
	fromApplication<SampleKanbanJvmApplication>()
		.with(TestcontainersConfiguration::class)
		.run(*args)
}
