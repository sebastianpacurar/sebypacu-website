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
	Items []MenuItem
}

type MenuItem struct {
	Title  string
	Link   string
	Active bool
}

func (n *NavBar) Render() app.UI {

	var testMenu = Menu{Items: []MenuItem{
		{Title: "Home", Link: "/", Active: true},
		{Title: "Countries", Link: "/countries", Active: false},
		{Title: "Quiz Game", Link: "/quiz-game", Active: false},
		{Title: "Mortăciuni!", Link: "/experiments", Active: false},
	}}

	return app.
		Nav().
		Body(
			app.
				Range(testMenu.Items).Slice(func(i int) app.UI {
				currentItem := testMenu.Items[i]
				linkId := "/"
				if currentItem.Title != "Home" {
					linkId = strings.ReplaceAll(strings.ToLower(currentItem.Title), " ", "-")
				}
				return app.
					If(i == len(testMenu.Items)-1,
						app.
							Div().
							Class("space-filler"),
						app.
							Div().
							Class("right").
							Body(
								app.
									A().
									ID(linkId).
									Href(currentItem.Link).
									Text(currentItem.Title),
							),
					).Else(
					app.
						If(linkId == "/" && currentItem.Active,
							app.
								Div().
								Class("active").
								Body(
									app.
										A().
										ID("home-icon").
										Href("/").
										Body(
											app.
												Span().
												Class("material-icons").
												Class("md-48").
												Text("home"),
										),
								),
						).
						ElseIf(linkId == "/" && !currentItem.Active,
							app.
								Div().
								Body(
									app.A().
										ID("home-icon").
										Href("/").
										Body(
											app.
												Span().
												Class("material-icons").
												Class("md-48").
												Text("home"),
										),
								),
						).
						ElseIf(linkId != "/" && currentItem.Active,
							app.
								Div().
								Class("active").
								Body(
									app.A().
										ID(linkId).
										Href(currentItem.Link).
										Text(currentItem.Title),
								),
						).
						Else(
							app.
								Div().
								Body(
									app.A().
										ID(linkId).
										Href(currentItem.Link).
										Text(currentItem.Title),
								),
						),
				)

			}),
		)
}
