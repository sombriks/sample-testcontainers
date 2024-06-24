package components

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)

func CategoryLanes(user *models.Person, status *models.Status, tasks *[]models.Task) g.Node {
	return Div(Class("cell is-flex lane"),
		ID(fmt.Sprint("status", status.Id)),
		Section(Class("card is-flex is-flex-grow-1 is-flex-direction-column"),
			Header(Class("card-header"),
				H1(Class("card-header-title"), g.Text(status.Description)),
			),
			Div(Class("card-content is-flex-grow-1 status"),
				Data("status", fmt.Sprint(status.Id)),
				g.Attr("@dragover.prevent"),
				g.Attr("@drop", `e => {
					const taskId = e.dataTransfer.getData('text/plain')
					const taskEl = document.getElementById(taskId)
					$el.appendChild(taskEl)
					$dispatch('update-status', {taskEl, lane: $el})
				 }`),
				// tasks var came as parameter
				g.Group(g.Map(*tasks, func(task models.Task) g.Node {
					if task.StatusId == status.Id {
						return TaskCard(user, status, &task)
					} else {
						return g.Raw("")
					}
				})),
			),
			Div(Class("card-footer"),
				g.Attr("x-data", "{mode:0}"),
				Button(Class("button card-footer-item m-1"),
					g.Attr("x-show", "mode==0"),
					g.Attr("@click", "mode++"),
					Span(Class("icon"), g.El("ion-icon", Name("add-outline"))),
				),
				Form(Class("card-footer-item"),
					g.Attr("x-show", "mode==1"),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-target", fmt.Sprint("#status", status.Id)),
					g.Attr("hx-post", "task"),
					g.Attr("@click.outside", "mode=0"),
					Input(Type("hidden"), Name("status"), Value(fmt.Sprint(status.Id))),
					Div(Class("field has-addons"),
						Div(Class("control"),
							Input(Class("input"),
								Type("text"),
								Name("description"),
								Placeholder("Describe new task"),
							),
						),
						Div(Class("control"),
							Button(Type("submit"),
								Class("button is-info"),
								Span(Class("icon"),
									g.El("ion-icon", Name("save-outline")),
								),
							),
						),
					),
				),
			),
		),
	)
}
