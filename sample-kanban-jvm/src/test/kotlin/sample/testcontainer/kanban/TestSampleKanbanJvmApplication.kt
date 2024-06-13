package sample.testcontainer.kanban

import org.springframework.boot.fromApplication
import org.springframework.boot.with


fun main(args: Array<String>) {
	fromApplication<SampleKanbanJvmApplication>().with(TestcontainersConfiguration::class).run(*args)
}
