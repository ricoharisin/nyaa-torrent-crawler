package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"os"

	"io"

	"golang.org/x/net/html"
)

const BaseURL string = "https://nyaa.si"

func main() {
	//crawl("https://nyaa.si/?f=0&c=0_0&q=horriblesubs+rokudenashi")
	//eps, islast := isLatestEpisode("[HorribleSubs] Rokudenashi Majutsu Koushi to Akashic Records - 12 [720p].mkv", 11)
	//fmt.Println(eps)
	//fmt.Println(islast)
	startCrawling("horriblesubs rokudenashi 720", 10)
	//downloadTorrent("asdsa")

}

func startCrawling(keyword string, prevEpisode int) {
	isSuccess, urlDetail, lastEpisode := getDetailUrlForLastEpisode(keyword, prevEpisode)
	fmt.Println(urlDetail)
	fmt.Println(lastEpisode)
	if isSuccess {
		_, torrentUrl := getTorrentUrlFromDetail(urlDetail)
		fmt.Println(torrentUrl)
		downloadTorrent(torrentUrl)
	}
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
	defer body.Close() // close Body when the function returns

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

func isLatestEpisode(title string, prevEpisode int) (bool, int) {
	targetEpisode := prevEpisode + 1

	for i, r := range title {
		c := string(r)
		if c == "-" {
			strEpisode := string(title[i+2]) + string(title[i+3])
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

func downloadTorrent(torrentUrl string) {
	if _, err := os.Stat("./torrent"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("torrent", 0777)
		}
	}

	location := "./torrent/" + strings.Replace(torrentUrl, "/download/", "", -1)

	output, err := os.Create(location)
	if err != nil {
		fmt.Println("Error while creating", location, "-", err)
		return
	}

	defer output.Close()

	/*if _, err := os.Stat(location); err != nil {
		if os.IsNotExist(err) {
			output, _ := os.Create(location)
		} else {
			output, _ := os.Open(location)
		}
	}*/

	url := BaseURL + torrentUrl
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: \"" + url + "\"")
		return
	}

	body := resp.Body
	defer body.Close()

	io.Copy(output, body)
}
