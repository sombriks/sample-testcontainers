package layouts

import (
	"fmt"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
)

func MainPage(title string, inHead []g.Node, inBody []g.Node, inFooter []g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: fmt.Sprintf("Page - %s", title),
		Head:  append([]g.Node{}, inHead...),
		Body:  append(append([]g.Node{}, inBody...), inFooter...),
	})
}
