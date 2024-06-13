package sample.testcontainer.kanban.models

import jakarta.persistence.*
import java.time.LocalDateTime

@Entity
@Table(name = "task", schema = "kanban")
data class Task(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
    var description: String? = null,
    @Column(updatable = false)
    @Temporal(TemporalType.TIMESTAMP)
    var created: LocalDateTime? = null,
    @ManyToOne
    var status: Status? = null,
    @OneToMany
    @JoinColumn(name = "task_id")
    var messages: List<Message>? = null
)
