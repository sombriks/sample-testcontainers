package pages

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func BoardPage(user *models.Person) g.Node {
	return layouts.MainPage(user, layouts.PageSlots{
		Title: "BoardPage",
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
		},
	})
}
