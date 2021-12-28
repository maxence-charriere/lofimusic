package pwa

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type pageTop struct {
	app.Compo

	Iclass   string
	Icontent []app.UI

	isAppUpdatable bool
}

func newPageTop() *pageTop {
	return &pageTop{}
}

func (t *pageTop) Class(v string) *pageTop {
	t.Iclass = app.AppendClass(t.Iclass, v)
	return t
}

func (t *pageTop) Content(v ...app.UI) *pageTop {
	t.Icontent = app.FilterUIElems(v...)
	return t
}

func (t *pageTop) OnNav(ctx app.Context) {
	t.handleAppUpdatable(ctx)
}

func (t *pageTop) OnAppUpdate(ctx app.Context) {
	t.handleAppUpdatable(ctx)
}

func (t *pageTop) handleAppUpdatable(ctx app.Context) {
	t.isAppUpdatable = ctx.AppUpdateAvailable()
}

func (t *pageTop) Render() app.UI {
	return ui.Stack().
		Class("fill").
		Class("unselectable").
		Right().
		Middle().
		Content(
			app.Nav().
				Class("page-top").
				Class("text-right").
				Class(t.Iclass).
				Body(
					app.If(t.isAppUpdatable,
						app.Div().
							Class("animate").
							Class("bounce").
							Body(
								ui.Link().
									Class(linkClass).
									Icon(downloadSVG).
									Label("Update").
									Help("Click to apply the latest updates.").
									OnClick(t.onUpdateClick),
							),
					),
					app.Range(t.Icontent).Slice(func(i int) app.UI {
						return t.Icontent[i]
					}),
				),
		)
}

func (t *pageTop) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.NewAction(updateApp)
}
