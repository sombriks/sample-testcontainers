package components

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)
import . "github.com/maragudk/gomponents/html"

func TaskComments(user *models.Person, task *models.Task) g.Node {
	return Div(ID(fmt.Sprint("task-comments-", task.Id)),
		H2(g.Text("Comments")),
		Form(Class("card-footer-item"),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", fmt.Sprint("#task-comments-", task.Id)),
			g.Attr("hx-post", fmt.Sprint("task/", task.Id, "/comments")),
			Input(Type("hidden"), Name("taskId"), Value(fmt.Sprint(task.Id))),
			Input(Type("hidden"), Name("personId"), Value(fmt.Sprint(user.Id))),
			Div(Class("field has-addons"),
				Div(Class("control"),
					Input(Class("input"), Required(),
						Type("text"),
						Name("content"),
						Placeholder("Add comment"),
					),
				),
				Div(Class("control"),
					Button(Class("button is-info"), Type("submit"),
						Span(Class("icon"),
							g.El("ion-icon", Name("save-outline")),
						),
					),
				),
			),
		),
		g.Group(g.Map(*task.SafeMessageList(), func(message models.Message) g.Node {
			return P(Class("notification"),
				Title(fmt.Sprint("comment #", message.Id, " by ",
					message.Person.Name, " at ", message.Created)),
				g.Text(message.Content))
		})),
	)
}
