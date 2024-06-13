package sample.testcontainer.kanban.models

import jakarta.persistence.*
import java.time.LocalDateTime

@Entity
@Table(name = "status", schema = "kanban")
class Status(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
    var description: String? = null,
    var meansComplete: Boolean? = null,
    @Column(updatable = false)
    @Temporal(TemporalType.TIMESTAMP)
    var created: LocalDateTime? = null
)
