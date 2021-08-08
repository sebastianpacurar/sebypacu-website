package pages

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Third struct {
	app.Compo
}

func (t *Third) Render() app.UI {
	return app.Body().Body().Text("Third Page")
}
