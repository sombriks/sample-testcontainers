package sample.testcontainer.kanban.models.to

data class MessageTO(
    var content: String? = null,
    var personId: Long? = null,
    var taskId: Long? = null,
)
