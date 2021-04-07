package main

import (
	"fmt"
	"path"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	controlIconSize     = 18
	controlMainIconSize = 30
)

type youTubePlayer struct {
	app.Compo

	Iradio liveRadio

	player    app.Value
	isPlaying bool
}

func newYouTubePlayer() *youTubePlayer {
	return &youTubePlayer{}
}

func (p *youTubePlayer) Radio(v liveRadio) *youTubePlayer {
	p.Iradio = v
	return p
}

func (p *youTubePlayer) OnMount(ctx app.Context) {
}

func (p *youTubePlayer) OnDismount() {
}

func (p *youTubePlayer) Render() app.UI {
	return app.Div().
		Class("youtube").
		Class("fill").
		Body(
			app.Div().
				Class("youtube-video").
				Body(
					app.Script().
						Src("//www.youtube.com/iframe_api").
						Async(true).
						OnLoad(p.onScriptLoaded),
					app.IFrame().
						ID("youtube-"+p.Iradio.Slug).
						Allow("autoplay").
						Allow("accelerometer").
						Allow("encrypted-media").
						Allow("picture-in-picture").
						Sandbox("allow-presentation allow-same-origin allow-scripts allow-popups").
						Src(fmt.Sprintf(
							"https://www.youtube.com/embed/%s?controls=0&showinfo=0&autoplay=1&loop=1&enablejsapi=1&playsinline=1",
							path.Base(p.Iradio.URL),
						)),
				),
			app.Div().
				Class("youtube-controls").
				Class("hspace-out").
				Class("vspace-top").
				Class("vspace-bottom").
				Body(
					app.Stack().
						Class("fit").
						Class("center").
						Center().
						Content(
							app.If(p.isPlaying,
								newControl().Icon(newSVGIcon().
									Size(controlIconSize).
									RawSVG(pauseSVG)).
									OnClick(p.onPauseClicked),
							).Else(
								newControl().Icon(newSVGIcon().
									Size(controlIconSize).
									RawSVG(playSVG)).
									OnClick(p.onPlayClicked),
							),
							newControl().
								Class("control-main").
								Icon(newSVGIcon().
									Size(controlMainIconSize).
									RawSVG(shuffleSVG)).
								OnClick(p.onShuffleClicked),
							newControl().Icon(newSVGIcon().
								Size(controlIconSize).
								RawSVG(soundHighSVG)).
								OnClick(p.onMuteClicked),
						),
				),
		)
}

func (p *youTubePlayer) onScriptLoaded(ctx app.Context, e app.Event) {
	fmt.Println("youtube scrip loaded")
}

func (p *youTubePlayer) onPlayClicked(ctx app.Context, e app.Event) {
	p.play(ctx)
}

func (p *youTubePlayer) onPauseClicked(ctx app.Context, e app.Event) {
	p.pause(ctx)
}

func (p *youTubePlayer) onShuffleClicked(ctx app.Context, e app.Event) {
	fmt.Println("shuffle clicked")
}

func (p *youTubePlayer) onMuteClicked(ctx app.Context, e app.Event) {
	fmt.Println("mute clicked")
}

func (p youTubePlayer) play(ctx app.Context) {
	p.player.Call("playVideo")
}

func (p youTubePlayer) pause(ctx app.Context) {
	p.player.Call("pauseVideo")
}
