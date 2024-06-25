package components

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)

func TaskCard(user *models.Person, status *models.Status, task *models.Task) g.Node {
	return Article(Class("message task"),
		g.Attr("draggable", "true"),
		g.Attr("x-data", "{openModal:false}"),
		g.Attr("hx-ext", "hx-dataset-include"),
		g.Attr("hx-swap", "outerHTML"),
		g.Attr("hx-trigger", "put-task"),
		ID(fmt.Sprint("task", task.Id)),
		Data("task", fmt.Sprint("", task.Id)),
		Data("description", fmt.Sprint("", task.Description)),
		Data("status", fmt.Sprint("", task.StatusId)), // task.Status.Id
		g.Attr("hx-put", fmt.Sprint("task/", task.Id)),
		g.Attr("@dragstart", "e => e.dataTransfer.setData('text/plain', $el.id)"),
		g.Attr("@update-status.window", `e => {
			if(e.detail.taskEl == $el) { // needed to filter real target
				$el.dataset.status = e.detail.lane.dataset.status
				$refs['form-status-'+$el.dataset.task].value = e.detail.lane.dataset.status
				$dispatch('put-task', $el.dataset)
			}
		 }`),
		Div(Class("message-header"), g.Text(fmt.Sprint("#", task.Id, " - ", task.Description))),
		Div(Class("message-body is-flex is-justify-content-space-between is-align-content-center"),
			Span(Class("icon-text"),
				Span(Class("icon"), g.El("ion-icon", Name("chatbox-ellipses-outline"))),
				Span(g.Text(fmt.Sprint(task.SafeMessageCount()))),
				Span(Class("icon"), g.El("ion-icon", Name("people-circle-outline"))),
				Span(g.Text(fmt.Sprint(task.SafePeopleCount()))),
			),
			Button(Class("button"),
				Span(Class("icon"), g.El("ion-icon", Name("information-circle-outline"))),
				g.Attr("@click", "openModal=true"),
			),
			TaskModal(user, status, task),
		),
	)
}
