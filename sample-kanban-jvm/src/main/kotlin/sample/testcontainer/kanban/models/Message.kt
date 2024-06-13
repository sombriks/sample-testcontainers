package sample.testcontainer.kanban.models

import jakarta.persistence.*
import java.time.LocalDateTime

@Entity
@Table(name = "message", schema = "kanban")
data class Message(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
    var content: String? = null,
    @Column(updatable = false)
    @Temporal(TemporalType.TIMESTAMP)
    var created: LocalDateTime? = null,
    @ManyToOne
    var person: Person? = null,
    @ManyToOne
    var task: Task? = null,
)
