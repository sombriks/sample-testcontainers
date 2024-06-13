package sample.testcontainer.kanban.repositories

import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import sample.testcontainer.kanban.models.Message

@Repository
interface MessageRepository : JpaRepository<Message, Long> {

    fun findByContentContainingIgnoreCase(q: String, pageable: Pageable): Page<Message>
}
