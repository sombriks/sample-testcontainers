package sample.testcontainer.kanban

import org.springframework.boot.fromApplication
import org.springframework.boot.with
import org.springframework.core.env.AbstractEnvironment

fun main(args: Array<String>) {
	System.setProperty(AbstractEnvironment.ACTIVE_PROFILES_PROPERTY_NAME, "test");
	fromApplication<SampleKanbanJvmApplication>()
		.with(TestcontainersConfiguration::class)
		.run(*args)
}
