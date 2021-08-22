package countries

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/utils"
)

type Table struct {
	app.Compo
	data []CountryInfo
}

func (t *Table) Render() app.UI {
	return app.
		Table().
		Body(
			app.
				Caption().
				ID("table-title").
				Text("All Countries"),
			app.
				THead().
				Body(
					app.
						Tr().
						Body(
							app.Th().
								Text("Name"),
							app.Th().
								Text("Flag"),
							app.Th().
								Text("Capital"),
							app.Th().
								Text("Region"),
						),
				),
			app.
				TBody().
				Body(
					app.
						Range(t.data).Slice(func(i int) app.UI {
						current := t.data[i]
						return app.
							Tr().
							ID(current.Alpha2Code).
							DataSet("country", current.Name).
							Body(
								app.
									Td().
									Text(current.Name),
								app.
									Td().
									Class("image-cell").
									Body(
										app.
											Img().
											Src(current.Flag).
											Alt(current.Name)),
								app.
									Td().
									Body(
										app.If(current.Capital == "",
											app.P().
												Text("N/A"),
										).Else(
											app.P().
												Text(current.Capital),
										),
									),
								app.
									Td().
									Body(
										app.If(current.Region == "",
											app.P().
												Text("N/A"),
										).Else(
											app.P().
												Text(current.Region),
										),
									),
							).OnClick(NavigateToCountry)
					}),
				),
			app.
				Tfoot().
				Body(
					app.
						Tr().
						Body(
							app.
								Td().
								ColSpan(4).
								Text("Go to Top"),
						),
				).OnClick(utils.ScrollToTop),
		)
}
