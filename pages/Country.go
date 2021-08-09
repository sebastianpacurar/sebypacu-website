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

type CountryPage struct {
	app.Compo
	AccessDetails CountryDetails
}

type CountryDetails struct {
	CountryName    string            `json:"name"`
	CountryFlag    string            `json:"flag"`
	CountryCapital string            `json:"capital"`
	CountryRegion  string            `json:"region"`
	Population     int               `json:"population"`
	Code           string            `json:"code"`
	CurrencyName   string            `json:"currency-name"`
	Symbol         string            `json:"symbol"`
	Translations   map[string]string `json:"translations"`
}

func (cp *CountryPage) OnMount (ctx app.Context) {
	country := strings.ReplaceAll(ctx.Page.URL().Path, "/location-details/name/", "")
	cp.initCountry(ctx, country)
	cp.Update()
}

func (cp *CountryPage) Render() app.UI {

	return app.Div().
		Body(
			&components.Header{},
			app.Div().
				Class("card mb-3").
				Body(
					app.Img().
						Class("card-img-top").
						Src(cp.AccessDetails.CountryFlag).
						Alt("Flag not found"),
					app.Div().
						Class("card-body").
						Body(
							app.H5().
								Class("card-title").
								Text(cp.AccessDetails.CountryName),
							app.Div().
								Class("container"),
							//Body(
							//	app.Range(cp.AccessDetails).Map(func(k string) app.UI {
							//		return app.Div().
							//			Class("row").
							//			Body(
							//				app.Div().Text(k),
							//				app.Div().
							//					Class("col").
							//					Body(
							//						app.P().
							//							Class("card-text").
							//							Text(k),
							//					),
							//				app.Div().
							//					Class("col").
							//					Body(
							//						app.P().
							//							Class("card-text").
							//							Text(parsedDetails[k]),
							//					),
							//			)
							//	}),
							//),
						),
				))
}

func (cp *CountryPage) initCountry(ctx app.Context, country string) {
	if err := json.Unmarshal(API.FetchData(fmt.Sprintf("name/%s", country)), &cp.AccessDetails); err != nil {
		log.Fatalln(err.Error())
	}
	cp.Update()
}
