package pwa

import (
	"math/rand"
	"path"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/lofimusic/pkg/radio"
)

const (
	loadVideo = "/player/video/load"

	currentVideoState = "/player/video/current"
)

var (
	videos = newVideoStore()
)

func handleLoadVideo(ctx app.Context, a app.Action) {
	var v radio.Video

	if slug := videoSlugFromPath(a.Tags.Get("path")); slug == "" {
		leastWatchedVideos := videos.leastWatchedVideos()
		v = leastWatchedVideos[rand.Intn(len(leastWatchedVideos))]
	} else {
		v = videos.videoBySlug(slug)
	}

	videos.inc(v)
	ctx.SetState(currentVideoState, v)
}

func videoSlugFromPath(v string) string {
	slug := path.Base(v)
	slug = strings.TrimLeft(slug, "/")
	slug = strings.TrimLeft(slug, ".")
	return slug
}
