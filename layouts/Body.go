package layouts

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"go-app/components"
)

type Body struct {
	Value string
}

func (b *Body) Render() app.UI {
	return app.Article().
		ID("go-wasm-pwa").
		Style("margin", "10px").
		Body(
			&components.Header{},

			&components.Footer{},
		)
}
