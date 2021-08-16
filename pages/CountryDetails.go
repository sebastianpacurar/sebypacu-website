package pages

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"pwa/API"
	"pwa/components"
	"strings"
)

type CountryDetails struct {
	app.Compo
	Details []CountryInfo
}

type CountryInfo struct {
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

func (cd *CountryDetails) OnNav(ctx app.Context) {
	if err := cd.initCountry(ctx); err != nil {
		return
	}
	cd.Update()
}

func (cd *CountryDetails) Render() app.UI {

	return app.
		Div().
		Body(
			&components.Header{},
			&components.NavBar{},
			app.
				Ul().
				Body(
					app.Range(cd.Details).Slice(func(i int) app.UI {
						country := cd.Details[i]
						return app.
							Div().
							Text(country.Area)
					}),
				),
		)
}

func (cd *CountryDetails) initCountry(ctx app.Context) error {
	code := strings.Split(ctx.Page.URL().String(), "/")
	res, err := API.FetchCountries(fmt.Sprintf("alpha?codes=%s", code[len(code)-1]))
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(res, &cd.Details); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountries()", err.Error())
		return err
	}
	return nil
}
