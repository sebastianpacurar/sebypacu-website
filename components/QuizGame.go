package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/partials"
)

type QuizGame struct {
	app.Compo
	Question      string
	CorrectAnswer string
	WrongAnswers  []string
}

func (qg *QuizGame) Render() app.UI {
	return app.Div().Body(
		&partials.Header{},
		&partials.NavBar{},
		&partials.Footer{},
	)
}
