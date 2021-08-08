package pages

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Fourth struct {
	app.Compo
}

func (f *Fourth) Render() app.UI {
	return app.Body().Body().Text("First Page")
}
