package main

import (
	"path"
	"sort"
	"strings"
)

type liveRadio struct {
	Slug  string
	Name  string
	URL   string
	Cards []string
	Links []socialLink
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			URL:   "https://youtu.be/5qap5aO4i9A",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5qap5aO4i9A",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/lofigirl",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/lofigirl",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lofigirl",
				},
			},
		},
		{
			Slug:  "lofigirl-sleepy",
			Name:  "Lofi Sleepy Girl",
			URL:   "https://youtu.be/DWcJFNfaw9c",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/DWcJFNfaw9c",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/lofigirl",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/lofigirl",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lofigirl",
				},
			},
		},
		{
			Slug:  "chillhop",
			Name:  "Chillhop Raccoon",
			URL:   "https://youtu.be/5yx6BWlEVcY",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:  "chillhop-relax",
			Name:  "Chillhop Relaxing Raccoon",
			URL:   "https://youtu.be/7NOSDKb0HlU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/7NOSDKb0HlU",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:  "collegemusic",
			Name:  "College Girl",
			URL:   "https://youtu.be/MCkTebktHVc",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/MCkTebktHVc",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
			Slug:  "collegemusic-guy",
			Name:  "College Guy",
			URL:   "https://youtu.be/2atQnvunGCo",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/2atQnvunGCo",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
			Slug:  "collegemusic-lonely",
			Name:  "College Lonely",
			URL:   "https://youtu.be/bM0Iw7PPoU4",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/bM0Iw7PPoU4",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/collegemusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/collegemusic",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/collegemusicyt",
				},
			},
		},
		{
			Slug:  "lofi-code-beats",
			Name:  "Coding Beats",
			URL:   "https://youtu.be/bmVKaAV_7-A",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/bmVKaAV_7-A",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jomaoppa",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/jomaoppa",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/jomaoppa",
				},
			},
		},
		{
			Slug:  "steezyasfuck-coffee-show",
			Name:  "Steezy Coffee Shop",
			URL:   "https://youtu.be/-5KAN9_CzSA",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.stzzzy.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/-5KAN9_CzSA",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/stzzyasfvck/",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/stzzyasfvck",
				},
			},
		},
		{
			Slug:  "steezyasfuck-junky-fluff",
			Name:  "Steezy Junky Fluff",
			URL:   "https://youtu.be/rc9cjjEun_k",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.stzzzy.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/rc9cjjEun_k",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/lofigirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/stzzyasfvck/",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/stzzyasfvck",
				},
			},
		},
		{
			Slug:  "closedonsunday-pop-culture",
			Name:  "Pop Culture Sunday",
			URL:   "https://youtu.be/mOe8VEMuPo0",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/mOe8VEMuPo0",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/artist/1LwjR2mIm78OJRTYdkMLl3",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/closedonsunday",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/closedonsundayy",
				},
			},
		},
		{
			Slug:  "closedonsunday-starwars",
			Name:  "Star Wars Sunday",
			URL:   "https://youtu.be/o33l32ZrIy8",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/o33l32ZrIy8",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/artist/1LwjR2mIm78OJRTYdkMLl3",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/closedonsunday",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/closedonsundayy",
				},
			},
		},
		{
			Slug:  "tokyolosttracks",
			Name:  "サクラチル",
			URL:   "https://youtu.be/WBfbkPTqUtU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/WBfbkPTqUtU",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/37i9dQZF1DX87D1EaNZxW1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/fHkhEgc",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/LofiGirl",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/TokyoLosTTracks",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/TokyoLosTTracks",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/TTT_SakuraChill",
				},
			},
		},
		{
			Slug:  "thebootlegboy",
			Name:  "Bootleg Smoke",
			URL:   "https://youtu.be/l7TxwBhtTUY",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.thebootlegboy.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/l7TxwBhtTUY",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/71019EDcRamfMmOEEoTdEu?si=XePP-REWQDSuzJT6-SXwSQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/FZrUkey",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/thebootlegboy",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/thebootlegboy",
				},
			},
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
