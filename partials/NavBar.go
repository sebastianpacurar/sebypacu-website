package partials

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"strings"
)

type NavBar struct {
	app.Compo
	Menu
	SideMenu
}

type Menu struct {
	MenuItems []MenuItem
}

type MenuItem struct {
	Title string
	Link  string
}

func (n *NavBar) Render() app.UI {

	var testMenu = Menu{MenuItems: []MenuItem{
		{Title: "SideMenu", Link: "/"},
		{Title: "Home", Link: "/"},
		{Title: "countries", Link: "/countries/all"},
		{Title: "Quiz", Link: "/quiz-game"},
		{Title: "Mortăciuni!", Link: "/experiments"},
	}}

	return app.
		Nav().
		Body(
			app.
				Range(testMenu.MenuItems).Slice(func(i int) app.UI {
				currentItem := testMenu.MenuItems[i]
				itemContainerID := "/"
				if currentItem.Title != "Home" {
					itemContainerID = strings.ReplaceAll(strings.ToLower(currentItem.Title), " ", "-")
				}
				return app.
					If(app.Window().Get("innerWidth").Int() > 600,
						app.
							If(i == len(testMenu.MenuItems)-1,
								app.
									Div().
									Class("space-filler"),
								app.
									Div().
									Class("nav-item").
									Class("right").
									Body(
										app.
											A().
											ID(itemContainerID).
											Href(currentItem.Link).
											Text(currentItem.Title),
									),
							).
							ElseIf(currentItem.Title == "Home",
								app.
									Div().
									ID("home-icon").
									Body(
										app.
											A().
											Href(currentItem.Link).
											Body(
												app.
													Span().
													Class("navbar-icons").
													Class("material-icons").
													Text("home"),
											),
									),
							).
							ElseIf(currentItem.Title == "SideMenu",
								app.
									Div().
									ID("side-menu-toggle").
									Body(
										app.
											Button().
											Type("button").
											Body(
												app.
													Span().
													Class("navbar-icons").
													Class("material-icons").
													Text("menu"),
											),
									).OnClick(n.ToggleSideMenu),
							).
							Else(
								app.
									Div().
									Class("nav-item").
									Body(
										app.
											A().
											ID(itemContainerID).
											Href(currentItem.Link).
											Text(currentItem.Title),
									),
							),
					).Else()
			}),
		)
}

func (n *NavBar) ToggleSideMenu(ctx app.Context, e app.Event) {

}
