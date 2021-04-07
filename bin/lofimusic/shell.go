package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type shell struct {
	app.Compo

	ChannelSlug string
}

func (s *shell) Render() app.UI {
	c := channels.Get(s.ChannelSlug)

	return app.Div().Body(
		&player{Channel: c},
		app.Shell().
			Menu(&menu{CurrentChannel: c}).
			OverlayMenu(
				app.Div().
					Class("overlay-menu").
					Body(&menu{CurrentChannel: c}),
			).
			Content(&detail{Channel: c}),
	)
}
