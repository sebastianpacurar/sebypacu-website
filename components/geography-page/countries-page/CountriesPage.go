package countries_page

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"pwa/API"
	"pwa/partials"
	"strings"
)

type CountriesPage struct {
	app.Compo
	Country
	Grid
	layout map[string]bool
}

func (cp *CountriesPage) OnNav(ctx app.Context) {
	if err := cp.initCountriesPage(ctx); err != nil {
		return
	}
	cp.Update()
}

func (cp *CountriesPage) Render() app.UI {

	/// set default layout to "table"
	if len(cp.layout) == 0 {
		cp.layout = make(map[string]bool)
		cp.layout["table"] = true
		cp.layout["cards"] = false
	}

	return app.
		Div().
		Body(
			&partials.Header{},
			&partials.NavBar{},
			app.
				Main().ID("countries-page").
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
										Placeholder("Filter CountriesPage by Letters"),
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
										OnClick(cp.switchCountriesPageView),
									app.
										Span().
										ID("cards-layout-icon").
										Class("material-icons").
										Class("table-layout-view").
										Text("grid_view").
										OnClick(cp.switchCountriesPageView),
								),
						),
					app.
						If(len(cp.Details) > 0,
							app.
								If(cp.layout["table"],
									&Table{
										data: cp.Details,
									},
								).
								ElseIf(cp.layout["cards"],
									app.If(cp.Mounted(),
										&Grid{
											data: cp.Details,
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

func (cp *CountriesPage) initCountriesPage(ctx app.Context) error {

	data, err := API.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(data, &cp.Details); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountriesPage()", err.Error())
		return err
	}

	return nil
}

func NavigateToCountry(ctx app.Context, e app.Event) {
	ctx.Navigate(fmt.Sprintf("/geography/country/alpha2/%s", ctx.JSSrc.Get("id").String()))
}

// SwitchCountriesPageView
/// Set "active" icon class to the icon which has the key set to true.
/// "clickedIcon" can be either 'table' or 'cards' to match the key names.
func (cp *CountriesPage) switchCountriesPageView(ctx app.Context, e app.Event) {
	clickedIcon := strings.Split(ctx.JSSrc.Get("id").String(), "-")[0]

	if !cp.layout[clickedIcon] {
		for k := range cp.layout {
			cp.layout[k] = !cp.layout[k]
		}
	} else {
		return
	}

	cp.Update()
}
