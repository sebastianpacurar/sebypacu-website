package pages

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"pwa/API"
	"pwa/components"
)

type Countries struct {
	app.Compo
	CountryData
}

type CountryData struct {
	Info []Country
}

type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Region     string `json:"region"`
	Flag       string `json:"flag"`
	Alpha2Code string `json:"alpha2code"`
}

func (c *Countries) OnNav(ctx app.Context) {
	if err := c.initCountries(ctx); err != nil {
		return
	}
	c.Update()
}

func (c *Countries) Render() app.UI {
	return app.
		Div().
		Body(
			&components.Header{},
			&components.NavBar{},
			app.
				If(len(c.Info) > 0,
					app.
						Main().
						ID("countries-container").
						Body(
							app.
								Form().
								ID("search-countries").
								Body(
									app.
										Input().
										ID("country-input").
										Type("text").
										Placeholder("Filter Countries by Letters"),
									app.
										Button().
										Type("submit").
										Value("Submit").
										Text("Fetch!"),
								),
							app.
								Table().
								Body(
									app.
										Caption().
										ID("table-title").
										Text("List of All Countries and independent Islands, Regions"),
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
												Range(c.Info).Slice(func(i int) app.UI {
												current := c.Info[i]
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
																app.P().
																	Class("h4").
																	Text(current.Capital),
															),
														app.
															Td().
															Body(
																app.P().
																	Class("h4").
																	Text(current.Region),
															),
													).OnClick(c.OnCountryClick)
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
										).OnClick(scrollToCountriesSearch),
								),
							&components.Footer{},
						),
				).
				Else(
					&components.Spinner{},
				),
		)
}

func (c *Countries) initCountries(ctx app.Context) error {

	data, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(data, &c.Info); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountries()", err.Error())
		return err
	}

	return nil
}

func scrollToCountriesSearch(ctx app.Context, e app.Event) {
	app.Window().ScrollToID("country-input")
}

func (c *Countries) OnCountryClick(ctx app.Context, e app.Event) {
	ctx.Navigate(fmt.Sprintf("/country/%s", ctx.JSSrc.Get("id").String()))
}
