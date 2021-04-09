package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
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

	Iradio            liveRadio
	IonPlaybackChange func(app.Context, bool)

	initPlayer  sync.Once
	radio       liveRadio
	player      app.Value
	isPlaying   bool
	isBuffering bool
	canBack     bool
	volume      volume
	err         error

	realeaseOnReady      func()
	releaseOnStateChange func()
	releaseOnError       func()
}

func newYouTubePlayer() *youTubePlayer {
	return &youTubePlayer{}
}

func (p *youTubePlayer) Radio(v liveRadio) *youTubePlayer {
	p.Iradio = v
	return p
}

func (p *youTubePlayer) OnPlaybackChange(v func(app.Context, bool)) *youTubePlayer {
	p.IonPlaybackChange = v
	return p
}

func (p *youTubePlayer) OnMount(ctx app.Context) {
	p.isBuffering = true
	p.volume = loadVolume(ctx)
	p.Update()
}

func (p *youTubePlayer) OnNav(ctx app.Context) {
	p.canBack = app.Window().Get("history").Get("length").Int() > 1
	p.Update()
}

func (p *youTubePlayer) OnDismount() {
	if p.realeaseOnReady != nil {
		p.realeaseOnReady()
	}
	if p.releaseOnStateChange != nil {
		p.releaseOnStateChange()
	}
	if p.releaseOnError != nil {
		p.releaseOnError()
	}
	p.player = nil
}

func (p *youTubePlayer) OnResize(ctx app.Context) {
	p.Update()
}

func (p *youTubePlayer) loadVideo(ctx app.Context) {
	if isOnYouTubeIframeAPIReady := app.Window().Get("isOnYouTubeIframeAPIReady").Bool(); !isOnYouTubeIframeAPIReady && app.IsClient {
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
			p.Defer(func(ctx app.Context) {
				p.setVolume(ctx, p.volume.Value)
				p.play(ctx)
			})
			return nil
		})
		p.realeaseOnReady = onReady.Release

		onStateChange := app.FuncOf(p.onStateChange)
		p.releaseOnStateChange = onStateChange.Release

		onError := app.FuncOf(p.onError)
		p.releaseOnError = onError.Release

		p.player = app.Window().
			Get("YT").
			Get("Player").
			New("youtube-player", map[string]interface{}{
				"videoId": p.radio.youtubeID(),
				"playerVars": map[string]interface{}{
					"autoplay":       1,
					"controls":       0,
					"modestbranding": 1,
					"disablekb":      1,
					"iv_load_policy": 3,
					"playsinline":    1,
					"origin":         "lofimusic.app",
				},
				"events": map[string]interface{}{
					"onReady":       onReady,
					"onStateChange": onStateChange,
					"onError":       onError,
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
			if p.err == nil {
				p.play(ctx)
			}

		case playing:
			p.isPlaying = true
			p.isBuffering = false
			p.err = nil

		case paused:
			p.isPlaying = false

		case buffering:
			p.isBuffering = true
			p.err = nil
		}
		p.Update()

		if p.IonPlaybackChange != nil {
			p.Defer(func(ctx app.Context) {
				p.IonPlaybackChange(ctx, p.isPlaying)
			})
		}
	})
	return nil
}

func (p *youTubePlayer) onError(this app.Value, args []app.Value) interface{} {
	p.Defer(func(ctx app.Context) {
		code := args[0].Get("data").Int()
		msg := ""

		switch code {
		case 2:
			msg = "invalid video parameter values"

		case 5:
			msg = "loading video failed"

		case 100:
			msg = "video not found"

		case 101, 150:
			msg = "video cannot be played in embedded players"

		default:
			msg = "unkown error"

		}

		p.err = errors.New("youtube player error").
			Tag("code", code).
			Tag("description", msg)

		fmt.Println("error:", p.err)
		p.Update()
	})
	return nil
}

func (p *youTubePlayer) Render() app.UI {
	if p.Iradio.Slug != "" && p.radio.Slug != p.Iradio.Slug {
		p.Defer(p.loadVideo)
	}

	volumeDisplay := ""
	if p.player == nil {
		volumeDisplay = "disabled"
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
						Class("unselectable").
						Body(
							app.Script().Src("https://www.youtube.com/iframe_api"),
						),
				),
			app.If(!p.isPlaying || p.isBuffering || p.err != nil,
				app.Div().
					Class("youtube-noplay").
					Class("background-overlay").
					Class("fill").
					Body(
						newLoader().
							Class("hspace-out").
							Class("vspace-stretch").
							Size(loaderSize).
							Title("Buffering").
							Description(p.radio.Name).
							Loading(p.isBuffering).
							Err(p.err),
					),
			),
			app.Stack().
				Class("youtube-controls").
				Class("hspace-out").
				Class("vspace-top").
				Class("vspace-bottom").
				Center().
				Content(
					app.Div().Class("youtube-left-space"),
					newControl().
						Class("youtube-back").
						Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(backwardSVG)).
						Disabled(!p.canBack).
						OnClick(p.onBackClicked),
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
					app.If(p.volume.Value > 60,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundHighSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).ElseIf(p.volume.Value > 20,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundMediumSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).ElseIf(p.volume.Value > 0,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundLowSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).Else(
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundMutedSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					),
					app.Div().
						Class("youtube-volume").
						Class(volumeDisplay).
						Body(
							app.Input().
								ID("youtube-volume").
								Type("range").
								Min("0").
								Max("100").
								Value(p.volume.Value).
								OnChange(p.onVolumeChanged).
								OnInput(p.onVolumeChanged),
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

func (p *youTubePlayer) onBackClicked(ctx app.Context, e app.Event) {
	app.Window().Get("history").Call("back")
}

func (p *youTubePlayer) onShuffleClicked(ctx app.Context, e app.Event) {
	ctx.Navigate("/")
}

func (p *youTubePlayer) onSoundClicked(ctx app.Context, e app.Event) {
	if p.volume.Value == 0 {
		p.setVolume(ctx, p.volume.LastHearableValue)
		return
	}
	p.setVolume(ctx, 0)
}

func (p *youTubePlayer) onVolumeChanged(ctx app.Context, e app.Event) {
	volume, _ := strconv.Atoi(ctx.JSSrc.Get("value").String())
	p.setVolume(ctx, volume)
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

func (p *youTubePlayer) setVolume(ctx app.Context, v int) {
	defer p.Update()

	if v == 0 {
		p.player.Call("mute")
	} else {
		p.volume.LastHearableValue = v
		p.player.Call("unMute")
		p.player.Call("setVolume", v)
	}

	p.volume.Value = v
	saveVolume(ctx, p.volume)
}

type volume struct {
	Value             int
	LastHearableValue int
}

func saveVolume(ctx app.Context, v volume) {
	if err := ctx.LocalStorage().Set("player-volume", v); err != nil {
		app.Log(errors.New("saving volume values in local storage failed").
			Wrap(err))
	}
}

func loadVolume(ctx app.Context) volume {
	var v volume
	if err := ctx.LocalStorage().Get("player-volume", &v); err != nil {
		app.Log(errors.New("saving player status in local storage failed").
			Wrap(err))
	}

	if v == (volume{}) {
		v.Value = 100
		v.LastHearableValue = 100
	}
	return v
}
