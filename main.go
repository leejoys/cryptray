package main

import (
	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed aw.ico
var ico []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Sup")
	systray.SetIcon(ico)
	q := systray.AddMenuItem("Quit", "")
	e := systray.AddMenuItem("Exit", "")
	systray.SetTooltip("Sup")
	for {
		select {
		case <-q.ClickedCh:
			systray.Quit()
		case <-e.ClickedCh:
			systray.Quit()
		}
	}
}

func onExit() {

}
