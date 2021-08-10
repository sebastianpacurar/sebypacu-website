package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	page "go-app/pages"
	"log"
	"net/http"
)

func main() {
	app.Route("/", &page.Home{})

	app.Route("/countries", &page.Countries{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name: "Pure By Go",
		RawHeaders: []string{
			`<!--Bootstrap 4.6 CSS-->
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css" integrity="sha384-B0vP5xmATw1+K9KRQjQERJvTumQW0nPEzvF6L/Z6nronJ3oUOFUFpCjEUQouq2+l" crossorigin="anonymous">

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
