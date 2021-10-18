package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type link struct {
	app.Compo

	Iid      string
	Iclass   string
	Ilabel   string
	Ihelp    string
	Ihref    string
	Ifocus   bool
	IonClick func(app.Context)
	Iicon    app.UI
}

func newLink() *link {
	return &link{}
}

func (l *link) ID(v string) *link {
	l.Iid = v
	return l
}

func (l *link) Class(v string) *link {
	if v == "" {
		return l
	}
	if l.Iclass != "" {
		l.Iclass += " "
	}
	l.Iclass += v
	return l
}

func (l *link) Label(v string) *link {
	l.Ilabel = v
	return l
}

func (l *link) Help(v string) *link {
	l.Ihelp = v
	return l
}

func (l *link) Href(v string) *link {
	l.Ihref = v
	return l
}

func (l *link) Focus(v bool) *link {
	l.Ifocus = v
	return l
}

func (l *link) OnClick(v func(app.Context)) *link {
	l.IonClick = v
	return l
}

func (l *link) Icon(v app.UI) *link {
	l.Iicon = v
	return l
}

func (l *link) Render() app.UI {
	iconVisibility := ""
	if l.Iicon == nil {
		iconVisibility = "hide"
	}

	focus := ""
	if l.Ifocus {
		focus = "focus"
	}

	return app.A().
		ID(l.Iid).
		Class("link").
		Class("heading").
		Class("fit").
		Class(l.Iclass).
		Class(focus).
		Title(l.Ihelp).
		Href(l.Ihref).
		OnClick(l.onClick).
		Body(
			ui.Stack().
				Middle().
				Content(
					app.Div().
						Class(iconVisibility).
						Class("link-icon").
						Body(l.Iicon),
					app.Div().Text(l.Ilabel),
				),
		)
}

func (l *link) onClick(ctx app.Context, e app.Event) {
	if l.IonClick == nil {
		return
	}

	e.PreventDefault()
	l.IonClick(ctx)
}

type infoLink struct {
	app.Compo

	Iclass string
	Ihref  string
	Ihelp  string
	Iicon  app.UI
}

func newInfoLink() *infoLink {
	return &infoLink{}
}

func (l *infoLink) Class(v string) *infoLink {
	if v == "" {
		return l
	}
	if l.Iclass != "" {
		l.Iclass += " "
	}
	l.Iclass += v
	return l
}

func (l *infoLink) Href(v string) *infoLink {
	l.Ihref = v
	return l
}

func (l *infoLink) Help(v string) *infoLink {
	l.Ihelp = v
	return l
}

func (l *infoLink) Icon(v app.UI) *infoLink {
	l.Iicon = v
	return l
}

func (l *infoLink) Render() app.UI {
	return app.A().
		Class("infolink").
		Class(l.Iclass).
		Title(l.Ihelp).
		Href(l.Ihref).
		Body(
			l.Iicon,
		)
}
