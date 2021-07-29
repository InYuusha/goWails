package main

import (
	_ "embed"
	"stats/pkg/sys"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
stats:=&sys.Stats{}
  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "stats",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(stats)
  app.Run()
}
