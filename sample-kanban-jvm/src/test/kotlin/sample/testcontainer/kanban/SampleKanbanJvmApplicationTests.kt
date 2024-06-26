package sample.testcontainer.kanban

import org.junit.jupiter.api.Test
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.context.annotation.Import
import org.springframework.test.context.ActiveProfiles

@SpringBootTest
@ActiveProfiles("test")
// just add that and you have a full-featured, predictable, database for test!
@Import(TestcontainersConfiguration::class)
class SampleKanbanJvmApplicationTests {

	@Test
	fun contextLoads() {}

}
