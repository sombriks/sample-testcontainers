package sample.testcontainer.kanban.controllers

import jakarta.servlet.http.Cookie
import jakarta.servlet.http.HttpServletResponse
import jakarta.servlet.http.HttpSession
import org.slf4j.LoggerFactory
import org.springframework.data.domain.PageRequest
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.ui.set
import org.springframework.web.bind.annotation.*
import sample.testcontainer.kanban.models.Person
import sample.testcontainer.kanban.models.Task
import sample.testcontainer.kanban.models.to.TaskStatusTO
import sample.testcontainer.kanban.services.BoardService

@Controller // not a RestController, we want to render views
@RequestMapping("/")
class BoardController(private val boardService: BoardService) {

    val logger by lazy { LoggerFactory.getLogger(BoardController::class.java) }

    @GetMapping
    fun index(@CookieValue("x-user-info") info: String?): String {
        logger.info("x-user-info: $info")
        return if (info == null) "redirect:/login" else "redirect:/board"
    }

    @GetMapping("login")
    fun login(model: Model): String {
        logger.info("login")
        // it's just an example, no real login here
        val users = boardService.listPeople("", PageRequest.ofSize(10)).content
        model.set("users", users)
        return "pages/login"
    }

    @PostMapping("login")
    fun doLogin(model: Model, userId: Long?, response: HttpServletResponse): String {
        logger.info("userId: $userId")
        if (userId == null) return "redirect:/login"
        val user = boardService.findPerson(userId)
        model.set("user", user)
        val cookie = Cookie("x-user-info", "name=${user?.name}&id=${user?.id}")
        cookie.maxAge = -1
        response.addCookie(cookie)
        return "redirect:/board"
    }

    @GetMapping("logout")
    fun logout(response: HttpServletResponse, session: HttpSession): String {
        logger.info("logout")
        val cookie = Cookie("x-user-info", null)
        cookie.maxAge = 0
        response.addCookie(cookie)
        session.invalidate()
        return "redirect:/login"
    }

    @GetMapping("board")
    fun board(model: Model, @CookieValue("x-user-info") info: String?): String {
        logger.info("board")
        if (info == null) return "redirect:/logout"
        // TODO consider use https://docs.spring.io/spring-framework/reference/web/webmvc/mvc-controller/ann-methods/sessionattribute.html
        model.set("user", Person.fromCookie(info))
        val people = boardService.listPeople("", PageRequest.ofSize(10)).content
        model.set("people", people)
        model.set("statuses", boardService.listStatuses())
        val tasks = boardService.listTasks("", PageRequest.ofSize(100)).content
        model.set("tasks", tasks)
        return "pages/board"
    }

    @GetMapping("table")
    fun table(model: Model, @CookieValue("x-user-info") info: String?): String {
        logger.info("table")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        val people = boardService.listPeople("", PageRequest.ofSize(10)).content
        model.set("people", people)
        model.set("statuses", boardService.listStatuses())
        val tasks = boardService.listTasks("", PageRequest.ofSize(100)).content
        model.set("tasks", tasks)
        return "pages/table"
    }

    @PostMapping("task")
    fun createTask(
        model: Model, @CookieValue("x-user-info") info: String?,
        data: TaskStatusTO,
    ): String {
        logger.info("createTask")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        val status = boardService.findStatus(data.status!!)
        model.set("status", status)
        val task = Task(description = data.description, status = status)
        boardService.saveTask(task)
        val tasks = boardService.listTasks("", PageRequest.ofSize(100)).content
        model.set("tasks", tasks)
        return "components/category-lanes"
    }

    @PutMapping("task/{id}")
    fun updateTask(
        model: Model,
        @CookieValue("x-user-info") info: String?,
        @PathVariable id: Long,
        data: TaskStatusTO,
    ): String {
        logger.info("updateTask")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        if (id != data.task) return "redirect:/error"
        model.set("task", boardService.updateTask(data))
        val status = boardService.findStatus(data.status!!)
        model.set("status", status)
        return "components/task-card"
    }

    @DeleteMapping("task/{id}")
    fun deleteTask(
        model: Model,
        @CookieValue("x-user-info") info: String?,
        @PathVariable id: Long,
    ): String {
        logger.info("deleteTask")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        val status = boardService.findStatusByTaskId(id)
        model.set("status", status)
        boardService.deleteTask(id)
        val tasks = boardService.listTasks("", PageRequest.ofSize(100)).content
        model.set("tasks", tasks)
        return "components/category-lanes"
    }

    @PostMapping("task/{id}/join")
    fun joinTask(
        model: Model,
        @CookieValue("x-user-info") info: String?,
        @PathVariable id: Long,): String {
        logger.info("joinTask")
        if (info == null) return "redirect:/logout"
        val person = Person.fromCookie(info)
        boardService.joinTask(taskId = id, personId = person.id)
        model.set("user", person)
        val task = boardService.findTask(id)
        model.set("task", task)
        return "components/task-members"
    }

    @DeleteMapping("task/{taskId}/person/{id}")
    fun removePersonFromTask(
        model: Model,
        @CookieValue("x-user-info") info: String?,
        @PathVariable taskId: Long,
        @PathVariable id: Long,
    ): String {
        logger.info("removePersonFromTask")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        boardService.removePersonFromTask(
            taskId = taskId,
            personId = id,
        )
        val task = boardService.findTask(taskId)
        model.set("task", task)
        return "components/task-members"
    }
}
