package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
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
	atomPrice, err := getPrice("https://coinmarketcap.com/currencies/cosmos/")
	if err != nil {
		return
	}
	btcPrice, err := getPrice("https://coinmarketcap.com/currencies/bitcoin/")
	if err != nil {
		return
	}
	etherPrice, err := getPrice("https://coinmarketcap.com/currencies/ethereum/")
	if err != nil {
		return
	}
	// Create UI with basic HTML passed via data URI
	page := fmt.Sprintf(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Cosmos price: %s</h1></body>
		<body><h1>Bitcoin price: %s</h1></body>
		<body><h1>Ethereum price: %s</h1></body>
	</html>
	`, atomPrice, btcPrice, etherPrice)
	ui, err := lorca.New("data:text/html,"+url.PathEscape(page), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

func onExit() {

}

func getPrice(path string) (string, error) {
	res, err := http.Get(path)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	price := doc.Find(".priceValue")
	return price.Text(), nil
}
