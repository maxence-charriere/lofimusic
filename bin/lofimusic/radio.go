package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	menuWidth = 282
)

type radio struct {
	app.Compo

	lives             []liveRadio
	current           liveRadio
	isPlaying         bool
	isUpdateAvailable bool
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

	r.isUpdateAvailable = ctx.AppUpdateAvailable
	r.isPlaying = false
	r.Update()
}

func (r *radio) OnAppUpdate(ctx app.Context) {
	r.isUpdateAvailable = ctx.AppUpdateAvailable
	r.Update()
}

func (r *radio) Render() app.UI {
	return app.Main().
		Class("radio").
		Class("fill").
		Body(
			newYouTubePlayer().
				Radio(r.current).
				OnPlaybackChange(r.onPlaybackChange),
			app.Shell().
				Class("radio-shell").
				MenuWidth(menuWidth).
				Menu(newNav().
					LiveRadios(r.lives).
					CurrentRadio(r.current)).
				OverlayMenu(newNav().
					Class("background-overlay").
					LiveRadios(r.lives).
					CurrentRadio(r.current)).
				Content(
					app.Aside().
						Class("app-title").
						Class("hspace-out").
						Body(
							app.Stack().
								Class("fit").
								Class("vspace-stretch").
								Class("right").
								Center().
								Content(
									app.If(r.isUpdateAvailable,
										newLink().
											Class("link-update").
											Class("glow").
											Label("Update").
											Icon(newSVGIcon().RawSVG(downloadSVG)).
											OnClick(r.onUpdateClick),
									),
								),
						),
					app.Div().
						Class("hspace-out").
						Class("vspace-content").
						Body(
							newInfo().
								Radio(r.current).
								Playing(r.isPlaying),
						),
				),
		)
}

func (r *radio) onPlaybackChange(ctx app.Context, isPlaying bool) {
	r.isPlaying = isPlaying
	r.Update()
}

func (r *radio) onUpdateClick(ctx app.Context) {
	ctx.Reload()
}
