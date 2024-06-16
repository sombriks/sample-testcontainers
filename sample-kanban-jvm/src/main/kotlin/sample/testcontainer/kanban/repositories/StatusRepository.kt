package sample.testcontainer.kanban.repositories

import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.jpa.repository.Query
import org.springframework.stereotype.Repository
import sample.testcontainer.kanban.models.Status

@Repository
interface StatusRepository : JpaRepository<Status, Long> {

    @Query("select t.status from Task t where t.id = :taskId")
    fun findStatusByTaskId(taskId: Long): Status?
}
