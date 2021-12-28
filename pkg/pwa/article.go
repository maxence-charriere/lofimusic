package pwa

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"github.com/maxence-charriere/lofimusic/pkg/radio"
)

const (
	headerLinkSize      = 22
	cardVisibleDuration = time.Second * 10
	cardHiddenDuration  = time.Second * 5
)

type article struct {
	app.Compo

	video         radio.Video
	headerVisible bool
	initCycle     sync.Once
	cardVisible   bool
	currentCard   int
}

func newArticle() *article {
	return &article{}
}

func (a *article) OnMount(ctx app.Context) {
	ctx.ObserveState(currentVideoState).
		OnChange(func() {
			a.load(ctx)
			a.cycleCards(ctx)
		}).
		Value(&a.video)
}

func (a *article) Render() app.UI {
	links := a.video.Links()

	headerVisibility := "article-header-hide"
	if a.headerVisible {
		headerVisibility = "article-header-show"
	}

	return app.Article().
		Class("relative").
		Class("fill").
		Body(
			app.Header().
				Class("article-header").
				Class("fit").
				Class(headerVisibility).
				Body(
					app.H1().
						Class("article-header-padding-h").
						Class("text-center").
						Text(a.video.Name),
					app.Div().
						Class("fill-h").
						Class("below-xs").
						Class("underline"),
					ui.Stack().
						Class("below-xs").
						Center().
						Middle().
						Content(
							app.Range(links).Slice(func(i int) app.UI {
								l := links[i]
								return ui.Link().
									Class(linkClass).
									Class("article-link").
									Icon(socialIcon(l.Slug)).
									IconSize(headerLinkSize).
									Href(l.URL).
									Help(fmt.Sprintf("Go to %s's %s",
										strings.Title(a.video.Owner),
										strings.Title(l.Slug),
									))
							}),
						),
				),
			app.Range(a.video.Cards).Slice(func(i int) app.UI {
				cardVisibility := "article-card-hide"
				if i == a.currentCard && a.headerVisible && a.cardVisible {
					cardVisibility = "info-card-show"
				}

				return app.P().
					Class("article-card").
					Class(cardVisibility).
					Text(a.video.Cards[i])
			}),
		)
}

func (a *article) load(ctx app.Context) {
	a.headerVisible = false

	ctx.Defer(func(ctx app.Context) {
		ctx.Dispatch(func(ctx app.Context) {
			a.headerVisible = true
		})
	})
}

func (a *article) cycleCards(ctx app.Context) {
	a.currentCard = -1
	a.cardVisible = false

	a.initCycle.Do(func() {
		ctx.Async(func() {
			ticker := time.NewTicker(cardHiddenDuration)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return

				case <-ticker.C:
					ctx.Dispatch(func(ctx app.Context) {
						if a.cardVisible {
							ticker.Reset(cardHiddenDuration)
							a.cardVisible = false
							return
						}

						if len(a.video.Cards) == 0 {
							return
						}

						a.currentCard = (a.currentCard + 1) % len(a.video.Cards)
						ticker.Reset(cardVisibleDuration)
						a.cardVisible = true
					})
				}
			}
		})
	})
}
