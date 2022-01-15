package main

import (
	_ "embed"
	"log"

	"github.com/getlantern/systray"
	"github.com/zserge/lorca"
)

//go:embed aw.ico
var ico []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Sup")
	systray.SetIcon(ico)
	v := systray.AddMenuItem("View", "")
	systray.AddSeparator()
	q := systray.AddMenuItem("Quit", "")
	systray.SetTooltip("Sup")

	for {
		select {
		case <-q.ClickedCh:
			systray.Quit()
		case <-v.ClickedCh:
			view()
		}
	}
}

func view() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("https://leejoys.github.io/", "", 640, 480)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

func onExit() {

}
