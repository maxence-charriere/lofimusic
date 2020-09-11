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
	"github.com/maxence-charriere/go-app/v7/pkg/logs"
)

const (
	backgroundColor = "#000000"
)

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

func main() {
	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	opts := options{Port: 4000}
	cli.Register().
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&opts)
	cli.Load()

	app.Log("%s", logs.New("starting lofimusic app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr: fmt.Sprintf(":%v", opts.Port),
		Handler: &app.Handler{
			Author:          "Maxence Charriere",
			BackgroundColor: backgroundColor,
			Description:     "Lofi music player",
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
		},
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		app.Log("command failed: %s", err)
		os.Exit(-1)
	}
}
