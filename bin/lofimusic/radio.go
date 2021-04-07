package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type radio struct {
	app.Compo

	lives        []liveRadio
	currentRadio string
}

func newRadio() *radio {
	return &radio{}
}

func (r *radio) OnMount(ctx app.Context) {
	r.lives = getLiveRadios()
	r.Update()
}

func (r *radio) OnNav(ctx app.Context) {
	r.currentRadio = strings.TrimPrefix(ctx.Page.URL().Path, "/")
	r.Update()
}

func (r *radio) Render() app.UI {
	return app.Main().
		Class("radio").
		Class("fill").
		Body(
			newYouTubePlayer(),
			app.Shell().
				Class("radio-shell").
				Menu(newNav().
					LiveRadios(r.lives).
					CurrentRadio(r.currentRadio)).
				OverlayMenu(newNav().
					LiveRadios(r.lives).
					CurrentRadio(r.currentRadio)),
		)
}
