// +build !wasm

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/maxence-charriere/go-app/v7/pkg/cli"
	"github.com/maxence-charriere/go-app/v7/pkg/errors"
	"github.com/maxence-charriere/go-app/v7/pkg/logs"
)

const (
	backgroundColor = "#000000"
)

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	h := app.Handler{
		Author:          "Maxence Charriere",
		BackgroundColor: backgroundColor,
		Description:     "Lofi music player",
		Icon: app.Icon{
			Default: "/web/logo.png",
		},
		Keywords: []string{
			"lofi",
			"music",
			"chilled cow",
		},
		LoadingLabel: "Lofi music player to work, study or relax.",
		Name:         "Lofimusic",
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"https://fonts.googleapis.com/css2?family=Roboto&display=swap",
			"/web/lofimusic.css",
		},
		ThemeColor: backgroundColor,
		Title:      "Lofimusic.app",
	}

	opts := options{Port: 4000}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&opts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run Lofimusic app on GitHub Pages.`).
		Options(&githubOpts)
	cli.Load()

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, opts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}

}

func runLocal(ctx context.Context, h http.Handler, opts options) {
	app.Log("%s", logs.New("starting lofimusic app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	if err := app.GenerateStaticWebsite(opts.Output, h, channels.Slugs()...); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Log("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
