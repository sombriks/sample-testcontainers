package layouts

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)

// PageSlots - our slots for layout filling
type PageSlots struct {
	Title    string
	InHead   []g.Node
	InBody   []g.Node
	InFooter []g.Node
}

// MainPage - core layout with all frontend styles, scripting and other resources
func MainPage(user *models.Person, slots PageSlots) g.Node {
	return Doctype(HTML(
		Head(Meta(Charset("utf-8")),
			Meta(Name("viewport"),
				Content("width=device-width,initial-scale=1.0,minimum-scale=0.5,maximum-scale=2.0")),
			TitleEl(g.Text("Simple Kanban - "), g.Text(slots.Title)),
			Link(Rel("icon"), Href("favicon.png")),
			Link(Rel("stylesheet"), Href("node_modules/bulma/css/bulma.css")),
			Script(Type("text/javascript"), Src("node_modules/htmx.org/dist/htmx.js")),
			Script(Type("text/javascript"), Src("node_modules/hx-dataset-include/lib/hx-dataset-include.js")),
			Script(Type("text/javascript"), Defer(), Src("node_modules/alpinejs/dist/cdn.js")),
			Script(Type("module"), Src("node_modules/ionicons/dist/ionicons/ionicons.esm.js")),
			Script(Type("text/javascript"), g.Attr("nomodule"),
				Src("node_modules/ionicons/dist/ionicons/ionicons.js")),
			g.Group(slots.InHead),
		),
		Body(g.Attr("x-data"),
			g.Attr("hx-boost"),
			// adding x-data here makes entire document capable of use alpinejs
			//  adding hx-boost makes progressive enhancements in the application
			Div(Class("container is-max-widescreen"),
				g.If(user != nil, Nav(Class("navbar"),
					Div(Class("navbar-menu"),
						Div(Class("navbar-start"),
							A(Class("navbar-item"), Href("board"), g.Text("Board")),
							A(Class("navbar-item"), Href("table"), g.Text("Table")),
						),
						Div(Class("navbar-end"),
							A(Class("navbar-item"), Href("logout"), g.Text("Logout")),
						),
					),
				)),
				Section(Class("section"), g.Group(slots.InBody)),
				Footer(Class("footer"), g.Group(slots.InFooter)),
			),
		),
	))
}
