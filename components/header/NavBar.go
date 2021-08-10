package header

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"strings"
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
		{Title: "countries", Link: "/countries"},
	}}

	return app.Nav().
		Class("navbar").
		Class("navbar-expand-lg").
		Class("navbar-light").
		Style("background-color", "lightseagreen").
		Style("margin-bottom", "8px").
		Body(
			app.Ul().
				Class("nav").
				Class("nav-pills").
				Class("nav-fill").
				Body(
					app.Li().
						Class("nav-item").
						Body(
							app.A().
								Class("nav-link").
								Class("active").
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
									ID(fmt.Sprintf("nav-link-%s", testMenu.Items[i].Title)).
									Class("link-dark").
									Class("nav-link").
									Href(testMenu.Items[i].Link).
									Text(testMenu.Items[i].Title),
							)
					}),

					/// Check to see if Countries link, contains the class "active". If true, then render Countries Search Bar
					app.If(strings.Contains(app.Window().GetElementByID("nav-link-countries").Get("class").String(), "active"),
						app.Form().
							Class("form-inline").
							Body(
								app.Input().
									Type("search").
									Class("form-control").
									Class("mr-sm-2").
									Placeholder("Search Country").
									Aria("label", "Search"),
							),
					),
				),
		)

}

func (n *NavBar) OnNavLinkClick(ctx app.Context, e app.Event) {
	ctx.JSSrc.Set("class", "active")
	n.Update()
}
