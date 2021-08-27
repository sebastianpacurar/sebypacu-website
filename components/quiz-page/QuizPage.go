package quiz_page

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"pwa/partials"
)

type QuizPage struct {
	app.Compo
	Question      string
	CorrectAnswer string
	WrongAnswers  []string
}

func (qp *QuizPage) Render() app.UI {
	return app.Div().Body(
		&partials.Header{},
		&partials.NavBar{},
		&partials.Footer{},
	)
}
