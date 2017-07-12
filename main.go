package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ricoharisin/nyaa-torrent-crawler/crawler"
	"github.com/ricoharisin/nyaa-torrent-crawler/downloader"
	"github.com/ricoharisin/nyaa-torrent-crawler/subscriber"
)

func main() {
	args := os.Args
	fmt.Println(len(args))
	if len(args) < 2 {
		showHelp()
	} else {
		switch args[1] {
		case "subscribe":
			subscribe(args[2], args[3])
			break
		case "crawl":
			crawl()
			break
		default:
			showHelp()
			break
		}
	}
}

func subscribe(args1 string, args2 string) {
	intargs, _ := strconv.Atoi(args2)
	subscriber.InsertNewSubscribe(args1, intargs)
}

func showHelp() {
	fmt.Println("usage: ")
	fmt.Println("nyaa-torrent-crawler subscribe \"keyword\" start_episode_num")
	fmt.Println("nyaa-torrent-crawler crawl")
}

func crawl() {
	listSubscribe := subscriber.GetListSubscriber()
	for i := range listSubscribe {
		keyword, eps := subscriber.GetSubscribeInfo(i)
		fmt.Println(eps)
		isSuccess, torrentUrl := crawler.StartCrawling(keyword, eps)
		if isSuccess {
			isSuccessDownload := downloader.DownloadTorrent(torrentUrl)
			if isSuccessDownload {
				subscriber.UpdateSubscribeEpisode(i)
			}
		}
	}
}
