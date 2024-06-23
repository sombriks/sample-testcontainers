package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func TablePage(user *models.Person) g.Node {
	return layouts.MainPage(user, layouts.PageSlots{
		Title: "TablePage",
	})
}
