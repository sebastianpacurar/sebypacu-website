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

	return app.Nav().
		Class("navbar").
		Class("navbar-expand-lg").
		Class("navbar-light").
		Style("background-color", "lightseagreen").
		Style("margin-bottom", "15").
		Body(
			app.Ul().
				Class("navbar-nav").
				Body(
					app.Li().
						Class("nav-item").
						Body(
							app.A().
								Class("nav-link").
								Class("link-dark").
								Href("/").
								Body(
									app.Span().
										Class("material-icons").
										Class("md-48").
										Text("home"),
								),
						),

					app.Range(testMenu.Items).Slice(func(i int) app.UI {
						return app.Li().
							Class("nav-item").
							Body(
								app.A().
									Class("link-dark").
									Class("nav-link").
									Href(testMenu.Items[i].Link).
									Text(testMenu.Items[i].Title),
							)
					}),
				),
		)

}
