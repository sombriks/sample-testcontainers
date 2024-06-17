package sample.testcontainer.kanban.services

import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.Matchers.notNullValue
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Disabled
import org.junit.jupiter.api.Test
import org.mockito.ArgumentMatchers.anyString
import org.mockito.Mock
import org.mockito.kotlin.anyOrNull
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.boot.test.mock.mockito.MockBean
import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.test.context.ActiveProfiles
import sample.testcontainer.kanban.models.Person
import sample.testcontainer.kanban.repositories.PersonRepository
import org.mockito.Mockito.`when` as _when

@SpringBootTest
@ActiveProfiles("test")
class BoardServiceTestWithMockito {

    @Autowired
    private lateinit var boardService: BoardService

    @MockBean
    private lateinit var personRepository: PersonRepository

    @Mock
    private lateinit var pageable: Pageable

    @Mock
    private lateinit var personPage: Page<Person>

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

    @Test
    @Disabled("We can keep mocking but we don't trust the test anymore")
    fun `should save people`() {
        val person = Person(name = "Ferdinando")
        boardService.savePerson(person)
        assertThat(person.id, notNullValue()) //new person should have an id now
    }
}
