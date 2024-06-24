package requests

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/services"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/components"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/pages"
	"log"
	"net/http"
)

type BoardRequest struct {
	service *services.BoardService
}

// NewBoardRequest - provision the request handlers for the kanban
func NewBoardRequest(service *services.BoardService) (*BoardRequest, error) {
	request := &BoardRequest{service: service}

	return request, nil
}

func getUser(c echo.Context) *models.Person {
	inContext := c.Get("user")
	if inContext == nil {
		return nil
	}
	return (inContext).(*models.Person)
}

// Index - simple redirect to board page
func (r *BoardRequest) Index(c echo.Context) error {
	return c.Redirect(302, "/board")
}

// BoardPage - route to provision and serve the kanban board page
func (r *BoardRequest) BoardPage(c echo.Context) error {
	user := getUser(c)
	statuses, err := r.service.ListStatus()
	if err != nil {
		return err
	}
	tasks, err := r.service.ListTasks("")
	if err != nil {
		return err
	}
	return pages.BoardPage(user, statuses, tasks).Render(c.Response().Writer)
}

func (r *BoardRequest) LoginPage(c echo.Context) error {
	user := getUser(c)
	people, err := r.service.ListPeople("")
	if err != nil {
		return err
	}
	return pages.Login(user, people).Render(c.Response().Writer)
}

func (r *BoardRequest) FakeLogin(c echo.Context) error {
	userId := c.FormValue("userId")
	var id int64
	_, _ = fmt.Sscan(userId, &id)
	user, _ := r.service.FindPerson(id)
	cookie := http.Cookie{
		Name:  "x-user-info",
		Value: user.UserToCookie(),
	}
	c.SetCookie(&cookie)
	return c.Redirect(302, "/board")
}

func (r *BoardRequest) FakeLogout(c echo.Context) error {
	cookie, err := c.Cookie("x-user-info")
	if err != nil {
		log.Println("[WARN] ", err.Error())
		return c.Redirect(302, "/login")
	}
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.Redirect(302, "/login")
}

func (r *BoardRequest) CookieCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("x-user-info")
		if err != nil {
			log.Println("[WARN] ", err.Error())
			return c.Redirect(302, "/login")
		}
		if cookie == nil || cookie.Value == "" {
			return c.Redirect(302, "/login")
		}
		c.Set("user", models.UserFromCookie(cookie.Value))
		return next(c)
	}
}

func (r *BoardRequest) TablePage(c echo.Context) error {
	user := getUser(c)
	return pages.TablePage(user).Render(c.Response().Writer)
}

func (r *BoardRequest) AddTask(c echo.Context) error {
	user := getUser(c)
	var statusId int64 = 0
	fmt.Sscan(c.FormValue("status"), &statusId)
	description := ""
	fmt.Sscan(c.FormValue("description"), &description)
	result, err := r.service.InsertTask(&models.Task{
		Description: description,
		StatusId:    statusId,
	})
	log.Printf("inserted task: %v\n", result)
	if err != nil {
		return err
	}
	status, err := r.service.FindStatus(statusId)
	if err != nil {
		return err
	}
	tasks, err := r.service.ListTasks("")
	if err != nil {
		return err
	}
	return components.CategoryLanes(user, status, tasks).Render(c.Response().Writer)
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
