package radio

import (
	"sort"
	"strings"
)

// List returns the available videos.
func List() []Video {
	videos := []Video{
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
			links: map[string]string{
				"website":   "https://chillhop.com/releases/sleepy-fish-everything-fades-to-blue",
				"spotify":   "https://open.spotify.com/artist/1IJe80moz409PtxW4llPFw",
				"instagram": "https://www.instagram.com/sleepyfishmusic",
			},
		},
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/5qap5aO4i9A",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is shown studying in Lyon, a city from France where her designer Juan Pablo used to live.",
				"The view through the window depicts the buildings on the slopes of Croix-Rousse, where the bell tower of the Bon-Pasteur church can be spotted.",
			},
			links: map[string]string{
				"website":   "https://lofigirl.com",
				"spotify":   "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				"discord":   "https://discord.com/invite/lofigirl",
				"reddit":    "https://www.reddit.com/r/LofiGirl",
				"instagram": "https://www.instagram.com/lofigirl",
				"facebook":  "https://www.facebook.com/lofigirl",
				"twitter":   "https://twitter.com/lofigirl",
			},
		},
		{
			Slug:  "lofigirl-sleepy",
			Name:  "Lofi Sleepy Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/DWcJFNfaw9c",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is living in Lyon, a city from France where her designer Juan Pablo used to live.",
			},
			links: map[string]string{
				"website":   "https://lofigirl.com",
				"spotify":   "https://open.spotify.com/playlist/0vvXsWCC9xrXsKd4FyS8kM?si=sQXk5Y-GTUeB7OlCRKZ__Q",
				"discord":   "https://discord.com/invite/lofigirl",
				"reddit":    "https://www.reddit.com/r/LofiGirl",
				"instagram": "https://www.instagram.com/lofigirl",
				"facebook":  "https://www.facebook.com/lofigirl",
				"twitter":   "https://twitter.com/lofigirl",
			},
		},
		{
			Slug:  "chillhop",
			Name:  "Chillhop Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/5yx6BWlEVcY",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://chillhop.com",
				"spotify":   "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				"discord":   "https://discord.com/invite/chillhop",
				"reddit":    "https://www.reddit.com/r/chillhop",
				"instagram": "https://www.instagram.com/chillhopmusic",
				"facebook":  "https://www.facebook.com/groups/1561371024098016",
			},
		},
		{
			Slug:  "chillhop-relax",
			Name:  "Chillhop Relaxing Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/7NOSDKb0HlU",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://chillhop.com",
				"spotify":   "https://open.spotify.com/playlist/0CFuMybe6s77w6QQrJjW7d?si=3co_7Q6XT0OJZwkBlqWoDQ",
				"discord":   "https://discord.com/invite/chillhop",
				"reddit":    "https://www.reddit.com/r/chillhop",
				"instagram": "https://www.instagram.com/chillhopmusic",
				"facebook":  "https://www.facebook.com/groups/1561371024098016",
			},
		},
		{
			Slug:  "collegemusic",
			Name:  "College Girl",
			Owner: "College Music",
			URL:   "https://youtu.be/MCkTebktHVc",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://www.collegemusic.co.uk/",
				"spotify":   "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				"instagram": "https://www.instagram.com/collegemusic",
				"facebook":  "https://www.facebook.com/collegemusic",
				"twitter":   "https://twitter.com/collegemusicyt",
			},
		},
		{
			Slug:  "collegemusic-guy",
			Name:  "College Guy",
			Owner: "College Music",
			URL:   "https://youtu.be/XDh0JcxrbPM",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://www.collegemusic.co.uk/",
				"spotify":   "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				"instagram": "https://www.instagram.com/collegemusic",
				"facebook":  "https://www.facebook.com/collegemusic",
				"twitter":   "https://twitter.com/collegemusicyt",
			},
		},
		{
			Slug:  "collegemusic-lonely",
			Name:  "College Lonely",
			Owner: "College Music",
			URL:   "https://youtu.be/bM0Iw7PPoU4",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://www.collegemusic.co.uk/",
				"spotify":   "https://open.spotify.com/playlist/32hJXySZtt9YvnwcYINGZ0",
				"instagram": "https://www.instagram.com/collegemusic",
				"facebook":  "https://www.facebook.com/collegemusic",
				"twitter":   "https://twitter.com/collegemusicyt",
			},
		},
		{
			Slug:  "lofi-code-beats",
			Name:  "Coding Beats",
			Owner: "Joma Tech",
			URL:   "https://youtu.be/PY8f1Z3nARo",
			Cards: []string{},
			links: map[string]string{
				"instagram": "https://www.instagram.com/jomaoppa",
				"facebook":  "https://www.facebook.com/jomaoppa",
				"twitter":   "https://twitter.com/jomaoppa",
			},
		},
		{
			Slug:  "steezyasfuck-coffee-show",
			Name:  "Steezy Coffee Shop",
			Owner: "STEEZYASFUCK",
			URL:   "https://youtu.be/-5KAN9_CzSA",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://www.stzzzy.com",
				"spotify":   "https://open.spotify.com/playlist/2s9R059mmdc8kz6lrUqZZd",
				"instagram": "https://www.instagram.com/stzzyasfvck/",
				"twitter":   "https://twitter.com/stzzyasfvck",
			},
		},
		{
			Slug:  "closedonsunday-pop-culture",
			Name:  "Pop Culture Sunday",
			Owner: "Closed on Sunday",
			URL:   "https://youtu.be/Tgme8K4XlHQ",
			Cards: []string{},
			links: map[string]string{
				"spotify":   "https://open.spotify.com/artist/1LwjR2mIm78OJRTYdkMLl3",
				"discord":   "https://discord.com/invite/closedonsunday",
				"instagram": "https://www.instagram.com/closedonsundayy",
			},
		},
		{
			Slug:  "closedonsunday-starwars",
			Name:  "Star Wars Sunday",
			Owner: "Closed on Sunday",
			URL:   "https://youtu.be/o33l32ZrIy8",
			Cards: []string{},
			links: map[string]string{
				"spotify":   "https://open.spotify.com/artist/1LwjR2mIm78OJRTYdkMLl3",
				"discord":   "https://discord.com/invite/closedonsunday",
				"instagram": "https://www.instagram.com/closedonsundayy",
			},
		},
		{
			Slug:  "tokyolosttracks",
			Name:  "サクラチル",
			Owner: "Tokyo LosT Track -Sakura Chill-",
			URL:   "https://youtu.be/WBfbkPTqUtU",
			Cards: []string{},
			links: map[string]string{
				"spotify":   "https://open.spotify.com/playlist/37i9dQZF1DX87D1EaNZxW1",
				"discord":   "https://discord.com/invite/fHkhEgc",
				"reddit":    "https://www.reddit.com/r/LofiGirl",
				"instagram": "https://www.instagram.com/TokyoLosTTracks",
				"facebook":  "https://www.facebook.com/TokyoLosTTracks",
				"twitter":   "https://twitter.com/TTT_SakuraChill",
			},
		},
		{
			Slug:  "thebootlegboy",
			Name:  "Bootleg Smoke",
			Owner: "the bootleg boy",
			URL:   "https://youtu.be/l7TxwBhtTUY",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://www.thebootlegboy.com",
				"spotify":   "https://open.spotify.com/playlist/71019EDcRamfMmOEEoTdEu?si=XePP-REWQDSuzJT6-SXwSQ",
				"discord":   "https://discord.com/invite/FZrUkey",
				"instagram": "https://www.instagram.com/thebootlegboy",
				"twitter":   "https://twitter.com/thebootlegboy",
			},
		},
		{
			Slug:  "sleepynatula",
			Name:  "Sleepy Natula",
			Owner: "Tone by Gridge",
			URL:   "https://youtu.be/pCuy6U-RVFk",
			Cards: []string{
				"Sleepy Natula is a radio that broadcasts Neo chill beats style songs, produced by Tone by Gridge in the Japanese city of Tokyo.",
				"The animation shows Natula, a girl that fell asleep after drinking a relaxation drink named CHILL OUT.",
				"Natula lives in Sephora, a city located on a musical planet named Tone.",
				"On this planet, music is so important that it touches every species, nature, climate, culture, and civilization of each country.",
				"Tone is a multicultural planet where there are thousands of countries populated by music-loving beings that come from all over the galaxy.",
			},
			links: map[string]string{
				"website":   "https://linktr.ee/tonebygridge?fbclid=IwAR2OYWJlav7MM66780DDtuu2k2viElVF3kgjn-GEK8c2RsvYt0TGOYz31zQ",
				"spotify":   "https://open.spotify.com/playlist/1stD3Nr9W5HWf47ft5TfKh?si=fDM2-nsIRlSdv2IC_KaanQ",
				"discord":   "https://discord.com/invite/eqH6RRj",
				"facebook":  "https://www.facebook.com/tonedotcom",
				"instagram": "https://www.instagram.com/tonebygridge",
				"twitter":   "https://twitter.com/tonedotcom",
			},
		},
		{
			Slug:  "dreamhop",
			Name:  "Dreamhop",
			Owner: "Dreamhop Music",
			URL:   "https://youtu.be/tCs48OFv7xA",
			Cards: []string{},
			links: map[string]string{
				"website":   "https://dreamhopmusic.com",
				"spotify":   "https://open.spotify.com/user/91jfqzlv0htrqrvvc60138qma",
				"discord":   "https://discord.com/invite/FxF9kng",
				"instagram": "https://www.instagram.com/dreamhopp",
				"twitter":   "https://twitter.com/Dreamhopp",
			},
		},
	}

	sort.Slice(videos, func(a, b int) bool {
		return strings.Compare(videos[a].Name, videos[b].Name) < 0
	})

	return videos
}
