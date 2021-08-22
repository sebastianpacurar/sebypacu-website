package countries

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"pwa/API"
	"pwa/partials"
	"strings"
)

type Page struct {
	app.Compo
	Country
	Grid
	layout map[string]bool
}

func (p *Page) OnNav(ctx app.Context) {
	if err := p.initPage(ctx); err != nil {
		return
	}
	p.Update()
}

func (p *Page) Render() app.UI {

	/// set default layout to "table"
	if len(p.layout) == 0 {
		p.layout = make(map[string]bool)
		p.layout["table"] = true
		p.layout["cards"] = false
	}

	return app.
		Div().
		Body(
			&partials.Header{},
			&partials.NavBar{},
			&partials.SideMenu{},
			app.
				Main().ID("countries-main").
				Body(
					app.
						Form().
						Body(
							app.
								Div().
								ID("country-search-container").
								Body(
									app.
										Input().
										ID("country-input").
										Type("text").
										Placeholder("Filter Page by Letters"),
									app.
										Button().
										Type("submit").
										Value("Submit").
										Text("Fetch!"),
								),
							app.
								Div().
								ID("layout-toggle-container").
								Body(
									app.
										Span().
										ID("table-layout-icon").
										Class("material-icons").
										Class("table-layout-view").
										Text("table_view").
										OnClick(p.switchPageView),
									app.
										Span().
										ID("cards-layout-icon").
										Class("material-icons").
										Class("table-layout-view").
										Text("grid_view").
										OnClick(p.switchPageView),
								),
						),
					app.
						If(len(p.Details) > 0,
							app.
								If(p.layout["table"],
									&Table{
										data: p.Details,
									},
								).
								ElseIf(p.layout["cards"],
									app.If(p.Mounted(),
										&Grid{
											data: p.Details,
										},
									),
								),
							&partials.Footer{},
						).
						Else(
							&partials.Spinner{},
						),
				),
		)
}

func (p *Page) initPage(ctx app.Context) error {

	data, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(data, &p.Details); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initPage()", err.Error())
		return err
	}

	return nil
}

func NavigateToCountry(ctx app.Context, e app.Event) {
	ctx.Navigate(fmt.Sprintf("/countries/alpha2/%s", ctx.JSSrc.Get("id").String()))
}

// SwitchPageView
/// Set "active" icon class to the icon which has the key set to true.
/// "clickedIcon" can be either 'table' or 'cards' to match the key names.
func (p *Page) switchPageView(ctx app.Context, e app.Event) {
	clickedIcon := strings.Split(ctx.JSSrc.Get("id").String(), "-")[0]

	if !p.layout[clickedIcon] {
		for k := range p.layout {
			p.layout[k] = !p.layout[k]
		}
	} else {
		return
	}

	p.Update()
}
