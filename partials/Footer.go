package partials

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Footer struct {
	app.Compo
}

func (f *Footer) Render() app.UI {
	return app.
		Footer().
		Body(
			app.
				H5().
				Text("Footer"),
		)
}
