package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	crawl("https://nyaa.si/?f=0&c=0_0&q=horriblesubs+rokudenashi")
}

func crawl(url string) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	body := resp.Body
	defer body.Close() // close Body when the function returns

	tokenizer := html.NewTokenizer(body)

	for {
		currToken := tokenizer.Next()
		switch {
		case currToken == html.ErrorToken:
			// End of the document, we're done
			return
		case currToken == html.StartTagToken:
			t := tokenizer.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if isAnchor {
				ok, url := getHref(t)
				if !ok {
					fmt.Println(url)
				}
				fmt.Println(url)
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {

			}
		}
	}
}

func getHref(t html.Token) (bool, string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			return true, a.Val
		}
	}

	return false, ""
}

func getStringFromResponse(response *http.Response) string {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}
