package sample.testcontainer.kanban.services

import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.Matchers.*
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.Import
import org.springframework.data.domain.PageRequest
import org.springframework.data.domain.Sort
import org.springframework.test.context.ActiveProfiles
import org.springframework.transaction.annotation.Transactional
import sample.testcontainer.kanban.TestcontainersConfiguration
import sample.testcontainer.kanban.models.Message
import sample.testcontainer.kanban.repositories.PersonRepository
import kotlin.test.assertNotNull

@SpringBootTest
@ActiveProfiles("test")
@Import(TestcontainersConfiguration::class)
class BoardServiceTestWithTestContainers {

    @Autowired
    private lateinit var boardService: BoardService

    @Autowired
    private lateinit var personRepository: PersonRepository

    @Test
    fun `should list messages`() {
        val result = boardService.listMessages(
            pageable = PageRequest.of(
                0, 10,
                Sort.by(Sort.Direction.DESC, "id")
            ),
            q = ""
        )
        assertNotNull(result)
        assertThat(result.content.map { it.content }, hasItem("Need this ASAP"))
        assertThat(result.content.map { it.person?.name }, hasItem("Caesar"))
        assertThat(result.content.map { it.task?.description }, hasItem("feature listing"))
    }

    @Test
    fun `should list people`() {
        val result = boardService.listPeople("", PageRequest.ofSize(10))
        assertNotNull(result)
        assertThat(result.content.map { it.name }, hasItems("Bob", "Alice"))
    }

    @Test
    fun `should list statuses`() {
        val result = boardService.listStatuses()
        assertNotNull(result)
        assertThat(result.map { it.description }, containsInAnyOrder("TODO", "DOING", "DONE"))
        assertThat(result.filter { it.meansComplete!! } .map { it.description }, hasItem("DONE"))
    }

    @Test
    fun `should list tasks`() {
        val result = boardService.listTasks("", PageRequest.ofSize(10))
        assertThat(result, notNullValue())
        assertThat(result.map { it.description }, hasItems("feature listing", "design", "data provision"))
    }

    @Test
    @Transactional
    fun `should comment a task`() {
        val person = boardService.findPerson(1)
        val task = boardService.findTask(1)
        val message = Message(person = person, task = task, content = "How is it going?")
        boardService.saveMessage(message)
        assertThat(message.id, notNullValue())

        val (result) = boardService.findTask(1)?.messages!!
        assertThat(result, notNullValue())
        assertThat(result.id, notNullValue())
        assertThat(result.content, equalTo("How is it going?"))
    }

    @Test
    @Transactional
    fun `should add person to task`() {
        val person = boardService.findPerson(5)!! // person with no tasks
        val task = boardService.findTask(2)!!
        task.people?.add(person)
        boardService.saveTask(task)

        personRepository.flush() // can't wait transaction end

        val (result) = boardService.findPerson(5)?.tasks!!
        assertThat(result, notNullValue())
        assertThat(result.id, notNullValue())
        assertThat(result, equalTo(task))
    }
}
