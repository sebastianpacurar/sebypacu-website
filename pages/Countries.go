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
	Data           struct {
		Children []struct {
			Data Countries
		}
	}
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
	c.initCountries(ctx)
	c.Update()
}

func (c *Countries) Render() app.UI {

	return app.Div().Body(
		&components.Header{},
		&components.NavBar{},
		//app.Button().Class("test").Text("generate").OnClick(c.InitCountries),

		app.If(len(c.Info) > 0,

			app.Table().
				Class("table").
				Class("table-sm").
				Class("table-hover").
				Class("table-bordered").
				Class("border-info").
				Class("table-stripped").
				Class("caption-top").
				Body(
					app.Caption().
						Text("List of All Countries and independent Islands, Regions"),
					app.THead().
						Class("thead-light").
						Body(
							app.Tr().
								Body(
									app.Th().
										Scope("col").
										Text("#").
										Style("width", "20%"),
									app.Th().
										Scope("col").
										Text("Name").
										Style("width", "20%"),
									app.Th().
										Scope("col").
										Text("Flag").
										Style("width", "20%"),
									app.Th().
										Scope("col").
										Text("Capital").
										Style("width", "20%"),
									app.Th().
										Scope("col").
										Text("Region").
										Style("width", "20%"),
								),
						),
					app.TBody().
						Body(
							app.Range(c.Info).Slice(func(i int) app.UI {
								current := c.Info[i]

								return app.Tr().
									Body(
										app.Th().
											Scope("row").
											Text(i+1),
										app.Td().
											Body(
												app.P().
													Class("h4").
													Text(current.Name),
											),
										app.Td().
											Body(
												app.Img().
													Class("img-thumbnail").
													Class("img-rounded").
													Src(current.Flag).
													Alt(current.Name).
													Style("width", "150px"),
											),
										app.Td().
											Body(
												app.P().
													Class("h4").
													Text(current.Capital),
											),
										app.Td().
											Body(
												app.P().
													Class("h4").
													Text(current.Region),
											),
									)
							}),
						),
				),
		).Else(
			app.Div().
				Text("not displayed"),
		),
		&components.Footer{},
	)
}

func (c *Countries) initCountries(ctx app.Context) {
	res, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := json.Unmarshal(res, &c.Info); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountries()", err.Error())
	}
}
