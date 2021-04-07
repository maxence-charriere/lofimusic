package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	loaderSize          = 78
	controlIconSize     = 18
	controlMainIconSize = 30
)

type youtubeState int

const (
	unstarted = -1
	ended     = 0
	playing   = 1
	paused    = 2
	buffering = 3
	videoCued = 5
)

type youTubePlayer struct {
	app.Compo

	Iradio liveRadio

	initPlayer           sync.Once
	radio                liveRadio
	player               app.Value
	isPlaying            bool
	isBuffering          bool
	realeaseOnReady      func()
	releaseOnStateChange func()
}

func newYouTubePlayer() *youTubePlayer {
	return &youTubePlayer{}
}

func (p *youTubePlayer) Radio(v liveRadio) *youTubePlayer {
	p.Iradio = v
	return p
}

func (p *youTubePlayer) OnNav(ctx app.Context) {
	p.Update()
}

func (p *youTubePlayer) OnDismount() {
	if p.realeaseOnReady != nil {
		p.realeaseOnReady()
	}
	if p.releaseOnStateChange != nil {
		p.releaseOnStateChange()
	}
	p.player = nil
}

func (p *youTubePlayer) loadVideo(ctx app.Context) {
	if isOnYouTubeIframeAPIReady := app.Window().Get("isOnYouTubeIframeAPIReady").Bool(); !isOnYouTubeIframeAPIReady {
		ctx.Async(func() {
			time.Sleep(time.Millisecond * 1000)
			p.Defer(p.loadVideo)
		})
		return
	}

	if p.Iradio.Slug != p.radio.Slug {
		p.radio = p.Iradio
		if p.player != nil {
			p.loadVideoByID(ctx, p.radio.youtubeID())
			return
		}
	}

	p.initPlayer.Do(func() {
		onReady := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			p.Defer(p.play)
			return nil
		})
		p.realeaseOnReady = onReady.Release

		onStateChange := app.FuncOf(p.onStateChange)
		p.releaseOnStateChange = onStateChange.Release

		p.player = app.Window().
			Get("YT").
			Get("Player").
			New("youtube-player", map[string]interface{}{
				"videoId": p.radio.youtubeID(),
				"playerVars": map[string]interface{}{
					"controls":       0,
					"modestbranding": 1,
				},
				"events": map[string]interface{}{
					"onReady":       onReady,
					"onStateChange": onStateChange,
				},
			})
	})
}

func (p *youTubePlayer) onStateChange(this app.Value, args []app.Value) interface{} {
	p.Defer(func(ctx app.Context) {
		switch args[0].Get("data").Int() {
		case unstarted:
			p.isPlaying = false

		case ended:
			p.isPlaying = false

		case playing:
			p.isPlaying = true
			p.isBuffering = false

		case paused:
			p.isPlaying = false

		case buffering:
			p.isBuffering = true
		}
		p.Update()
	})
	return nil
}

func (p *youTubePlayer) Render() app.UI {
	if p.Iradio.Slug != "" && p.radio.Slug != p.Iradio.Slug {
		p.Defer(p.loadVideo)
	}

	return app.Div().
		Class("youtube").
		Class("fill").
		Class("unselectable").
		Body(
			app.Div().
				Class("youtube-video").
				Body(
					app.Div().
						ID("youtube-player").
						Body(
							app.Script().Src("https://www.youtube.com/iframe_api"),
						),
				),
			app.If(!p.isPlaying || p.isBuffering,
				app.Div().
					Class("youtube-noplay").
					Class("fill").
					Body(
						newLoader().
							Class("hspace-out").
							Class("vspace-stretch").
							Size(loaderSize).
							Title("Buffering").
							Description(p.radio.Name).
							Loading(p.isBuffering),
					),
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
							app.If(p.isPlaying || p.isBuffering,
								newControl().Icon(newSVGIcon().
									Size(controlIconSize).
									RawSVG(pauseSVG)).
									Disabled(p.player == nil).
									OnClick(p.onPauseClicked),
							).Else(
								newControl().Icon(newSVGIcon().
									Size(controlIconSize).
									RawSVG(playSVG)).
									Disabled(p.player == nil).
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
								Disabled(p.player == nil).
								OnClick(p.onMuteClicked),
						),
				),
		)
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

func (p *youTubePlayer) loadVideoByID(ctx app.Context, id string) {
	p.player.Call("loadVideoById", id, 0)
}

func (p *youTubePlayer) play(ctx app.Context) {
	p.player.Call("playVideo")
}

func (p *youTubePlayer) pause(ctx app.Context) {
	p.player.Call("pauseVideo")
}
