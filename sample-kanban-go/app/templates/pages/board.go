package pages

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/components"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func BoardPage(user *models.Person, statuses *[]models.Status, tasks *[]models.Task) g.Node {
	return layouts.MainPage(user, layouts.PageSlots{
		Title: "Board",
		InHead: []g.Node{
			StyleEl(g.Raw(`        
				.lanes {
					min-width: 600px;
					min-height: 60vh;
				}
				.lane {
					min-width: 200px;
				}
			`)),
		},
		InBody: []g.Node{
			H1(g.Text(fmt.Sprint("Welcome to the board, ", user.Name))),
			Div(Class(fmt.Sprint("lanes fixed-grid has-", len(*statuses), "-cols")),
				Div(Class("grid"),
					g.Group(g.Map(*statuses, func(status models.Status) g.Node {
						return components.CategoryLanes(user, &status, tasks)
					})),
				),
			),
		},
	})
}
