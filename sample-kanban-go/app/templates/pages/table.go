package pages

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func TablePage(user *models.Person) g.Node {
	return layouts.MainPage(user, layouts.PageSlots{
		Title: "Table",
		InBody: []g.Node{
			html.H1(g.Text(fmt.Sprint("Table for user ", user.Name))),
		},
	})
}
