package main

import (
	"fmt"

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
					app.Stack().
						Class("social").
						Content(
							app.Range(d.Channel.SocialMedia).Slice(func(i int) app.UI {
								link := d.Channel.SocialMedia[i]
								media := medias[link.MediaSlug]

								return newButton().
									Help(fmt.Sprintf("See %s on %s.", d.Channel.Name, media.Name)).
									Icon(media.Icon).
									URL(link.URL)
							}),
						),
				),
		)
}

type button struct {
	app.Compo

	Iurl  string
	Iicon string
	Ihelp string
}

func newButton() *button {
	return &button{}
}

func (b *button) URL(v string) *button {
	b.Iurl = v
	return b
}

func (b *button) Icon(svg string) *button {
	b.Iicon = svg
	return b
}

func (b *button) Help(v string) *button {
	b.Ihelp = v
	return b
}

func (b *button) Render() app.UI {
	return app.A().
		Class("button").
		Title(b.Ihelp).
		Href(b.Iurl).
		Target("_blank").
		Body(
			app.Raw(b.Iicon),
		)
}
