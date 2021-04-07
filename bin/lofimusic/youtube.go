package main

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type youTubePlayer struct {
	app.Compo
}

func newYouTubePlayer() *youTubePlayer {
	return &youTubePlayer{}
}

func (p *youTubePlayer) Render() app.UI {
	return app.Div().
		Class("youtube").
		Class("fill").
		Body(
			app.Div().Class("youtube-video"),
			app.Div().Class("youtube-controls"),
		)
}
