package pwa

import (
	"math/rand"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/lofimusic/pkg/radio"
)

const (
	buyMeACoffeeURL     = "https://www.buymeacoffee.com/maxence"
	githubURL           = "https://github.com/maxence-charriere/lofimusic"
	twitterURL          = "https://twitter.com/jonhymaxoo"
	coinbaseBusinessURL = "https://commerce.coinbase.com/checkout/851320a4-35b5-41f1-897b-74dd5ee207ae"
)

// The progressive web app info.
var Info = struct {
	Name               string
	Author             string
	ThemeColor         string
	DefaultTitle       string
	DefaultDescription string
	LogoURL            string
	ImageURL           string
	Keywords           []string
}{
	Name:               "Lofimusic",
	ThemeColor:         "black",
	DefaultTitle:       "Music to work, study and relax.",
	Author:             "Maxence Charriere",
	DefaultDescription: "LoFi music player to work, study and relax.",
	LogoURL:            "/web/logo.png",
	ImageURL:           "https://lofimusic.app/web/covers/lofimusic.png",
	Keywords: []string{
		"lofi",
		"lo-fi",
		"music",
		"lofimusic",
		"chill",
		"chilled",
		"beats",
		"relax",
		"study",
		"sleep",
		"hiphop",
		"app",
		"pwa",
	},
}

// InitAndRunOnBrowser initializes the PWA and run it when launched on a web
// browser.
func InitAndRunOnBrowser() {
	rand.Seed(time.Now().UnixNano())
	analytics.Add(analytics.NewGoogleAnalytics())

	app.Handle(loadVideo, handleLoadVideo)

	for _, v := range radio.List() {
		app.Route("/"+v.Slug, newPage())
	}
	app.Route("/", newPage())
	app.RunWhenOnBrowser()
}
