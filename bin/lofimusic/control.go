package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type control struct {
	app.Compo

	Iclass    string
	Iicon     app.UI
	Idisabled bool
	IonClick  app.EventHandler
}

func newControl() *control {
	return &control{}
}

func (c *control) Class(v string) *control {
	if v == "" {
		return c
	}
	if c.Iclass != "" {
		c.Iclass += " "
	}
	c.Iclass += v
	return c
}

func (c *control) Icon(v app.UI) *control {
	c.Iicon = v
	return c
}

func (c *control) Disabled(v bool) *control {
	c.Idisabled = v
	return c
}

func (c *control) OnClick(v app.EventHandler) *control {
	c.IonClick = v
	return c
}

func (c *control) Render() app.UI {
	disabled := ""
	if c.Idisabled {
		disabled = "disabled"
	}

	return app.Button().
		Class("control").
		Class(disabled).
		Class(c.Iclass).
		OnClick(c.onClick).
		Body(c.Iicon)
}

func (c *control) onClick(ctx app.Context, e app.Event) {
	if c.IonClick != nil {
		c.IonClick(ctx, e)
	}
}
