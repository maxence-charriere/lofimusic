package main

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	infoLinkIconSize = 18
)

type info struct {
	app.Compo

	Iclass   string
	Iradio   liveRadio
	Iplaying bool
}

func newInfo() *info {
	return &info{}
}

func (i *info) Class(v string) *info {
	if v == "" {
		return i
	}
	if i.Iclass != "" {
		i.Iclass += " "
	}
	i.Iclass += v
	return i
}

func (i *info) Radio(v liveRadio) *info {
	i.Iradio = v
	return i
}

func (i *info) Playing(v bool) *info {
	i.Iplaying = v
	return i
}

func (i *info) Render() app.UI {
	titleVisibility := ""
	if i.Iplaying {
		titleVisibility = "info-title-show"
	}

	return app.Article().
		Class("info").
		Class("fill").
		Class("no-overflow").
		Body(
			app.Header().
				Class("info-title").
				Class(titleVisibility).
				Class("center").
				Class("fit").
				Body(
					app.H1().
						Class("h1").
						Class("glow").
						Text(i.Iradio.Name),
					app.Div().Class("info-title-separator"),
					app.Stack().
						Class("info-links").
						Class("fit").
						Class("center").
						Center().
						Content(
							app.Range(i.Iradio.Links).Slice(func(j int) app.UI {
								l := i.Iradio.Links[j]
								return newInfoLink().
									Help(fmt.Sprintf("Visit %s's %s.",
										strings.Title(i.Iradio.Owner),
										strings.Title(l.Slug),
									)).
									Href(l.URL).
									Icon(newSVGIcon().
										Size(infoLinkIconSize).
										RawSVG(socialIcon(l.Slug)))
							}),
						),
				),
		)
}
