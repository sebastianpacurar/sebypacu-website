package pages

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Second struct {
	app.Compo
}

func (s *Second) Render() app.UI {
	return app.Body().Body().Text("Second Page")
}
