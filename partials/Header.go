package partials

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.
		Header().
		ID("app-header").
		Body(
			app.
				H2().
				Text("Bye and FU, JavaScript! ╭∩╮(・◞・) "),
		)
}
