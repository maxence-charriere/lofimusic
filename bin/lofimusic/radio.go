package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type radio struct {
	app.Compo

	lives   []liveRadio
	current liveRadio
}

func newRadio() *radio {
	return &radio{}
}

func (r *radio) OnMount(ctx app.Context) {
	r.lives = getLiveRadios()
	r.Update()
}

func (r *radio) OnNav(ctx app.Context) {
	slug := strings.TrimPrefix(ctx.Page.URL().Path, "/")

	for _, lr := range r.lives {
		if slug == lr.Slug {
			r.current = lr
			break
		}
	}
	r.Update()
}

func (r *radio) Render() app.UI {
	return app.Main().
		Class("radio").
		Class("fill").
		Body(
			newYouTubePlayer().Radio(r.current),
			app.Shell().
				Class("radio-shell").
				Menu(newNav().
					LiveRadios(r.lives).
					CurrentRadio(r.current)).
				OverlayMenu(newNav().
					LiveRadios(r.lives).
					CurrentRadio(r.current)),
		)
}
