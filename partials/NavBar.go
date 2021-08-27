package partials

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"strings"
)

type NavBar struct {
	app.Compo
	Menu
	SideMenu
	layoutNav map[string]bool
}

type Menu struct {
	MenuItems []MenuItem
}

type MenuItem struct {
	Title    string
	Link     string
	IsActive bool // will implement in the future
}

func (n *NavBar) Render() app.UI {

	var menu = Menu{MenuItems: []MenuItem{
		{Title: "SideMenu", Link: "", IsActive: false},
		{Title: "Home", Link: "/", IsActive: false},
		{Title: "Geography", Link: "/geography/all", IsActive: false},
		{Title: "Quiz", Link: "/quiz", IsActive: false},
		{Title: "Mortăciuni!", Link: "/experiments", IsActive: false},
	}}

	return app.
		Nav().
		Body(
			app.
				Range(menu.MenuItems).Slice(func(i int) app.UI {
				currentItem := menu.MenuItems[i]
				itemContainerID := "/"
				if currentItem.Title != "Home" {
					itemContainerID = strings.ReplaceAll(strings.ToLower(currentItem.Title), " ", "-")
				}
				return app.If(app.Window().Get("innerWidth").Int() > 600,
					app.
						If(i == len(menu.MenuItems)-1,
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
								ID("home-page-icon").
								Body(
									app.
										A().
										Href(currentItem.Link).
										Body(
											app.
												Span().
												Class("navbar-icons").
												Class("material-icons").
												Text("home-page"),
										),
								),
						).
						ElseIf(currentItem.Title == "SideMenu",
							app.
								Div().
								ID("sm").
								Class("sm-hidden").
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
				).Else(
					&SideMenu{},
				)
			}),
		)
}

func (n *NavBar) ToggleSideMenu(ctx app.Context, e app.Event) {
	n.SideMenu.IsOpen = !n.SideMenu.IsOpen
	n.Update()
}
