package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/partials"
)

type Home struct {
	app.Compo
}

func (hp *Home) Render() app.UI {
	return app.
		Div().
		Body(
			&partials.Header{},
			&partials.NavBar{},
			//&partials.Footer{},
		)
}
