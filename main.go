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

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name: "Pure By Go",
		Styles: []string{
			"web/styles/css/bootstrap.css",
			"web/styles/css/bootstrap-grid.css",
			"web/styles/css/bootstrap-reboot.css",
			"web/styles/css/bootstrap-utilities.css",
		},
		Scripts: []string{
			"web/styles/js/bootstrap.js",
		},
		Icon: app.Icon{
			Default:    "",
			Large:      "",
			AppleTouch: "",
		},

		Title: "go+wasm=rulz",
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
