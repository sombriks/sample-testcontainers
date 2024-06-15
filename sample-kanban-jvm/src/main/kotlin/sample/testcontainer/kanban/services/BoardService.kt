package sample.testcontainer.kanban.services

import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.data.domain.Sort
import org.springframework.data.repository.findByIdOrNull
import org.springframework.stereotype.Service
import sample.testcontainer.kanban.models.Message
import sample.testcontainer.kanban.models.Person
import sample.testcontainer.kanban.models.Status
import sample.testcontainer.kanban.models.Task
import sample.testcontainer.kanban.models.to.TaskStatusTO
import sample.testcontainer.kanban.repositories.MessageRepository
import sample.testcontainer.kanban.repositories.PersonRepository
import sample.testcontainer.kanban.repositories.StatusRepository
import sample.testcontainer.kanban.repositories.TaskRepository
import java.time.LocalDateTime

@Service
class BoardService(
    private val messageRepository: MessageRepository,
    private val personRepository: PersonRepository,
    private val statusRepository: StatusRepository,
    private val taskRepository: TaskRepository,
) {

    fun listMessages(q: String, pageable: Pageable): Page<Message> {
        return messageRepository.findByContentContainingIgnoreCase(q, pageable)
    }

    fun listPeople(q: String, pageable: Pageable): Page<Person> {
        return personRepository.findByNameContainingIgnoreCase(q, pageable)
    }

    fun listStatuses(): List<Status> {
        return statusRepository.findAll(Sort.by(Sort.Direction.ASC, "id"))
    }

    fun listTasks(q: String, pageable: Pageable): Page<Task> {
        return taskRepository.findByDescriptionContainingIgnoreCase(q, pageable)
    }

    fun findMessage(id: Long): Message? {
        return messageRepository.findByIdOrNull(id)
    }

    fun findPerson(id: Long): Person? {
        return personRepository.findByIdOrNull(id)
    }

    fun findStatus(id: Long): Status? {
        return statusRepository.findByIdOrNull(id)
    }

    fun findTask(id: Long): Task? {
        return taskRepository.findByIdOrNull(id)
    }

    fun saveMessage(message: Message) {
        if (message.id == null) message.created = LocalDateTime.now()
        messageRepository.save(message)
    }

    fun savePerson(person: Person) {
        if (person.id == null) person.created = LocalDateTime.now()
        personRepository.save(person)
    }

    fun saveTask(task: Task) {
        if (task.id == null) task.created = LocalDateTime.now()
        taskRepository.save(task)
    }

    fun updateTaskStatus(data: TaskStatusTO): Task {
        val task = findTask(data.task!!)!!
        val status = findStatus(data.status!!)!!
        task.status = status
        saveTask(task)
        return task
    }
}
