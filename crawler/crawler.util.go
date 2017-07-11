package crawler

import (
	"golang.org/x/net/html"
)

func getHref(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}

	return ""
}

func getTitle(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "title" {
			return a.Val
		}
	}

	return ""
}
