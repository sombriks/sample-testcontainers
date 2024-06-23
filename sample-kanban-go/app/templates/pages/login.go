package pages

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/templates/layouts"
)

func Login(user *models.Person, people *[]models.Person) g.Node {
	return layouts.MainPage(user, layouts.PageSlots{
		Title: "Login", InBody: []g.Node{
			Div(Class("is-mobile is-centered columns"),
				Div(Class("is-half column"),
					Form(Class("card"),
						Action("login"),
						Method("post"),
						Header(Class("card-header"),
							H1(Class("card-header-title title"), g.Text("Login")),
						),
						Section(Class("card-content"),
							Div(Class("field"),
								Label(Class("label"), For("userId"), g.Text("Pick an user")),
								Div(Class("control select is-primary is-fullwidth"),
									Select(ID("userId"), Name("userId"),
										g.Group(g.Map(*people, func(person models.Person) g.Node {
											return Option(Value(fmt.Sprint(person.Id)), g.Text(person.Name))
										})),
									),
								),
							),
						),
						Footer(Class("card-footer"),
							Button(Class("button is-primary card-footer-item m-1"),
								Type("submit"),
								g.Text("Pick this user"),
							),
						),
					),
				),
			),
		},
	})
}
