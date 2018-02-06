package main


/// go get -u github.com/gopherjs/gopherjs

/// go get github.com/fabioberger/chrome
/// go get honnef.co/go/js/dom

import (
	"strconv"
	"honnef.co/go/js/dom"
	"github.com/fabioberger/chrome"
)

func main() {
	c := chrome.NewChrome()

	tabDetails := chrome.Object{
		"active": false,

	}
	c.Tabs.Create(tabDetails, func(tab chrome.Tab) {
		notification := "Tab with id: " + strconv.Itoa(tab.Id) + " created!"
		dom.GetWindow().Document().GetElementByID("notification").SetInnerHTML(notification)
		tab.Url = "www.google.dk"
	})







}