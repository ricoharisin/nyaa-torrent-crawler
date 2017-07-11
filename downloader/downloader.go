package downloader

import (
	"fmt"
	"net/http"
	"strings"

	"io"
	"os"
)

const BaseURL string = "https://nyaa.si"

func DownloadTorrent(torrentUrl string) bool {
	if _, err := os.Stat("./torrent"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("torrent", 0777)
		}
	}

	location := "./torrent/" + strings.Replace(torrentUrl, "/download/", "", -1)

	output, err := os.Create(location)
	if err != nil {
		fmt.Println("Error while creating", location, "-", err)
		return false
	}

	defer output.Close()

	url := BaseURL + torrentUrl
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR: \"" + url + "\"")
		return false
	}

	body := resp.Body
	defer body.Close()

	_, errCopy := io.Copy(output, body)

	if errCopy != nil {
		fmt.Println("Error while copying to", location, "-", err)
		return false
	}

	return true

}
