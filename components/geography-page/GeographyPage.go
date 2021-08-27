package geography_page

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/partials"
)

type GeographyPage struct {
	app.Compo
}

func (gp *GeographyPage) Render() app.UI {
	return app.
		Div().
		Body(
			&partials.Header{},
			&partials.NavBar{},
			app.
				Main().
				ID("geography-page").
				Body(
					app.A().Href("/geography/countries/all").Text("Under Construction! Go To Countries Page"),
				),
		)
}
