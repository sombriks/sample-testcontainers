package sample.testcontainer.kanban.models

import jakarta.persistence.*
import java.time.LocalDateTime

@Entity
@Table(name = "person", schema = "kanban")
data class Person(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
    var name: String? = null,
    @Column(updatable = false)
    @Temporal(TemporalType.TIMESTAMP)
    var created: LocalDateTime? = null
)
