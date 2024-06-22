package requests

import (
	"github.com/labstack/echo/v4"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/services"
)

type BoardRequest struct {
	service *services.BoardService
}

// NewBoardRequest - provision the request handlers for the kanban
func NewBoardRequest(service *services.BoardService) (*BoardRequest, error) {
	request := &BoardRequest{service: service}
	return request, nil
}

func (r *BoardRequest) BoardPage(c echo.Context) error {

	return c.HTML(200, "ok - board")
}

func (r *BoardRequest) LoginPage(c echo.Context) error {

	return c.HTML(200, "ok - login")
}
func (r *BoardRequest) TablePage(c echo.Context) error {

	return c.HTML(200, "ok - table")
}
