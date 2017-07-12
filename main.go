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
		case "list":
			list()
			break
		case "removesubscribe":
			list()
			removesubscribe()
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
	fmt.Println("nyaa-torrent-crawler <option>")
	fmt.Println("\noption list:")
	fmt.Println("subscribe <keyword> <current episode> | subscribe anime based on nyaa.si search keyword")
	fmt.Println("crawl | start crawling based on susbcribe list")
	fmt.Println("list | show subscribe list")
	fmt.Println("removesubscribe | remove subscribe from list")
}

func crawl() {
	listSubscribe := subscriber.GetListSubscriber()
	for i := range listSubscribe {
		keyword, eps := subscriber.GetSubscribeInfo(i)
		isSuccess, torrentUrl := crawler.StartCrawling(keyword, eps)
		if isSuccess {
			isSuccessDownload := downloader.DownloadTorrent(torrentUrl)
			if isSuccessDownload {
				subscriber.UpdateSubscribeEpisode(i)
			}
		}
	}
}

func list() {
	fmt.Printf("|%-2s|%s|%-2s|\n", "Index", "Keyword", "Current Episode")
	listSubscribe := subscriber.GetListSubscriber()
	for i := range listSubscribe {
		keyword, eps := subscriber.GetSubscribeInfo(i)
		fmt.Printf("|%-2d|%s|%-2d|\n", i, keyword, eps)
	}
}

func removesubscribe() {
	var index int
	fmt.Printf("which index? ")
	fmt.Scanln(&index)
	subscriber.RemoveSubscribe(index)
}
