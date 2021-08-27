package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
	"pwa/components/geography-page"
	"pwa/components/geography-page/countries-page"
	comp "pwa/components/home-page"
	"pwa/components/quiz-page"
	"pwa/experiments"
)

func main() {
	app.Route("/", &comp.HomePage{})
	app.Route("/quiz", &quiz_page.QuizPage{})
	app.Route("/experiments", &experiments.Mortaciuni{})

	app.Route("/geography/all", &geography_page.GeographyPage{})
	app.Route("/geography/countries/all", &countries_page.CountriesPage{})
	app.RouteWithRegexp("^/geography/country/alpha2code/.*", &countries_page.Country{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Title:  "1 fel de 2 feluri",
		Name:   "Diverse aplicatii de kkt \\:D/",
		Author: "Sebastian Pacurar",
		RawHeaders: []string{
			`<!--Google Material Icons-->
			<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">`,
		},
		Styles: []string{
			"/web/style/components/geography/countries/CountriesPage.css",
			"/web/style/components/geography/countries/CardsView.css",
			"/web/style/components/geography/countries/TableView.css",
			"/web/style/partials/Header.css",
			"/web/style/partials/NavBar.css",
			"/web/style/partials/Spinner.css",
		},
	})

	if err := http.ListenAndServe(":7015", nil); err != nil {
		log.Fatal(err)
	}
}
