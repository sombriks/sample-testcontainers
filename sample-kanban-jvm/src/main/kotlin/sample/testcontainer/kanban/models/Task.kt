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
    @OrderBy("created desc")
    var messages: List<Message>? = null,
    @ManyToMany
    @JoinTable(
        schema = "kanban",
        name = "task_person",
        joinColumns = [JoinColumn(name = "task_id")],
        inverseJoinColumns = [JoinColumn(name = "person_id")],
    )
    @OrderBy("created desc")
    var people: MutableList<Person>? = null,
) {
    override fun toString(): String {
        return "Task(id=$id, description=$description, created=$created, status=$status)"
    }
}
