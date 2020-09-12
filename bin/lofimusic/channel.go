package main

import (
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

var (
	channels = newChannelStore()
)

type channel struct {
	Name        string
	ID          string
	Slug        string
	Description string
	SocialMedia []link
}

type link struct {
	URL       string
	MediaSlug string
}

type channelStore struct {
	channels []channel
}

func newChannelStore() *channelStore {
	rand.Seed(uint64(time.Now().UnixNano()))

	c := []channel{
		{
			Name: "ChilledCow - relax & study",
			ID:   "5qap5aO4i9A",
			Slug: "chilledcow",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/5qap5aO4i9A",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://chilledcow.com/password",
					MediaSlug: "web",
				},
				{
					URL:       "https://www.instagram.com/chilledcow_yt",
					MediaSlug: "instagram",
				},
				{
					URL:       "https://www.facebook.com/TheChilledCow",
					MediaSlug: "facebook",
				},
				{
					URL:       "https://twitter.com/chilledcow",
					MediaSlug: "twitter",
				},
				{
					URL:       "https://discord.com/invite/hUKvJnw",
					MediaSlug: "discord",
				},
			},
		},
		{
			Name: "ChilledCow - sleep & chill",
			ID:   "DWcJFNfaw9c",
			Slug: "chilledcow2",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/DWcJFNfaw9c",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://chilledcow.com/password",
					MediaSlug: "web",
				},
				{
					URL:       "https://www.instagram.com/chilledcow_yt",
					MediaSlug: "instagram",
				},
				{
					URL:       "https://www.facebook.com/TheChilledCow",
					MediaSlug: "facebook",
				},
				{
					URL:       "https://twitter.com/chilledcow",
					MediaSlug: "twitter",
				},
				{
					URL:       "https://discord.com/invite/hUKvJnw",
					MediaSlug: "discord",
				},
			},
		},
		{
			Name: "Chillhop Music - jazzy & hip hop",
			ID:   "5yx6BWlEVcY",
			Slug: "chillhopmusic",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/DWcJFNfaw9c",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://chillhop.com",
					MediaSlug: "web",
				},
				{
					URL:       "https://www.instagram.com/chillhopmusic",
					MediaSlug: "instagram",
				},
				{
					URL:       "https://discord.com/invite/chillhop",
					MediaSlug: "discord",
				},
				{
					URL:       "https://www.reddit.com/r/chillhop",
					MediaSlug: "reddit",
				},
			},
		},
		{
			Name: "Chillhop Music - study & relax",
			ID:   "7NOSDKb0HlU",
			Slug: "chillhopmusic2",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/7NOSDKb0HlU",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://chillhop.com",
					MediaSlug: "web",
				},
				{
					URL:       "https://www.instagram.com/chillhopmusic",
					MediaSlug: "instagram",
				},
				{
					URL:       "https://discord.com/invite/chillhop",
					MediaSlug: "discord",
				},
				{
					URL:       "https://www.reddit.com/r/chillhop",
					MediaSlug: "reddit",
				},
			},
		},
		{
			Name: "Tokyo LosT Tracks - サクラチル",
			ID:   "WBfbkPTqUtU",
			Slug: "tokyolosttracks",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/WBfbkPTqUtU",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/playlist/37i9dQZF1DX87D1EaNZxW1",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://discord.com/invite/fHkhEgc",
					MediaSlug: "discord",
				},
				{
					URL:       "https://twitter.com/TTT_SakuraChill",
					MediaSlug: "twitter",
				},
				{
					URL:       "https://www.facebook.com/TokyoLosTTracks",
					MediaSlug: "facebook",
				},
				{
					URL:       "https://www.instagram.com/TokyoLosTTracks/",
					MediaSlug: "instagram",
				},
			},
		},
		{
			Name: "The Jazz Hop Café",
			ID:   "OVPPOwMpSpQ",
			Slug: "thejazzhopcafe",
			SocialMedia: []link{
				{
					URL:       "https://youtu.be/OVPPOwMpSpQ",
					MediaSlug: "youtube",
				},
				{
					URL:       "https://open.spotify.com/user/thejazz",
					MediaSlug: "spotify",
				},
				{
					URL:       "https://twitter.com/jazzhopcafe",
					MediaSlug: "twitter",
				},
				{
					URL:       "https://www.facebook.com/jazzhopcafe",
					MediaSlug: "facebook",
				},
			},
		},
	}

	sort.Slice(c, func(a, b int) bool {
		return strings.Compare(c[a].Name, c[b].Name) < 0
	})

	for i := range c {
		sort.Slice(c[i].SocialMedia, func(a, b int) bool {
			return strings.Compare(c[i].SocialMedia[a].MediaSlug, c[i].SocialMedia[b].MediaSlug) < 0
		})

	}

	return &channelStore{
		channels: c,
	}
}

