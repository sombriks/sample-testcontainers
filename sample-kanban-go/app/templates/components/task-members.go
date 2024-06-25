package components

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)
import . "github.com/maragudk/gomponents/html"

func TaskMembers(user *models.Person, task *models.Task) g.Node {

	classLoggedUser := func(person *models.Person) string {
		if user.Id == person.Id {
			return "tag is-primary"
		}
		return "tag"
	}

	return Div(ID(fmt.Sprint("task-members-", task.Id)),
		H2(g.Text("People working on this task")),
		Div(Class("field is-grouped is-grouped-multiline"),
			g.Group(g.Map(*task.SafePeopleList(), func(person models.Person) g.Node {
				return Div(Class("control"),
					Div(Class("tags has-addons"),
						Span(Class(classLoggedUser(&person)), g.Text(person.Name)),
						Span(Class("tag is-delete"),
							g.Attr("hx-swap", "outerHTML"),
							g.Attr("hx-confirm", "Are you sure?"),
							g.Attr("hx-target", fmt.Sprint("#task-members-", task.Id)),
							g.Attr("hx-delete", fmt.Sprint("task/", task.Id, "/person/", person.Id)),
						),
					),
				)
			})),
		),
		g.If(task.MemberById(user.Id) == nil, Div(Class("buttons is-right"),
			Button(Class("button is-primary"),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-target", fmt.Sprint("#task-members-", task.Id)),
				g.Attr("hx-post", fmt.Sprint("task/", task.Id, "/join")),
				Span(Class("icon"), g.El("ion-icon", Name("person-add-outline"))),
				Span(g.Text("Join this task")),
			),
		)),
	)
}
