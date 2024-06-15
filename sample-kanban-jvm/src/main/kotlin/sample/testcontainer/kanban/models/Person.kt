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
    var created: LocalDateTime? = null,
    @ManyToMany
    @JoinTable(
        schema = "kanban",
        name = "task_person",
        joinColumns = [JoinColumn(name = "person_id")],
        inverseJoinColumns = [JoinColumn(name = "task_id")]
    )
    @OrderBy("created desc")
    var tasks: MutableList<Task>? = null,
) {
    companion object {
        fun fromCookie(info: String): Person {
            val person = Person()
            info.split("&").forEach {
                val (key, value) = it.split("=")
                when (key) {
                    "id" -> person.id = value.toLong()
                    "name" -> person.name = value
                }
            }
            return person
        }
    }
}
