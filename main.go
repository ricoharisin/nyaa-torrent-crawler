package main

import (
	"fmt"

	"github.com/ricoharisin/nyaa-torrent-crawler/crawler"
	"github.com/ricoharisin/nyaa-torrent-crawler/downloader"
	"github.com/ricoharisin/nyaa-torrent-crawler/subscriber"
)

func main() {
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
