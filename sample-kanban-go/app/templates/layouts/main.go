package layouts

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func MainPage(title string) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: fmt.Sprintf("Page - %s", title),
		Head:  []g.Node{},
		Body: []g.Node{
			h.H1(g.Text(title)),
		},
	})
}
