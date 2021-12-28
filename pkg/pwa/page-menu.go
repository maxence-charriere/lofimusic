package pwa

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"github.com/maxence-charriere/lofimusic/pkg/radio"
)

type pageMenu struct {
	app.Compo

	Iclass string

	videos           []radio.Video
	isAppInstallable bool
}

func newPageMenu() *pageMenu {
	return &pageMenu{
		videos: radio.List(),
	}
}

func (m *pageMenu) Class(v string) *pageMenu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *pageMenu) OnNav(ctx app.Context) {
	m.handleAppInstallable(ctx)
}

func (m *pageMenu) OnAppInstallChange(ctx app.Context) {
	m.handleAppInstallable(ctx)
}

func (m *pageMenu) handleAppInstallable(ctx app.Context) {
	m.isAppInstallable = ctx.IsAppInstallable()
}

func (m *pageMenu) Render() app.UI {
	return ui.Scroll().
		Class("page-menu").
		Class("fill").
		Class("unselectable").
		Class(m.Iclass).
		HeaderHeight(pageHeaderHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					app.Header().Body(
						app.A().
							Class("app-title").
							Class("accent-hover").
							Class("block").
							Href("/").
							Text(Info.Name),
					),
				),
		).
		Content(
			ui.Stack().
				Class("fill-v-min").
				Middle().
				Content(
					app.Nav().Body(
						app.Range(m.videos).Slice(func(i int) app.UI {
							v := m.videos[i]
							return ui.Link().
								Class(linkClass).
								Icon(playSVG).
								Label(v.Name).
								Class(pathFocus(v.Slug)).
								Href("/" + v.Slug)
						}),

						newSectionBreak(),

						ui.Link().
							Class(linkClass).
							Icon(twitterSVG).
							Label("Twitter").
							Help(fmt.Sprintf("Follow %s's creator on Twitter.", Info.Name)).
							Href(twitterURL),
						ui.Link().
							Class(linkClass).
							Icon(coffeeSVG).
							Label("Buy Me A Coffee").
							Help(fmt.Sprintf("Support %s's creator on Buy Me A Coffee.", Info.Name)).
							Href(buyMeACoffeeURL),
						ui.Link().
							Class(linkClass).
							Icon(githubSVG).
							Label("GitHub").
							Help(fmt.Sprintf("See %s code on GitHub.", Info.Name)).
							Href(githubURL),
						ui.Link().
							Class(linkClass).
							Icon(cryptoSVG).
							Label("Donate Crypto").
							Help(fmt.Sprintf("Support %s's creator by donating crypto.", Info.Name)).
							Href(coinbaseBusinessURL),
					),
				),
		)
}
