package home_page

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/partials"
)

type HomePage struct {
	app.Compo
}

func (hp *HomePage) Render() app.UI {
	return app.
		Div().
		Body(
			&partials.Header{},
			&partials.NavBar{},
			//&partials.Footer{},
		)
}
