package radio

import (
	"path"
)

// Video represents a video to play.
type Video struct {
	Slug  string
	Name  string
	Owner string
	URL   string
	Cards []string

	links map[string]string
}

func (v Video) Links() []Link {
	links := make([]Link, 0, len(v.links)+1)

	for k, v := range v.links {
		links = append(links, Link{
			Slug: k,
			URL:  v,
		})
	}

	return append(links, Link{
		Slug: "youtube",
		URL:  v.URL,
	})
}

func (v Video) YoutubeID() string {
	return path.Base(v.URL)
}

// Link represents a link to a media related to a video.
type Link struct {
	Slug string
	URL  string
}
