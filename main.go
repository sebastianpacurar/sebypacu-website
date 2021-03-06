package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
	comp "pwa/components"
	"pwa/components/countries"
	"pwa/experiments"
)

func main() {
	app.Route("/", &comp.Home{})
	app.Route("/quiz-game", &comp.QuizGame{})
	app.Route("/experiments", &experiments.Mortaciuni{})

	app.Route("/countries/all", &countries.Page{})
	app.RouteWithRegexp("^/countries/alpha2/.*", &countries.Country{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Title: "1 fel de 2 feluri",
		Name:  "Wasm Fun",
		RawHeaders: []string{
			`<!--Google Material Icons-->
			<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">`,
		},
		Styles: []string{
			"/web/style/Header.css",
			"/web/style/Navbar.css",
			"/web/style/Spinner.css",
			"/web/style/CountriesPage.css",
			//"/web/style/Footer.css",
		},
	})

	if err := http.ListenAndServe(":7005", nil); err != nil {
		log.Fatal(err)
	}
}
