package pages

import (
	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func Login(c echo.Context) error {
	return layouts.MainPage("Login", []g.Node{}, []g.Node{}, []g.Node{}).Render(c.Response().Writer)
}