func (s *channelStore) Get(slug string) channel {
	for _, c := range s.channels {
		if c.Slug == slug {
			return c
		}
	}

	idx := rand.Intn(len(s.channels))
	return s.channels[idx]
}

func (s *channelStore) Channels() []channel {
	c := make([]channel, len(s.channels))
	copy(c, s.channels)
	return c
}

func (s *channelStore) Slugs() []string {
	slugs := make([]string, len(s.channels))

	for i, c := range s.channels {
		slugs[i] = c.Slug
	}

	return slugs
}

type media struct {
	Name string
	Icon string
}

var medias = map[string]media{
	"youtube": {
		Name: "YouTube",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M10,15L15.19,12L10,9V15M21.56,7.17C21.69,7.64 21.78,8.27 21.84,9.07C21.91,9.87 21.94,10.56 21.94,11.16L22,12C22,14.19 21.84,15.8 21.56,16.83C21.31,17.73 20.73,18.31 19.83,18.56C19.36,18.69 18.5,18.78 17.18,18.84C15.88,18.91 14.69,18.94 13.59,18.94L12,19C7.81,19 5.2,18.84 4.17,18.56C3.27,18.31 2.69,17.73 2.44,16.83C2.31,16.36 2.22,15.73 2.16,14.93C2.09,14.13 2.06,13.44 2.06,12.84L2,12C2,9.81 2.16,8.2 2.44,7.17C2.69,6.27 3.27,5.69 4.17,5.44C4.64,5.31 5.5,5.22 6.82,5.16C8.12,5.09 9.31,5.06 10.41,5.06L12,5C16.19,5 18.8,5.16 19.83,5.44C20.73,5.69 21.31,6.27 21.56,7.17Z" />
		</svg>
		`,
	},
	"facebook": {
		Name: "Facebook",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M12 2.04C6.5 2.04 2 6.53 2 12.06C2 17.06 5.66 21.21 10.44 21.96V14.96H7.9V12.06H10.44V9.85C10.44 7.34 11.93 5.96 14.22 5.96C15.31 5.96 16.45 6.15 16.45 6.15V8.62H15.19C13.95 8.62 13.56 9.39 13.56 10.18V12.06H16.34L15.89 14.96H13.56V21.96A10 10 0 0 0 22 12.06C22 6.53 17.5 2.04 12 2.04Z" />
		</svg>
		`,
	},
	"instagram": {
		Name: "Instagram",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M7.8,2H16.2C19.4,2 22,4.6 22,7.8V16.2A5.8,5.8 0 0,1 16.2,22H7.8C4.6,22 2,19.4 2,16.2V7.8A5.8,5.8 0 0,1 7.8,2M7.6,4A3.6,3.6 0 0,0 4,7.6V16.4C4,18.39 5.61,20 7.6,20H16.4A3.6,3.6 0 0,0 20,16.4V7.6C20,5.61 18.39,4 16.4,4H7.6M17.25,5.5A1.25,1.25 0 0,1 18.5,6.75A1.25,1.25 0 0,1 17.25,8A1.25,1.25 0 0,1 16,6.75A1.25,1.25 0 0,1 17.25,5.5M12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9Z" />
		</svg>
		`,
	},
	"twitter": {
		Name: "Twitter",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M22.46,6C21.69,6.35 20.86,6.58 20,6.69C20.88,6.16 21.56,5.32 21.88,4.31C21.05,4.81 20.13,5.16 19.16,5.36C18.37,4.5 17.26,4 16,4C13.65,4 11.73,5.92 11.73,8.29C11.73,8.63 11.77,8.96 11.84,9.27C8.28,9.09 5.11,7.38 3,4.79C2.63,5.42 2.42,6.16 2.42,6.94C2.42,8.43 3.17,9.75 4.33,10.5C3.62,10.5 2.96,10.3 2.38,10C2.38,10 2.38,10 2.38,10.03C2.38,12.11 3.86,13.85 5.82,14.24C5.46,14.34 5.08,14.39 4.69,14.39C4.42,14.39 4.15,14.36 3.89,14.31C4.43,16 6,17.26 7.89,17.29C6.43,18.45 4.58,19.13 2.56,19.13C2.22,19.13 1.88,19.11 1.54,19.07C3.44,20.29 5.7,21 8.12,21C16,21 20.33,14.46 20.33,8.79C20.33,8.6 20.33,8.42 20.32,8.23C21.16,7.63 21.88,6.87 22.46,6Z" />
		</svg>
		`,
	},
	"spotify": {
		Name: "Spotify",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M17.9,10.9C14.7,9 9.35,8.8 6.3,9.75C5.8,9.9 5.3,9.6 5.15,9.15C5,8.65 5.3,8.15 5.75,8C9.3,6.95 15.15,7.15 18.85,9.35C19.3,9.6 19.45,10.2 19.2,10.65C18.95,11 18.35,11.15 17.9,10.9M17.8,13.7C17.55,14.05 17.1,14.2 16.75,13.95C14.05,12.3 9.95,11.8 6.8,12.8C6.4,12.9 5.95,12.7 5.85,12.3C5.75,11.9 5.95,11.45 6.35,11.35C10,10.25 14.5,10.8 17.6,12.7C17.9,12.85 18.05,13.35 17.8,13.7M16.6,16.45C16.4,16.75 16.05,16.85 15.75,16.65C13.4,15.2 10.45,14.9 6.95,15.7C6.6,15.8 6.3,15.55 6.2,15.25C6.1,14.9 6.35,14.6 6.65,14.5C10.45,13.65 13.75,14 16.35,15.6C16.7,15.75 16.75,16.15 16.6,16.45M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
		</svg>
		`,
	},
	"discord": {
		Name: "Discord",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M22,24L16.75,19L17.38,21H4.5A2.5,2.5 0 0,1 2,18.5V3.5A2.5,2.5 0 0,1 4.5,1H19.5A2.5,2.5 0 0,1 22,3.5V24M12,6.8C9.32,6.8 7.44,7.95 7.44,7.95C8.47,7.03 10.27,6.5 10.27,6.5L10.1,6.33C8.41,6.36 6.88,7.53 6.88,7.53C5.16,11.12 5.27,14.22 5.27,14.22C6.67,16.03 8.75,15.9 8.75,15.9L9.46,15C8.21,14.73 7.42,13.62 7.42,13.62C7.42,13.62 9.3,14.9 12,14.9C14.7,14.9 16.58,13.62 16.58,13.62C16.58,13.62 15.79,14.73 14.54,15L15.25,15.9C15.25,15.9 17.33,16.03 18.73,14.22C18.73,14.22 18.84,11.12 17.12,7.53C17.12,7.53 15.59,6.36 13.9,6.33L13.73,6.5C13.73,6.5 15.53,7.03 16.56,7.95C16.56,7.95 14.68,6.8 12,6.8M9.93,10.59C10.58,10.59 11.11,11.16 11.1,11.86C11.1,12.55 10.58,13.13 9.93,13.13C9.29,13.13 8.77,12.55 8.77,11.86C8.77,11.16 9.28,10.59 9.93,10.59M14.1,10.59C14.75,10.59 15.27,11.16 15.27,11.86C15.27,12.55 14.75,13.13 14.1,13.13C13.46,13.13 12.94,12.55 12.94,11.86C12.94,11.16 13.45,10.59 14.1,10.59Z" />
		</svg>
		`,
	},
	"reddit": {
		Name: "Reddit",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
    		<path fill="currentColor" d="M14.5 15.41C14.58 15.5 14.58 15.69 14.5 15.8C13.77 16.5 12.41 16.56 12 16.56C11.61 16.56 10.25 16.5 9.54 15.8C9.44 15.69 9.44 15.5 9.54 15.41C9.65 15.31 9.82 15.31 9.92 15.41C10.38 15.87 11.33 16 12 16C12.69 16 13.66 15.87 14.1 15.41C14.21 15.31 14.38 15.31 14.5 15.41M10.75 13.04C10.75 12.47 10.28 12 9.71 12C9.14 12 8.67 12.47 8.67 13.04C8.67 13.61 9.14 14.09 9.71 14.08C10.28 14.08 10.75 13.61 10.75 13.04M14.29 12C13.72 12 13.25 12.5 13.25 13.05S13.72 14.09 14.29 14.09C14.86 14.09 15.33 13.61 15.33 13.05C15.33 12.5 14.86 12 14.29 12M22 12C22 17.5 17.5 22 12 22S2 17.5 2 12C2 6.5 6.5 2 12 2S22 6.5 22 12M18.67 12C18.67 11.19 18 10.54 17.22 10.54C16.82 10.54 16.46 10.7 16.2 10.95C15.2 10.23 13.83 9.77 12.3 9.71L12.97 6.58L15.14 7.05C15.16 7.6 15.62 8.04 16.18 8.04C16.75 8.04 17.22 7.57 17.22 7C17.22 6.43 16.75 5.96 16.18 5.96C15.77 5.96 15.41 6.2 15.25 6.55L12.82 6.03C12.75 6 12.68 6.03 12.63 6.07C12.57 6.11 12.54 6.17 12.53 6.24L11.79 9.72C10.24 9.77 8.84 10.23 7.82 10.96C7.56 10.71 7.2 10.56 6.81 10.56C6 10.56 5.35 11.21 5.35 12C5.35 12.61 5.71 13.11 6.21 13.34C6.19 13.5 6.18 13.62 6.18 13.78C6.18 16 8.79 17.85 12 17.85C15.23 17.85 17.85 16.03 17.85 13.78C17.85 13.64 17.84 13.5 17.81 13.34C18.31 13.11 18.67 12.6 18.67 12Z" />
		</svg>
		`,
	},
	"web": {
		Name: "Website",
		Icon: `
		<svg style="width:18px;height:18px" viewBox="0 0 24 24">
			<path fill="currentColor" d="M16.36,14C16.44,13.34 16.5,12.68 16.5,12C16.5,11.32 16.44,10.66 16.36,10H19.74C19.9,10.64 20,11.31 20,12C20,12.69 19.9,13.36 19.74,14M14.59,19.56C15.19,18.45 15.65,17.25 15.97,16H18.92C17.96,17.65 16.43,18.93 14.59,19.56M14.34,14H9.66C9.56,13.34 9.5,12.68 9.5,12C9.5,11.32 9.56,10.65 9.66,10H14.34C14.43,10.65 14.5,11.32 14.5,12C14.5,12.68 14.43,13.34 14.34,14M12,19.96C11.17,18.76 10.5,17.43 10.09,16H13.91C13.5,17.43 12.83,18.76 12,19.96M8,8H5.08C6.03,6.34 7.57,5.06 9.4,4.44C8.8,5.55 8.35,6.75 8,8M5.08,16H8C8.35,17.25 8.8,18.45 9.4,19.56C7.57,18.93 6.03,17.65 5.08,16M4.26,14C4.1,13.36 4,12.69 4,12C4,11.31 4.1,10.64 4.26,10H7.64C7.56,10.66 7.5,11.32 7.5,12C7.5,12.68 7.56,13.34 7.64,14M12,4.03C12.83,5.23 13.5,6.57 13.91,8H10.09C10.5,6.57 11.17,5.23 12,4.03M18.92,8H15.97C15.65,6.75 15.19,5.55 14.59,4.44C16.43,5.07 17.96,6.34 18.92,8M12,2C6.47,2 2,6.5 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
		</svg>
		`,
	},
}
