package sample.testcontainer.kanban.controllers

import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import sample.testcontainer.kanban.services.BoardService

@Controller
@RequestMapping("/")
class BoardController(private val boardService: BoardService) {

    @GetMapping
    fun index(): String {

        return "login" // login, board, table (dialog)
    }

}
