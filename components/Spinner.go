package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Spinner struct {
	app.Compo
}

func (s *Spinner) Render() app.UI {
	return app.
		Div().
		ID("spinner")
}
