package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"go-app/components/header"
)

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Main().
		Body(
			&header.NavBar{},
		)
}
