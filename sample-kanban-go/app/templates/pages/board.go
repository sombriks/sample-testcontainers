package pages

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func Board(c echo.Context) error {
	return layouts.MainPage("Board", []g.Node{}, []g.Node{}, []g.Node{}).Render(c.Response().Writer)
}
