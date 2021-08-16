package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"strings"
)

type NavBar struct {
	app.Compo
	Menu
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
		{Title: "Home", Link: "/"},
		{Title: "Countries", Link: "/countries"},
		{Title: "Quiz Game", Link: "/quiz-game"},
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
							Class("nav-item").
							Body(
								app.
									A().
									Href(currentItem.Link).
									Body(
										app.
											Span().
											ID("nav-bar-home-icon").
											Class("material-icons").
											Text("home"),
									),
							),
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
					)
			}),
		)
}
