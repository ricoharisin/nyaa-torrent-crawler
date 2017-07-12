package crawler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const BaseURL string = "https://nyaa.si"

func StartCrawling(keyword string, prevEpisode int) (bool, string) {
	fmt.Println("crawling " + keyword + " .........")
	isSuccess, urlDetail, _ := getDetailUrlForLastEpisode(keyword, prevEpisode)
	if isSuccess {
		_, torrentUrl := getTorrentUrlFromDetail(urlDetail)
		fmt.Println("found the latest episode!! will download" + torrentUrl)
		return true, torrentUrl
	}
	fmt.Println("no latest episode found :(")
	return false, ""
}

func getDetailUrlForLastEpisode(keyword string, prevEpisode int) (bool, string, int) {
	query := strings.Replace(keyword, " ", "+", -1)
	url := BaseURL + "/?f=0&c=0_0&q=" + query
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return false, "", 0
	}

	body := resp.Body
	defer body.Close()

	tokenizer := html.NewTokenizer(body)

	for {
		currToken := tokenizer.Next()
		switch {
		case currToken == html.ErrorToken:
			return false, "", 0
		case currToken == html.StartTagToken:
			t := tokenizer.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				title := getTitle(t)
				if title != "" {
					isLatestEpisode, lastEpisode := isLatestEpisode(title, prevEpisode)
					if isLatestEpisode {
						href := getHref(t)
						if href != "" {
							return true, href, lastEpisode
						}
					}
				}
			}
		}
	}
}

func getTorrentUrlFromDetail(urlDetail string) (bool, string) {
	url := BaseURL + urlDetail
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return false, ""
	}

	body := resp.Body
	defer body.Close() // close Body when the function returns

	tokenizer := html.NewTokenizer(body)

	for {
		currToken := tokenizer.Next()
		switch {
		case currToken == html.ErrorToken:
			return false, ""
		case currToken == html.StartTagToken:
			t := tokenizer.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				href := getHref(t)
				if strings.Contains(href, "torrent") {
					return true, href
				}
			}
		}
	}
}

func isLatestEpisode(title string, prevEpisode int) (bool, int) {
	targetEpisode := prevEpisode + 1

	for i, r := range title {
		c := string(r)
		if c == "-" {
			strEpisode := string(title[i+2]) + string(title[i+3])
			//fmt.Println(strEpisode)
			intEpisode, err := strconv.Atoi(strEpisode)
			if err == nil {
				if intEpisode == targetEpisode {
					return true, intEpisode
				}
			}
		}
	}

	return false, 0
}
