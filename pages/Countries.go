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
	IsTableView bool
}

type CountryData struct {
	Data []CountryInfo
}

func (c *Countries) OnNav(ctx app.Context) {
	if err := c.initCountries(ctx); err != nil {
		return
	}
	c.IsTableView = true
	c.Update()
}

func (c *Countries) Render() app.UI {
	return app.
		Div().
		Body(
			&components.Header{},
			&components.NavBar{},
			app.
				Main().
				Body(
					app.
						Form().
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
							app.
								Div().
								ID("table-layout-container").
								Body(
									app.
										Span().
										ID("table-layout-icon").
										Class("material-icons").
										Class("table-layout-view").
										Text("table_view").
										OnClick(c.SwitchCountriesView),
									app.
										Span().
										ID("cards-layout-icon").
										Class("material-icons").
										Class("table-layout-view").
										Text("grid_view").
										OnClick(c.SwitchCountriesView),
								),
						),
					app.If(len(c.Data) > 0,
						app.If(c.IsTableView,
							app.
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
												Range(c.Data).Slice(func(i int) app.UI {
												current := c.Data[i]
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
						).
							Else(
								app.
									Article().
									ID("cards-container").
									Body(
										app.Range(c.Data).Slice(func(i int) app.UI {
											current := c.Data[i]

											return app.Article().ID("country-card").Body(
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
																	OnClick(c.OnCountryClick),
															),
													),
											)

										}),
									),
							),
						&components.Footer{},
					).
						Else(
							&components.Spinner{},
						),
				),
		)
}

func (c *Countries) initCountries(ctx app.Context) error {

	data, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(data, &c.Data); err != nil {
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

func (c *Countries) SwitchCountriesView(ctx app.Context, e app.Event) {
	clickedIcon := ctx.JSSrc.Get("id").String()
	if c.IsTableView && clickedIcon == "cards-layout-icon" {
		c.IsTableView = false
	} else if !c.IsTableView && clickedIcon == "table-layout-icon" {
		c.IsTableView = true
	} else {
		return
	}
	c.Update()
}
