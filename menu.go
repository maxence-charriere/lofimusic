package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type menu struct {
	app.Compo

	CurrentChannel channel
}

func (m *menu) Render() app.UI {
	chans := channels.Channels()

	return app.Stack().
		Class("menu").
		Center().
		Content(
			app.Nav().
				Body(
					app.H1().Text("Lofimusic"),
					app.Section().Body(
						app.Range(chans).Slice(func(i int) app.UI {
							c := chans[i]
							return newMenuItem().
								Text(c.Name).
								Selected(c.Slug == m.CurrentChannel.Slug).
								Href("/" + c.Slug).
								Help(fmt.Sprintf("Play %s.", c.Name)).
								Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
								<path fill="currentColor" d="M8,5.14V19.14L19,12.14L8,5.14Z" />
							</svg>
							`)
						}),
					),
					app.Section().Body(
						newMenuItem().
							Text("Buy me a coffee").
							Href("https://www.buymeacoffee.com/Lofimusicapp").
							Help("Buy a coffee to support Lofimusic.app.").
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
								<path fill="currentColor" d="M2,21H20V19H2M20,8H18V5H20M20,3H4V13A4,4 0 0,0 8,17H14A4,4 0 0,0 18,13V10H20A2,2 0 0,0 22,8V5C22,3.89 21.1,3 20,3Z" />
							</svg>
							`),
						newMenuItem().
							Text("Twitter").
							Href("https://twitter.com/jonhymaxoo").
							Help("Follow on Twitter").
							Icon(`
							<svg style="width:24px;height:24px" viewBox="0 0 24 24">
								<path fill="currentColor" d="M22.46,6C21.69,6.35 20.86,6.58 20,6.69C20.88,6.16 21.56,5.32 21.88,4.31C21.05,4.81 20.13,5.16 19.16,5.36C18.37,4.5 17.26,4 16,4C13.65,4 11.73,5.92 11.73,8.29C11.73,8.63 11.77,8.96 11.84,9.27C8.28,9.09 5.11,7.38 3,4.79C2.63,5.42 2.42,6.16 2.42,6.94C2.42,8.43 3.17,9.75 4.33,10.5C3.62,10.5 2.96,10.3 2.38,10C2.38,10 2.38,10 2.38,10.03C2.38,12.11 3.86,13.85 5.82,14.24C5.46,14.34 5.08,14.39 4.69,14.39C4.42,14.39 4.15,14.36 3.89,14.31C4.43,16 6,17.26 7.89,17.29C6.43,18.45 4.58,19.13 2.56,19.13C2.22,19.13 1.88,19.11 1.54,19.07C3.44,20.29 5.7,21 8.12,21C16,21 20.33,14.46 20.33,8.79C20.33,8.6 20.33,8.42 20.32,8.23C21.16,7.63 21.88,6.87 22.46,6Z" />
							</svg>
							`),
					),
				),
		)
}

type menuItem struct {
	app.Compo

	Iicon     string
	Ihref     string
	Iselected string
	Itext     string
	Ihelp     string
}

func newMenuItem() *menuItem {
	return &menuItem{}
}

func (i *menuItem) Icon(svg string) *menuItem {
	i.Iicon = svg
	return i
}

func (i *menuItem) Href(v string) *menuItem {
	i.Ihref = v
	return i
}

func (i *menuItem) Selected(v bool) *menuItem {
	if v {
		i.Iselected = "focus"
	}
	return i
}

func (i *menuItem) Text(v string) *menuItem {
	i.Itext = v
	return i
}

func (i *menuItem) Help(v string) *menuItem {
	i.Ihelp = v
	return i
}

func (i *menuItem) Render() app.UI {
	return app.A().
		Class("item").
		Class(i.Iselected).
		Href(i.Ihref).
		Title(i.Ihelp).
		Body(
			app.Stack().
				Center().
				Content(
					app.Div().
						Class("icon").
						Body(
							app.Raw(i.Iicon),
						),
					app.Div().
						Class("label").
						Text(i.Itext),
				),
		)
}
