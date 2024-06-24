package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)

func TaskCard(user *models.Person, status *models.Status, task *models.Task) g.Node {
	return Div(g.Text(task.Description))
}
