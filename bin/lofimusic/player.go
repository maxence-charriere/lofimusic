package main

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

var (
	youtubeAPIReady bool
)

type player struct {
	app.Compo

	Channel channel

	youtube                  app.Value
	releaseIframe            func()
	releaseOnPlayerReady     func()
	releasePlayerStateChange func()
	ready                    bool
	playing                  bool
	volume                   int
}

func (p *player) OnMount(ctx app.Context) {
	if !youtubeAPIReady {
		onYouTubeIframeAPIReady := app.FuncOf(p.onYoutubeIframeAPIReady)
		p.releaseIframe = onYouTubeIframeAPIReady.Release
		app.Window().Set("onYouTubeIframeAPIReady", onYouTubeIframeAPIReady)
		return
	}

	app.Dispatch(p.setupYoutubePlayer)
	p.Update()
}

func (p *player) onYoutubeIframeAPIReady(this app.Value, args []app.Value) interface{} {
	app.Dispatch(func() {
		youtubeAPIReady = true
		p.setupYoutubePlayer()
	})

	return nil
}

func (p *player) setupYoutubePlayer() {
	onPlayerReady := app.FuncOf(p.onPlayerReady)
	p.releaseOnPlayerReady = onPlayerReady.Release

	onPlayerStateChange := app.FuncOf(p.onPlayerStateChange)
	p.releasePlayerStateChange = onPlayerStateChange.Release

	p.volume = 50
	app.LocalStorage.Get("volume", &p.volume)

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

func (p *player) onPlayerReady(this app.Value, args []app.Value) interface{} {
	p.youtube.Call("setVolume", p.volume)
	p.youtube.Call("playVideo")
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

	app.Dispatch(func() {
		p.ready = ready
		p.playing = playing
		p.Update()
	})

	return nil
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
								app.Raw(
									`<svg style="width:24px;height:24px" viewBox="0 0 24 24">
										<path fill="currentColor" d="M14,3.23V5.29C16.89,6.15 19,8.83 19,12C19,15.17 16.89,17.84 14,18.7V20.77C18,19.86 21,16.28 21,12C21,7.72 18,4.14 14,3.23M16.5,12C16.5,10.23 15.5,8.71 14,7.97V16C15.5,15.29 16.5,13.76 16.5,12M3,9V15H7L12,20V4L7,9H3Z" />
									</svg>
									`),
								app.Input().
									Type("range").
									Placeholder("Volume").
									Min("0").
									Max("100").
									Value(strconv.Itoa(p.volume)).
									OnInput(p.onVolumeChange).
									OnClick(p.onVolumeChange),
							),
					),
				),
		)
}

func (p *player) onPlay(ctx app.Context, e app.Event) {
	p.play()
}

func (p *player) play() {
	p.youtube.Call("playVideo")
}

func (p *player) onVolumeChange(ctx app.Context, e app.Event) {
	volume, _ := strconv.Atoi(ctx.JSSrc.Get("value").String())
	p.volume = volume
	p.youtube.Call("setVolume", p.volume)
	app.LocalStorage.Set("volume", p.volume)
}

func (p *player) onShuffle(ctx app.Context, e app.Event) {
	for {
		c := channels.Get("/")
		if c.Slug != p.Channel.Slug {
			app.Navigate("/" + c.Slug)
			return
		}
	}
}
