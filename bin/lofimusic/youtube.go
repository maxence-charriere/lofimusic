package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	controlIconSize     = 18
	controlMainIconSize = 30
)

var (
	isOnYouTubeIframeAPIReady bool
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

	once                 sync.Once
	radio                liveRadio
	player               app.Value
	isPlaying            bool
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

func (p *youTubePlayer) OnMount(ctx app.Context) {
	if !isOnYouTubeIframeAPIReady {
		releaseOnYouTubeIframeAPIReady := func() {}
		onYouTubeIframeAPIReady := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			isOnYouTubeIframeAPIReady = true
			releaseOnYouTubeIframeAPIReady()
			return nil
		})
		releaseOnYouTubeIframeAPIReady = onYouTubeIframeAPIReady.Release
		app.Window().Set("onYouTubeIframeAPIReady", onYouTubeIframeAPIReady)
	}
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
	if !isOnYouTubeIframeAPIReady {
		ctx.Async(func() {
			time.Sleep(time.Millisecond * 100)
			p.Defer(p.loadVideo)
		})
		return
	}

	if p.Iradio.Slug != p.radio.Slug {
		p.radio = p.Iradio
		if p.player != nil {
			fmt.Println("loading video:", p.radio.Slug, p.radio.youtubeID())
			p.loadVideoByID(ctx, p.radio.youtubeID())
			return
		}
	}

	p.once.Do(func() {
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

		case paused:
			p.isPlaying = false

		case buffering:
		}
		p.Update()
	})
	return nil
}

func (p *youTubePlayer) Render() app.UI {
	if p.radio.Slug != p.Iradio.Slug {
		p.Defer(p.loadVideo)
	}

	return app.Div().
		Class("youtube").
		Class("fill").
		Body(
			app.Div().
				Class("youtube-video").
				Body(
					app.Div().
						ID("youtube-player").
						Body(
							app.Script().
								Src("https://www.youtube.com/iframe_api").
								Async(true),
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

func (p *youTubePlayer) loadVideoByID(ctx app.Context, id string) {
	p.player.Call("loadVideoById", id, 0)
}

func (p youTubePlayer) play(ctx app.Context) {
	p.player.Call("playVideo")
}

func (p youTubePlayer) pause(ctx app.Context) {
	p.player.Call("pauseVideo")
}
