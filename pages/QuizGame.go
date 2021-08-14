package pages

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/components"
)

type QuizGame struct {
	app.Compo
	Question      string
	CorrectAnswer string
	WrongAnswers  []string
}

func (qg *QuizGame) Render() app.UI {
	return app.Div().Body(
		&components.Header{},
		&components.Footer{},
	)
}
