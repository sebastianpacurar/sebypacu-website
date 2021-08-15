package pages

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/components"
)

type Home struct {
	app.Compo
}

func (hp *Home) Render() app.UI {
	return app.Div().
		Body(
			&components.Header{},
			&components.NavBar{},
			&components.Footer{},
		)
}
