package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type detail struct {
	app.Compo

	Channel channel
}

func (d *detail) Render() app.UI {
	return app.Main().
		Class("detail").
		Body(
			app.Section().
				Class("channel").
				Body(
					app.H1().Text(d.Channel.Name),
				),
		)
}
