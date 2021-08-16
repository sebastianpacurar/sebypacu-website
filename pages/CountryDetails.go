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
	Name           string            `json:"name"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2code"`
	Alpha3Code     string            `json:"alpha3code"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Population     int32             `json:"population"`
	Flag           string            `json:"flag"`
	LatLng         []float64         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Area           float64           `json:"area"`
	Gini           float64           `json:"gini"`
	Timezones      []string          `json:"timezones"`
	Borders        []string          `json:"borders"`
	NativeName     string            `json:"nativeName"`
	NumericCode    string            `json:"numericCode"`
	Currencies     []Currency        `json:"currencies"`
	Languages      []Language        `json:"languages"`
	Translations   map[string]string `json:"translations"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Language struct {
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

func (cd *CountryDetails) OnNav(ctx app.Context) {
	if err := cd.initCountry(ctx); err != nil {
		log.Fatalln("Eroare la navigarea pe tara, ", err.Error())
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
				If(len(cd.Details) > 0,
					app.
						Div().
						Body(
							app.Range(cd.Details).Slice(func(i int) app.UI {
								country := cd.Details[i]
								return app.
									Ul().
									Body(
										app.
											Li().
											Text(country.Name),
										app.
											Li().
											Text(country.TopLevelDomain),
										app.
											Li().
											Text(country.Alpha2Code),
										app.
											Li().
											Text(country.Alpha3Code),
										app.
											Li().
											Text(country.CallingCodes),
										app.
											Li().
											Text(country.Subregion),
										app.
											Li().
											Text(country.Population),
										app.
											Li().
											Text(country.Flag),
										app.
											Li().
											Text(country.LatLng),
										app.
											Li().
											Text(country.Demonym),
										app.
											Li().
											Text(country.Gini),
										app.
											Li().
											Text(country.Timezones),
										app.
											Li().
											Text(country.Borders),
										app.
											Li().
											Text(country.NativeName),
										app.
											Li().
											Text(country.NumericCode),

										app.
											Li().
											Body(
												app.Range(country.Currencies).Slice(func(i int) app.UI {
													currency := country.Currencies[i]
													return app.Ul().Body(
														app.Li().Text(currency.Name),
														app.Li().Text(currency.Code),
														app.Li().Text(currency.Symbol),
													)

												}),
											),
										app.
											Li().
											Body(
												app.Range(country.Languages).Slice(func(i int) app.UI {
													language := country.Languages[i]
													return app.
														Ul().
														Body(
															app.Li().Text(language.Name),
															app.Li().Text(language.NativeName),
														)
												}),
											),
										app.
											Li().
											Body(
												app.Range(country.Translations).Map(func(k string) app.UI {
													return app.
														Ul().
														Body(
															app.Li().Text(country.Translations[k]),
														)
												}),
											),
									)
							}),
						),
					&components.Footer{},
				).Else(
				&components.Spinner{},
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
