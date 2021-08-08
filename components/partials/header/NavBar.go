package header

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type NavBar struct {
	app.Compo
}

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Title string
	Link  string
}

func (n *NavBar) Render() app.UI {

	var testMenu = Menu{Items: []MenuItem{
		{Title: "Countries", Link: "/countries"},
		{Title: "Second", Link: "/second"},
		{Title: "Third", Link: "/third"},
		{Title: "Fourth", Link: "/fourth"},
	}}

	return app.Ul().Class("nav justify-content-center").Body(
		app.Range(testMenu.Items).Slice(func(i int) app.UI {
			return app.Li().Class("nav-item").Body(
				app.A().
					Class("nav-link").
					Href(testMenu.Items[i].Link).
					Text(testMenu.Items[i].Title),
			)
		}),
	)
}


