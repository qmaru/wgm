package main

import (
	"log"

	"github.com/egoist/mygo"
	myui "github.com/egoist/mygo/ui"

	"github.com/qmaru/wgm/backend"
	wgmui "github.com/qmaru/wgm/internal/ui"
)

func main() {
	mygo.App.WhenReady(func() {
		a := wgmui.NewApp(backend.DBStore{})
		mygo.NewWindow(mygo.WindowOptions{
			Title:     "wgm",
			Width:     1280,
			Height:    720,
			MinWidth:  860,
			MinHeight: 520,
			Vibrancy:  mygo.VibrancyMica,
			Content:   myui.View(a.View),
		})
	})

	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
