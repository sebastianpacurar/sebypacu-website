package partials

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type SideMenu struct {
	app.Compo
	IsOpen  bool
	Content []interface{}
}

func (sm *SideMenu) Render() app.UI {
	return app.
		Aside().
		ID("sebypacu-sidemenu").
		Body(
			app.Ul().Body(
				app.Range(sm.Content).Slice(func(i int) app.UI {
					current := structs.Map(sm.Content[i])
					fmt.Println(current)
					return app.Range(current).Map(func(k string) app.UI {
						val := fmt.Sprintf("%s: %v", k, current[k])

						return app.P().Text(val)
					})
				}),
			),
		)
}
