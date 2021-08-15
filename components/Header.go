package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Header().
		Body(
			app.
				H1().
				Text("Placeholder head"),
		)
}
