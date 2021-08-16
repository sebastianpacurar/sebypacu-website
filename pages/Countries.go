package pages

import (
	"encoding/json"
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
	Name         string   `json:"name"`
	Capital      string   `json:"capital"`
	Region       string   `json:"region"`
	Flag         string   `json:"flag"`
	AltSpellings []string `json:"altSpellings"`
	Subregion    string   `json:"subregion"`
	//Translations   map[string]string
	Population     int32     `json:"population"`
	LatLng         []float32 `json:"latlng"`
	Demonym        string    `json:"demonym"`
	Area           float32   `json:"area"`
	Gini           float32   `json:"gini"`
	Timezones      []string  `json:"timezones"`
	Borders        []string  `json:"borders"`
	NativeName     string    `json:"nativeName"`
	CallingCodes   []string  `json:"callingCodes"`
	NumericCode    string    `json:"numericCode"`
	TopLevelDomain []string  `json:"topLevelDomain"`
	Alpha2Code     string    `json:"alpha2code"`
	Alpha3Code     string    `json:"alpha3code"`
}

type Currencies struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Languages struct {
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
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
						Table().
						ID("countries").
						Body(
							app.
								Caption().
								ID("table-title").
								Text("List of All Countries and independent Islands, Regions"),
							app.
								THead().
								Body(
									app.Tr().
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
											)
									}),
								),
							app.
								Tfoot().
								Body(
									app.Tr().Body(
										app.
											Td().
											ColSpan(4).
											Text("Go to Top"),
									),
								).OnClick(scrollToTop),
						),
					&components.Footer{},
				).Else(
				&components.Spinner{},
			),
		)
}

func (c *Countries) initCountries(ctx app.Context) error {
	res, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(res, &c.Info); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountries()", err.Error())
		return err
	}
	return nil
}

func scrollToTop(ctx app.Context, e app.Event) {
	app.Window().ScrollToID("table-title")
}
