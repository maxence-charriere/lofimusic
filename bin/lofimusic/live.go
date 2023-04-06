package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "everything-fades-to-blue",
			Name:  "Everything Fades To Blue",
			Owner: "Sleepy Fish",
			URL:   "https://youtu.be/PfgS405CdXk",
			Cards: []string{
				"Everything Fades To Blue is a mix of indie/emo and ambient music produced by Sleep Fish, a Pennsylvania-based producer that is also a student in statistics, data science, and machine learning.",
				"Sleepy Fish is one of the few Lo-fi artists who actually sing in its creations.",
				"Everything Fades To Blue is the 3rd episode of a story where a tidal wave destroys an island along with the home where Sleepy Fish used to live.",
				"Toppled into the sea, on its own for the first time, Sleepy Fish uses its glow to search for family, to guide others, and to find its way.",
				`The episode comes after "My Room Becomes the Sea" and "Beneath Your Waves".`,
				"The undersea-themed animation has been made in collaboration with Tristan Gion and Bien à Vous Studio, a French studio based in Nantes.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com/releases/sleepy-fish-everything-fades-to-blue",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/PfgS405CdXk",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/artist/1IJe80moz409PtxW4llPFw",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/sleepyfishmusic",
				},
			},
		},
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/jfKfPfyJRdk",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is shown studying in Lyon, a city from France where her designer Juan Pablo used to live.",
				"The view through the window depicts the buildings on the slopes of Croix-Rousse, where the bell tower of the Bon-Pasteur church can be spotted.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/jfKfPfyJRdk",
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
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/rUxyKA_-grg",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is living in Lyon, a city from France where her designer Juan Pablo used to live.",
			},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://lofigirl.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/rUxyKA_-grg",
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
			Owner: "Chillhop Music",
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
			Owner: "Chillhop Music",
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
			Owner: "College Music",
			URL:   "https://youtu.be/3TASKrR6nrg",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/3TASKrR6nrg",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
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
			Owner: "College Music",
			URL:   "https://youtu.be/QwXHcgZUnFI",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/QwXHcgZUnFI",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
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
			Owner: "College Music",
			URL:   "https://youtu.be/dxUtV-zNv9w",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.collegemusic.co.uk/",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/dxUtV-zNv9w",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
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
		// {
		// 	Slug:  "lofi-code-beats",
		// 	Name:  "Coding Beats",
		// 	Owner: "Joma Tech",
		// 	URL:   "https://youtu.be/PY8f1Z3nARo",
		// 	Cards: []string{},
		// 	Links: []socialLink{
		// 		{
		// 			Slug: "youtube",
		// 			URL:  "https://youtu.be/PY8f1Z3nARo",
		// 		},
		// 		{
		// 			Slug: "instagram",
		// 			URL:  "https://www.instagram.com/jomaoppa",
		// 		},
		// 		{
		// 			Slug: "facebook",
		// 			URL:  "https://www.facebook.com/jomaoppa",
		// 		},
		// 		{
		// 			Slug: "twitter",
		// 			URL:  "https://twitter.com/jomaoppa",
		// 		},
		// 	},
		// },
		// {
		// 	Slug:  "steezyasfuck-coffee-show",
		// 	Name:  "Steezy Coffee Shop",
		// 	Owner: "STEEZYASFUCK",
		// 	URL:   "https://youtu.be/-5KAN9_CzSA",
		// 	Cards: []string{},
		// 	Links: []socialLink{
		// 		{
		// 			Slug: "website",
		// 			URL:  "https://www.stzzzy.com",
		// 		},
		// 		{
		// 			Slug: "youtube",
		// 			URL:  "https://youtu.be/-5KAN9_CzSA",
		// 		},
		// 		{
		// 			Slug: "spotify",
		// 			URL:  "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
		// 		},
		// 		{
		// 			Slug: "instagram",
		// 			URL:  "https://www.instagram.com/stzzyasfvck/",
		// 		},
		// 		{
		// 			Slug: "twitter",
		// 			URL:  "https://twitter.com/stzzyasfvck",
		// 		},
		// 	},
		// },
		// {
		// 	Slug:  "steezyasfuck-junky-fluff",
		// 	Name:  "Steezy Junky Fluff",
		// 	Owner: "STEEZYASFUCK",
		// 	URL:   "https://www.youtube.com/watch?v=xgirCNccI68",
		// 	Cards: []string{},
		// 	Links: []socialLink{
		// 		{
		// 			Slug: "website",
		// 			URL:  "https://www.youtube.com/watch?v=xgirCNccI68",
		// 		},
		// 		{
		// 			Slug: "youtube",
		// 			URL:  "https://youtu.be/-5KAN9_CzSA",
		// 		},
		// 		{
		// 			Slug: "spotify",
		// 			URL:  "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
		// 		},
		// 		{
		// 			Slug: "instagram",
		// 			URL:  "https://www.instagram.com/stzzyasfvck/",
		// 		},
		// 		{
		// 			Slug: "twitter",
		// 			URL:  "https://twitter.com/stzzyasfvck",
		// 		},
		// 	},
		// },
		{
			Slug:  "closedonsunday-pop-culture",
			Name:  "Pop Culture Sunday",
			Owner: "Closed on Sunday",
			URL:   "https://youtu.be/pixAeRe2rk0",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/pixAeRe2rk0",
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
			Owner: "Closed on Sunday",
			URL:   "https://youtu.be/zFTFx6nclUY",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/zFTFx6nclUY",
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
			Slug:  "thebootlegboy",
			Name:  "Bootleg Smoke",
			Owner: "the bootleg boy",
			URL:   "https://youtu.be/bLlloaA4b4g",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.thebootlegboy.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/bLlloaA4b4g",
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
		// {
		// 	Slug:  "sleepynatula",
		// 	Name:  "Sleepy Natula",
		// 	Owner: "Tone by Gridge",
		// 	URL:   "https://youtu.be/pCuy6U-RVFk",
		// 	Cards: []string{
		// 		"Sleepy Natula is a radio that broadcasts Neo chill beats style songs, produced by Tone by Gridge in the Japanese city of Tokyo.",
		// 		"The animation shows Natula, a girl that fell asleep after drinking a relaxation drink named CHILL OUT.",
		// 		"Natula lives in Sephora, a city located on a musical planet named Tone.",
		// 		"On this planet, music is so important that it touches every species, nature, climate, culture, and civilization of each country.",
		// 		"Tone is a multicultural planet where there are thousands of countries populated by music-loving beings that come from all over the galaxy.",
		// 	},
		// 	Links: []socialLink{
		// 		{
		// 			Slug: "website",
		// 			URL:  "https://linktr.ee/tonebygridge?fbclid=IwAR2OYWJlav7MM66780DDtuu2k2viElVF3kgjn-GEK8c2RsvYt0TGOYz31zQ",
		// 		},
		// 		{
		// 			Slug: "youtube",
		// 			URL:  "https://youtu.be/pCuy6U-RVFk",
		// 		},
		// 		{
		// 			Slug: "spotify",
		// 			URL:  "https://open.spotify.com/playlist/1stD3Nr9W5HWf47ft5TfKh?si=fDM2-nsIRlSdv2IC_KaanQ",
		// 		},
		// 		{
		// 			Slug: "discord",
		// 			URL:  "https://discord.com/invite/eqH6RRj",
		// 		},
		// 		{
		// 			Slug: "facebook",
		// 			URL:  "https://www.facebook.com/tonedotcom",
		// 		},
		// 		{
		// 			Slug: "instagram",
		// 			URL:  "https://www.instagram.com/tonebygridge",
		// 		},
		// 		{
		// 			Slug: "twitter",
		// 			URL:  "https://twitter.com/tonedotcom",
		// 		},
		// 	},
		// },
		{
			Slug:  "dreamhop",
			Name:  "Dreamhop",
			Owner: "Dreamhop Music",
			URL:   "https://youtu.be/wkhLHTmS_GI",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://dreamhopmusic.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/wkhLHTmS_GI",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/user/91jfqzlv0htrqrvvc60138qma",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/FxF9kng",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/dreamhopp",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Dreamhopp",
				},
			},
		},
		{
			Slug:  "taiki",
			Name:  "Taiki",
			Owner: "Chill with Taiki",
			URL:   "https://youtu.be/qH3fETPsqXU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/qH3fETPsqXU",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/2qvd2ngQGP",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/chillwithtaiki",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/chillwithtaiki",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillwithtaiki",
				},
				{
					Slug: "website",
					URL:  "https://taiki.shop",
				},
			},
		},
		{
			Slug:  "house-in-the-woods",
			Name:  "House in the Woods",
			Owner: "Lofi Zone",
			URL:   "https://youtu.be/MKzeyoX_w4g",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/MKzeyoX_w4g",
				},
				{
					Slug: "spotify",
					URL:  "https://open.spotify.com/playlist/0V5IsHm0VJbmeffuLzgoc3?si=eac2c90bc18f4a97",
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
