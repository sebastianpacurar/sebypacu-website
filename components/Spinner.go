package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Spinner struct {
	app.Compo
	Reason string
}

func (s *Spinner) Render() app.UI {
	return app.
		Div().
		ID("spinner-container").
		Body(
			app.
				Div().
				ID("spinner-image").
				Body(
					app.
						Div().
						ID("spinner"),
				),
			app.
				H4().
				Text(fmt.Sprintf("Loading %s...", s.Reason)),
		)
}
