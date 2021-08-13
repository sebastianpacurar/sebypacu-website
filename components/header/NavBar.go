package header

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type NavBar struct {
	app.Compo

	IsSearchBarDisplayed bool
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
		{Title: "Quiz Game", Link: "/universal-quiz-game"},
	}}

	return app.Nav().
		Class("navbar").
		Class("navbar-expand-lg").
		Class("navbar-light").
		Style("background-color", "#ff5F1f").
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
							app.Button().
								Type("button").
								ID("home-icon").
								Class("nav-link").
								Style("color", "white").
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
							Body(app.If(testMenu.Items[0].Title == "Quiz Game",
								app.Button().
									Type("button").
									ID(fmt.Sprintf("%s-nav-link", testMenu.Items[i].Title)).
									Class("btn").
									Class("btn-primary").
									Class("btn-large").
									Name(testMenu.Items[i].Title).
									Text(testMenu.Items[i].Title).
									Style("color", "white").
									Disabled(true),
							).Else(
								app.Button().
									Type("button").
									ID(fmt.Sprintf("%s-nav-link", testMenu.Items[i].Title)).
									Class("btn").
									Class("btn-primary").
									Class("btn-large").
									Name(testMenu.Items[i].Title).
									Text(testMenu.Items[i].Title).
									Style("color", "white").
									OnClick(n.onUpdateClick),
							))
					}),

					/// If update is true
					app.If(n.IsSearchBarDisplayed,
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

func (n *NavBar) OnAppUpdate(ctx app.Context, e app.Event) {
	n.IsSearchBarDisplayed = ctx.AppUpdateAvailable
	n.Update()
}

func (n *NavBar) onUpdateClick(ctx app.Context, e app.Event) {
	n.IsSearchBarDisplayed = !n.IsSearchBarDisplayed

	ctx.Navigate(fmt.Sprintf("localhost:9000/%s", ctx.JSSrc.Get("name").String()))
}
