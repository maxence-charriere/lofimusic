package pwa

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func focus(v bool) string {
	if !v {
		return ""
	}
	return "accent"
}

func pathFocus(v ...string) string {
	path := "/" + strings.Join(v, "/")
	currentPath := app.Window().URL().Path
	return focus(currentPath == path)
}

func navToFragment(ctx app.Context) {
	fragment := ctx.Page().URL().Fragment
	if fragment == "" {
		fragment = "page-top"
	}
	ctx.ScrollTo(fragment)
}

func hideElement(v bool) string {
	if v {
		return "hide"
	}
	return ""
}

func disableElement(v bool) string {
	if v {
		return "disable"
	}
	return ""
}
