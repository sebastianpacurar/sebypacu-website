package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	page "go-app/pages"
	"log"
	"net/http"
)

var links = map[string]app.Composer{
	"/countries": &page.Countries{},
	"/second":    &page.Second{},
	"/third":     &page.Third{},
	"/fourth":    &page.Fourth{},
}

func main() {
	app.Route("/", &page.Home{})

	for link, composer := range links {
		app.Route(link, composer)
	}

	// Match all countries routes to avoid looping 150 times
	app.RouteWithRegexp("^/location-details/.*", &page.CountryPage{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name: "Pure By Go",
		RawHeaders: []string{
			`<!--Bootstrap 4.6 CSS + JS-->
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css" integrity="sha384-B0vP5xmATw1+K9KRQjQERJvTumQW0nPEzvF6L/Z6nronJ3oUOFUFpCjEUQouq2+l" crossorigin="anonymous">
			<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-Piv4xVNRyMGpqkS2by6br4gNJ7DXjqk09RmUpJ8jgGtD7zP9yug3goQfGII0yAns" crossorigin="anonymous"></script>

			<!--Google Material Icons-->
			<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		`,
		},
		Title: "go+wasm=rulz",
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
