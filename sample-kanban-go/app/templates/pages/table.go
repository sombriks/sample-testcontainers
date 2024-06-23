package pages

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func Table(c echo.Context) error {
	return layouts.MainPage("Table", []g.Node{}, []g.Node{}, []g.Node{}).Render(c.Response().Writer)
}
