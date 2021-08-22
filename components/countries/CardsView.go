package countries

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Grid struct {
	app.Compo
	data []CountryInfo
}

func (g *Grid) Render() app.UI {
	return app.
		Article().
		ID("cards-container").
		Body(
			app.Range(g.data).Slice(func(i int) app.UI {
				current := g.data[i]

				return app.Div().Class("country-card").Body(
					app.
						Figure().
						ID(fmt.Sprintf("%s-figure", current.Name)).
						Body(
							app.
								FigCaption().
								ID(fmt.Sprintf("%s-title", current.Name)).
								Body(
									app.
										P().
										Text(current.Name),
								),
							app.
								Img().
								Src(current.Flag).
								Alt(fmt.Sprintf("%s Flag", current.Name)),
							app.
								Section().
								ID(fmt.Sprintf("%s-details", current.Name)).
								Body(
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Capital"),
											app.
												P().
												Text(current.Capital),
										),
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Region"),
											app.
												P().
												Text(current.Region),
										),
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Subregion"),
											app.
												P().
												Text(current.Subregion),
										),
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Population"),
											app.
												P().
												Text(current.Population),
										),
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Area"),
											app.
												P().
												Text(current.Area),
										),
									app.
										Div().
										Class("card-info").
										Body(
											app.
												P().
												Text("Native Name"),
											app.
												P().
												Text(current.NativeName),
										),
									app.
										Button().
										ID(current.Alpha2Code).
										Class("view-country-btn").
										Text("View Country").
										OnClick(NavigateToCountry),
								),
						),
				)
			}),
		)
}
