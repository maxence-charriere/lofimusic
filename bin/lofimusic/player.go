package main

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
)

var (
	youtubeAPIReady bool
)

type player struct {
	app.Compo

	Channel channel
	State   playerState

	youtube                  app.Value
	releaseIframe            func()
	releaseOnPlayerReady     func()
	releasePlayerStateChange func()
	ready                    bool
	playing                  bool
}

func (p *player) OnMount(ctx app.Context) {
	p.State.load(ctx)
	p.Update()

	if !youtubeAPIReady {
		onYouTubeIframeAPIReady := app.FuncOf(p.onYoutubeIframeAPIReady)
		p.releaseIframe = onYouTubeIframeAPIReady.Release
		app.Window().Set("onYouTubeIframeAPIReady", onYouTubeIframeAPIReady)
		return
	}

	p.Defer(p.setupYoutubePlayer)
}

func (p *player) onYoutubeIframeAPIReady(this app.Value, args []app.Value) interface{} {
	p.Defer(func(ctx app.Context) {
		youtubeAPIReady = true
		p.setupYoutubePlayer(ctx)
	})

	return nil
}

func (p *player) setupYoutubePlayer(ctx app.Context) {
	onPlayerReady := app.FuncOf(p.onPlayerReady)
	p.releaseOnPlayerReady = onPlayerReady.Release

	onPlayerStateChange := app.FuncOf(p.onPlayerStateChange)
	p.releasePlayerStateChange = onPlayerStateChange.Release

	p.youtube = app.Window().
		Get("YT").
		Get("Player").
		New("youtube-"+p.Channel.Slug, map[string]interface{}{
			"events": map[string]interface{}{
				"onReady":       onPlayerReady,
				"onStateChange": onPlayerStateChange,
			},
		})
}

func (p *player) OnDismount() {
	if p.releaseIframe != nil {
		p.releaseIframe()
	}

	if p.releaseOnPlayerReady != nil {
		p.releaseOnPlayerReady()
	}

	if p.releasePlayerStateChange != nil {
		p.releasePlayerStateChange()
	}
}

func (p *player) onPlayerReady(this app.Value, args []app.Value) interface{} {
	p.Defer(func(ctx app.Context) {
		p.setVolume(ctx, p.State.Volume)
		p.play()
	})

	return nil
}

func (p *player) onPlayerStateChange(this app.Value, args []app.Value) interface{} {
	ready := false
	playing := false

	switch args[0].Get("data").Int() {
	case 1:
		playing = true
		ready = true

	case -1, 0, 2:
		playing = false
		ready = true

	case 3:
		ready = false
	}

	p.Defer(func(ctx app.Context) {
		p.ready = ready
		p.playing = playing
		p.Update()
	})

	return nil
}

func (p *player) play() {
	p.youtube.Call("playVideo")
}

func (p *player) setVolume(ctx app.Context, volume int) {
	if volume == 0 {
		p.youtube.Call("mute")
	} else {
		p.youtube.Call("unMute")
		p.State.LastNonZeroVolume = volume
	}

	p.youtube.Call("setVolume", volume)
	p.State.Volume = volume
	p.saveState(ctx)
	p.Update()

	p.Defer(func(ctx app.Context) {
		app.Window().GetElementByID("volume-bar").Set("value", volume)
	})
}

func (p *player) saveState(ctx app.Context) {
	if err := p.State.save(ctx); err != nil {
		app.Logf("%s", errors.New("saving player state failed").Wrap(err))
	}
}

