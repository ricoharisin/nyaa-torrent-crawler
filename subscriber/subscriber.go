package subscriber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var list []SubscribeInfo

type SubscribeInfo struct {
	Keyword string `json:"keyword"`
	Episode int    `json:"episode"`
}

func GetListSubscriber() []SubscribeInfo {
	raw := getSubscribeInfoFile()
	json.Unmarshal(raw, &list)
	fmt.Println(list)
	return list
}

func getSubscribeInfoFile() []byte {
	raw, _ := ioutil.ReadFile("./subscribe.info.json")
	return raw
}

func GetSubscribeInfo(index int) (string, int) {
	return list[index].Keyword, list[index].Episode
}

func UpdateSubscribeEpisode(index int) {
	prevEps := list[index].Episode
	nextEps := prevEps + 1
	list[index].Episode = nextEps
	newList, _ := json.Marshal(list)
	ioutil.WriteFile("./subscribe.info.json", newList, 0777)
}
