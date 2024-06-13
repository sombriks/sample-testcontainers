package sample.testcontainer.kanban.repositories

import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import sample.testcontainer.kanban.models.Status

@Repository
interface StatusRepository : JpaRepository<Status, Long>
