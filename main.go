package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
	"pwa/experiments"
	page "pwa/pages"
)

func main() {
	app.Route("/", &page.Home{})
	app.Route("/countries", &page.Countries{})
	app.Route("/quiz-game", &page.QuizGame{})
	app.Route("/experiments", &experiments.Mortaciuni{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Title: "1 fel de 2 feluri",
		Name: "Wasm Fun",
		RawHeaders: []string{
			`<!--Google Material Icons-->
			<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">`,
		},
		Styles: []string{
			"/web/style/styles.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
