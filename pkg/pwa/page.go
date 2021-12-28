package pwa

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	pagePaneSize     = 282
	pageHeaderHeight = 90

	linkClass       = "page-menu-link fit heading accent-hover hover icon-circle"
	footerLinkClass = "page-menu-link fit accent-hover hover"
)

type page struct {
	app.Compo
}

func newPage() *page {
	return &page{}
}

func (p *page) OnPreRender(ctx app.Context) {
	p.init(ctx)
}

func (p *page) OnNav(ctx app.Context) {
	p.init(ctx)
}

func (p *page) init(ctx app.Context) {
	ctx.NewAction(loadVideo, app.T("path", ctx.Page().URL().Path))
}

func (p *page) Render() app.UI {
	return ui.Shell().
		Class("fill").
		Class("page-background").
		PaneWidth(pagePaneSize).
		Menu(newPageMenu()).
		HamburgerMenu(
			newPageMenu().
				Class("background-overlay").
				Class("page-menu-overlay"),
		).
		Content(
			ui.Scroll().
				Class("fill").
				HeaderHeight(pageHeaderHeight).
				Header(
				// newPageTop().
				// 	Content(p.ItopMenu...).
				// 	Subcontent(p.ItopIndex...),
				).
				Content(
					app.Main().
						Class("fill").
						Body(
							newArticle(),
						),
				),
		)
}
