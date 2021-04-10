package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"golang.org/x/exp/rand"
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
	randHistory       []liveRadio
}

func newRadio() *radio {
	return &radio{}
}

func (r *radio) OnPreRender(ctx app.Context) {
	r.init(ctx)
	r.load(ctx)
}

func (r *radio) OnMount(ctx app.Context) {
	r.init(ctx)
}

func (r *radio) OnNav(ctx app.Context) {
	r.load(ctx)
}

func (r *radio) OnResize(ctx app.Context) {
	r.ResizeContent()
	r.Update()
}

func (r *radio) init(ctx app.Context) {
	rand.Seed(uint64(time.Now().UnixNano()))
	r.lives = getLiveRadios()
	r.Update()
}

func (r *radio) load(ctx app.Context) {
	slug := strings.TrimPrefix(ctx.Page.URL().Path, "/")
	if slug == "" {
		r.current = r.randomRadio()
		u := *ctx.Page.URL()
		u.Path = "/" + r.current.Slug
		ctx.Page.ReplaceURL(&u)
	} else {
		for _, lr := range r.lives {
			if slug == lr.Slug {
				r.current = lr
				break
			}
		}
	}

	r.isUpdateAvailable = ctx.AppUpdateAvailable
	r.isPlaying = false

	r.Update()

	ctx.Page.SetTitle(fmt.Sprintf("%s Radio", r.current.Name))
	ctx.Page.SetDescription(fmt.Sprintf("Listen to Lo-fi music radio %s on the Lofimusic open-source player: an installable Progressive Web app (PWA) written in Go (Golang).", r.current.Name))

	if app.IsServer {
		ctx.Page.SetImage("https://lofimusic.app/web/covers/" + slug + ".png")
	}
}

func (r *radio) OnAppUpdate(ctx app.Context) {
	r.isUpdateAvailable = ctx.AppUpdateAvailable
	r.Update()
}

func (r *radio) randomRadio() liveRadio {
	if len(r.randHistory) == 0 {
		r.randHistory = make([]liveRadio, len(r.lives))
		copy(r.randHistory, r.lives)
	}

	idx := rand.Intn(len(r.randHistory))
	radio := r.randHistory[idx]

	copy(r.randHistory[idx:], r.randHistory[idx+1:])
	r.randHistory = r.randHistory[:len(r.randHistory)-1]

	return radio
}

func (r *radio) Render() app.UI {
	return app.Main().
		Class("radio").
		Class("fill").
		Body(
			newYouTubePlayer().
				Class("radio-player").
				Class("fill").
				Radio(r.current).
				OnPlaybackChange(r.onPlaybackChange),
			app.Shell().
				Class("radio-shell").
				Class("fill").
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
						Class("radio-update").
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
