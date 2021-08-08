package pages

import (
	"encoding/json"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"go-app/components"
	"io/ioutil"
	"log"
	"net/http"
)

type Countries struct {
	app.Compo
	CountriesList []Country
}

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Region  string `json:"region"`
}

func (c *Countries) Render() app.UI {
	return app.Div().Body(
		&components.Header{},
		app.Button().Class("test").Text("generate").OnClick(c.onClick),

		app.If(len(c.CountriesList) > 0,

			app.Table().Class("table").Body(
				app.THead().Class("thead-dark").Body(
					app.Tr().Body(
						app.Th().Scope("col").Text("#"),
						app.Th().Scope("col").Text("Name"),
						app.Th().Scope("col").Text("Capital"),
						app.Th().Scope("col").Text("Region"),
					),
				),
				app.TBody().Body(
					app.Range(c.CountriesList).Slice(func(i int) app.UI {
						return app.Tr().Body(
							app.Th().Scope("row").Text(i),
							app.Th().Text(c.CountriesList[i].Name),
							app.Th().Text(c.CountriesList[i].Capital),
							app.Th().Text(c.CountriesList[i].Region),
						)
					}),
				),
			),
		).Else(
			app.Div().Text("not displayed"),
		),
	)
}

func (c *Countries) onClick(ctx app.Context, e app.Event) {
	c.initCountries(ctx)
}

func (c *Countries) initCountries(ctx app.Context) {
	resp, err := http.Get("https://restcountries.eu/rest/v2/all")
	if err != nil {
		log.Fatalln(err)
	}

	////We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//Convert the body to type string
	if err := json.Unmarshal(body, &c.CountriesList); err != nil {
		log.Fatalln(err.Error())
	}

	c.Update()
}
