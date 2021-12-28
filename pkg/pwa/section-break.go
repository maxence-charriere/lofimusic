package pwa

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type sectionBreak struct {
	app.Compo
}

func newSectionBreak() *sectionBreak {
	return &sectionBreak{}
}

func (s *sectionBreak) Render() app.UI {
	return app.Div().Class("section-break")
}
