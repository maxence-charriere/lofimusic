// +build wasm

package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

func main() {
	for _, s := range channels.Slugs() {
		app.Route("/"+s, &shell{ChannelSlug: s})
	}

	app.Route("/", &shell{})
	app.Run()
}
