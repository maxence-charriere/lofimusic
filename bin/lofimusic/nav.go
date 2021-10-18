package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type nav struct {
	app.Compo

	Iclass        string
	IliveRadios   []liveRadio
	IcurrentRadio liveRadio

	isFirstLoad bool
}

func newNav() *nav {
	return &nav{}
}

func (n *nav) Class(v string) *nav {
	if v == "" {
		return n
	}
	if n.Iclass != "" {
		n.Iclass += " "
	}
	n.Iclass += v
	return n
}

func (n *nav) LiveRadios(v []liveRadio) *nav {
	n.IliveRadios = v
	return n
}

func (n *nav) CurrentRadio(v liveRadio) *nav {
	n.IcurrentRadio = v
	return n
}

func (n *nav) OnMount(ctx app.Context) {
	n.isFirstLoad = true
}

func (n *nav) OnNav(ctx app.Context) {
	if n.isFirstLoad {
		n.isFirstLoad = false
		ctx.ScrollTo(strings.TrimPrefix(ctx.Page().URL().Path, "/"))
	}
}

func (n *nav) Render() app.UI {
	return app.Div().
		Class("nav").
		Class("fill").
		Class("unselectable").
		Class(n.Iclass).
		Body(
			ui.Stack().
				Class("app-title").
				Class("hspace-out").
				Middle().
				Content(
					app.Header().
						Body(
							app.A().
								Class("hApp").
								Class("focus").
								Class("glow").
								Href("/").
								Text("Lofimusic"),
						),
				),
			app.Nav().
				Class("nav-content").
				Body(
					app.Div().
						Class("nav-radios").
						Body(
							ui.Stack().
								Class("nav-radios-stack").
								Middle().
								Content(
									app.Div().
										Class("hspace-out").
										Body(
											app.Range(n.IliveRadios).Slice(func(i int) app.UI {
												lr := n.IliveRadios[i]
												return newLink().
													ID(lr.Slug).
													Class("glow").
													Icon(newSVGIcon().RawSVG(playSVG)).
													Label(lr.Name).
													Href("/" + lr.Slug).
													Focus(lr.Slug == n.IcurrentRadio.Slug)
											}),
										),
								),
						),
					app.Div().
						Class("nav-support").
						Class("hspace-out").
						Body(
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(coffeeSVG)).
								Label("Buy me a coffee").
								Href(buyMeACoffeeURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(cryptoSVG)).
								Label("Donate cryptos").
								Href(coinbaseBusinessURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(githubSVG)).
								Label("GitHub").
								Href(githubURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(twitterSVG)).
								Label("Twitter").
								Href(twitterURL),
						),
				),
		)
}
