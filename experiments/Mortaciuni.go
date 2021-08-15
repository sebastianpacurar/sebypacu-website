package experiments

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/components"
)

type Mortaciuni struct {
	app.Compo
}

func (m *Mortaciuni) Render() app.UI {
	return app.
		Div().
		Body(
			&components.Header{},
			&components.NavBar{},
			app.
				Div().
				Style("height", "300px").
				Style("width", "100%").
				Style("background-color", "antiquewhite"),
			&components.Footer{},
		)
}
