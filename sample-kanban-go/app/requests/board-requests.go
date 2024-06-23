package requests

import (
	"github.com/labstack/echo/v4"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/services"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
	"log"
)

type BoardRequest struct {
	service *services.BoardService
}

// NewBoardRequest - provision the request handlers for the kanban
func NewBoardRequest(service *services.BoardService) (*BoardRequest, error) {
	request := &BoardRequest{service: service}
	return request, nil
}

func (r *BoardRequest) Index(c echo.Context) error {
	return c.Redirect(302, "/board")
}

func (r *BoardRequest) BoardPage(c echo.Context) error {

	return layouts.MainPage("Board").Render(c.Response().Writer)
}

func (r *BoardRequest) LoginPage(c echo.Context) error {

	return layouts.MainPage("Login").Render(c.Response().Writer)
}

func (r *BoardRequest) FakeLogin(c echo.Context) error {

	return c.HTML(200, "ok - login")
}

func (r *BoardRequest) FakeLogout(c echo.Context) error {

	return c.HTML(200, "ok - login")
}

func (r *BoardRequest) CookieCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("x-user-info")
		log.Println(cookie)
		if err != nil {
			log.Println("[WARN] ", err.Error())
			return c.Redirect(302, "/login")
		}
		if cookie == nil || cookie.Value == "" {
			return c.Redirect(302, "/login")
		}
		return next(c)
	}
}

func (r *BoardRequest) TablePage(c echo.Context) error {

	return layouts.MainPage("Table").Render(c.Response().Writer)
}

func (r *BoardRequest) AddTask(c echo.Context) error {

	return c.HTML(200, "ok - table")
}

func (r *BoardRequest) UpdateTask(c echo.Context) error {

	return c.HTML(200, "ok - table")
}

func (r *BoardRequest) DeleteTask(c echo.Context) error {

	return c.HTML(200, "ok - table")
}

func (r *BoardRequest) JoinTask(c echo.Context) error {

	return c.HTML(200, "ok - table")
}

func (r *BoardRequest) AddComent(c echo.Context) error {

	return c.HTML(200, "ok - table")
}
