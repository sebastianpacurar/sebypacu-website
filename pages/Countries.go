package pages

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"go-app/components"
	"go-app/pages/API"
	"log"
	"strings"
)

type Countries struct {
	app.Compo
	CountriesList []CountryTable
	ActivePage    string
}

type CountryTable struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Region  string `json:"region"`
	Flag    string `json:"flag"`
}

func (c *Countries) Render() app.UI {

	return app.Div().Body(
		&components.Header{},
		app.Button().Class("test").Text("generate").OnClick(c.generateCountries),

		app.If(len(c.CountriesList) > 0,

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
							app.Range(c.CountriesList).Slice(func(i int) app.UI {
								country := c.CountriesList[i]
								countryLink := strings.ReplaceAll(country.Name, " ", "-")
								return app.Tr().
									Body(
										app.Th().
											Scope("row").
											Text(i),
										app.Td().
											Body(
												app.Button().
													Type("button").
													Class("btn").
													Class("btn-info").
													Text(country.Name).
													OnClick(c.OnCountryClick).
													ID(countryLink),
											),
										app.Td().
											Body(
												app.Img().
													Class("img-thumbnail").
													Class("img-rounded").
													Src(country.Flag).
													Alt(country.Name+" img").
													Style("width", "150px"),
											),
										app.Td().
											Text(country.Capital),
										app.Td().
											Text(country.Region),
									)
							}),
						),
				),
		).Else(
			app.Div().
				Text("not displayed"),
		),
	)
}

func (c *Countries) OnCountryClick(ctx app.Context, e app.Event) {
	c.ActivePage = ctx.JSSrc.Get("id").String()
	ctx.Navigate(fmt.Sprintf("/location-details/name/%s", c.ActivePage))
}

func (c *Countries) generateCountries(ctx app.Context, e app.Event) {
	c.initCountries(ctx)
}

func (c *Countries) initCountries(ctx app.Context) {
	if err := json.Unmarshal(API.FetchData("all"), &c.CountriesList); err != nil {
		log.Fatalln(err.Error())
	}

	c.Update()
}