func (p *player) Render() app.UI {
	hide := "hide"
	if p.ready {
		hide = ""
	}

	blur := "blur"
	if p.playing {
		blur = ""
	}

	return app.Div().
		ID("player").
		Class("player").
		Body(
			app.Div().
				Class("video").
				Class(blur).
				Body(
					app.Script().
						Src("//www.youtube.com/iframe_api").
						Async(true),
					app.IFrame().
						ID("youtube-"+p.Channel.Slug).
						Allow("autoplay").
						Allow("accelerometer").
						Allow("encrypted-media").
						Allow("picture-in-picture").
						Sandbox("allow-presentation allow-same-origin allow-scripts allow-popups").
						Src(fmt.Sprintf(
							"https://www.youtube.com/embed/%s?controls=0&showinfo=0&autoplay=1&loop=1&enablejsapi=1&playsinline=1",
							p.Channel.ID,
						)),
				),
			app.Div().
				Class("overlay").
				Body(
					app.Footer().Body(
						app.Stack().
							Class("controls").
							Class(hide).
							Center().
							Content(
								app.If(!p.playing,
									app.Button().
										ID("play").
										Class("button").
										Title(fmt.Sprintf("Play %s.", p.Channel.Name)).
										OnClick(p.onPlay).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
												<path fill="currentColor" d="M8,5.14V19.14L19,12.14L8,5.14Z" />
											</svg>
											`),
										),
								).Else(
									app.Button().
										ID("shuffle").
										Class("button").
										Title("Play a random Lofi channel.").
										OnClick(p.onShuffle).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    											<path fill="currentColor" d="M14.83,13.41L13.42,14.82L16.55,17.95L14.5,20H20V14.5L17.96,16.54L14.83,13.41M14.5,4L16.54,6.04L4,18.59L5.41,20L17.96,7.46L20,9.5V4M10.59,9.17L5.41,4L4,5.41L9.17,10.58L10.59,9.17Z" />
											</svg>
											`),
										),
								),
							),

						app.Stack().
							Class("volume").
							Class(hide).
							Center().
							Content(
								app.If(p.State.Volume > 66,
									app.Button().
										Class("button").
										Title("Mute volume.").
										OnClick(p.onMute).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    											<path fill="currentColor" d="M14,3.23V5.29C16.89,6.15 19,8.83 19,12C19,15.17 16.89,17.84 14,18.7V20.77C18,19.86 21,16.28 21,12C21,7.72 18,4.14 14,3.23M16.5,12C16.5,10.23 15.5,8.71 14,7.97V16C15.5,15.29 16.5,13.76 16.5,12M3,9V15H7L12,20V4L7,9H3Z" />
											</svg>
											`),
										),
								).ElseIf(p.State.Volume > 33,
									app.Button().
										Class("button").
										Title("Mute volume.").
										OnClick(p.onMute).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
												<path fill="currentColor" d="M5,9V15H9L14,20V4L9,9M18.5,12C18.5,10.23 17.5,8.71 16,7.97V16C17.5,15.29 18.5,13.76 18.5,12Z" />
											</svg>
											`),
										),
								).ElseIf(p.State.Volume > 0,
									app.Button().
										Class("button").
										Title("Mute volume.").
										OnClick(p.onMute).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
												<path fill="currentColor" d="M7,9V15H11L16,20V4L11,9H7Z" />
											</svg>
											`),
										),
								).Else(
									app.Button().
										Class("button").
										Title("Unmute volume.").
										OnClick(p.onUnMute).
										Body(
											app.Raw(`
											<svg style="width:24px;height:24px" viewBox="0 0 24 24">
    											<path fill="currentColor" d="M3,9H7L12,4V20L7,15H3V9M16.59,12L14,9.41L15.41,8L18,10.59L20.59,8L22,9.41L19.41,12L22,14.59L20.59,16L18,13.41L15.41,16L14,14.59L16.59,12Z" />
											</svg>
											`),
										),
								),
								app.Input().
									ID("volume-bar").
									Class("volumebar").
									Type("range").
									Placeholder("Volume").
									Min("0").
									Max("100").
									Value(strconv.Itoa(p.State.Volume)).
									OnChange(p.onVolumeChange).
									OnInput(p.onVolumeChange),
							),
					),
				),
		)
}

func (p *player) onPlay(ctx app.Context, e app.Event) {
	p.play()
}

func (p *player) onVolumeChange(ctx app.Context, e app.Event) {
	volume, _ := strconv.Atoi(ctx.JSSrc.Get("value").String())
	p.setVolume(ctx, volume)
}

func (p *player) onMute(ctx app.Context, e app.Event) {
	p.setVolume(ctx, 0)
}

func (p *player) onUnMute(ctx app.Context, e app.Event) {
	p.setVolume(ctx, p.State.LastNonZeroVolume)
}

func (p *player) onShuffle(ctx app.Context, e app.Event) {
	for {
		c := channels.Get("/")
		if c.Slug != p.Channel.Slug {
			ctx.Navigate("/" + c.Slug)
			return
		}
	}
}

type playerState struct {
	Volume            int
	LastNonZeroVolume int
	Muted             bool
}

func (s *playerState) load(ctx app.Context) error {
	if err := ctx.LocalStorage().Get("player.state", s); err != nil {
		return errors.New("getting player status from local storage failed").
			Wrap(err)
	}

	if *s == (playerState{}) {
		s.Volume = 100
		s.LastNonZeroVolume = 100
	}

	return nil
}

func (s *playerState) save(ctx app.Context) error {
	if err := ctx.LocalStorage().Set("player.state", s); err != nil {
		return errors.New("saving player status in local storage failed").
			Wrap(err)
	}

	return nil
}
