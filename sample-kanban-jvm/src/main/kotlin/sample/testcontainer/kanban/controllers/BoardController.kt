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
        // it's just an example
        model.set("users", boardService
            .listPeople("", PageRequest.ofSize(10)).content)
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
        model.set("people", boardService
            .listPeople("", PageRequest.ofSize(10)).content)
        model.set("statuses", boardService.listStatuses())
        model.set("tasks", boardService
            .listTasks("", PageRequest.ofSize(100)).content)
        return "pages/board"
    }

    @GetMapping("table")
    fun table(model: Model, @CookieValue("x-user-info") info: String?): String {
        logger.info("table")
        if (info == null) return "redirect:/logout"
        model.set("user", Person.fromCookie(info))
        model.set("people", boardService
            .listPeople("", PageRequest.ofSize(10)).content)
        model.set("statuses", boardService.listStatuses())
        model.set("tasks", boardService
            .listTasks("", PageRequest.ofSize(100)).content)
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
        model.set("tasks", boardService
            .listTasks("", PageRequest.ofSize(100)).content)
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
        model.set("task", boardService.updateTaskStatus(data))
        return "components/task-card"
    }
}
