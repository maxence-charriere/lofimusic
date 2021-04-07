package main

import (
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type nav struct {
	app.Compo

	IliveRadios   []liveRadio
	IcurrentRadio string

	isFirstLoad bool
}

func newNav() *nav {
	return &nav{}
}

func (n *nav) LiveRadios(v []liveRadio) *nav {
	n.IliveRadios = v
	return n
}

func (n *nav) CurrentRadio(v string) *nav {
	n.IcurrentRadio = v
	return n
}

func (n *nav) OnMount(ctx app.Context) {
	n.isFirstLoad = true
}

func (n *nav) OnNav(ctx app.Context) {
	if n.isFirstLoad {
		n.isFirstLoad = false
		ctx.ScrollTo(strings.TrimPrefix(ctx.Page.URL().Path, "/"))
	}
}

func (n *nav) Render() app.UI {
	return app.Div().
		Class("nav").
		Class("fill").
		Class("unselectable").
		Body(
			app.Stack().
				Class("app-title").
				Class("hspace-out").
				Center().
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
							app.Stack().
								Class("nav-radios-stack").
								Center().
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
													Focus(lr.Slug == n.IcurrentRadio)
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
