package subscription

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

func GetListSubscription() []SubscribeInfo {
	raw := getSubscribeInfoFile()
	json.Unmarshal(raw, &list)
	return list
}

func getSubscribeInfoFile() []byte {
	raw, _ := ioutil.ReadFile("./subscribe.info.json")
	return raw
}

func GetSubscriptionInfo(index int) (string, int) {
	return list[index].Keyword, list[index].Episode
}

func UpdateSubscriptionEpisode(index int) {
	prevEps := list[index].Episode
	nextEps := prevEps + 1
	list[index].Episode = nextEps
	newList, _ := json.Marshal(list)
	ioutil.WriteFile("./subscribe.info.json", newList, 0777)
}

func InsertNewSubscription(keyword string, episode int) {
	oldList := GetListSubscription()
	var newSubscribe SubscribeInfo
	newSubscribe.Keyword = keyword
	newSubscribe.Episode = episode
	oldList = append(oldList, newSubscribe)
	newList, _ := json.Marshal(oldList)
	ioutil.WriteFile("./subscribe.info.json", newList, 0777)
}

func RemoveSubscription(index int) {
	oldList := GetListSubscription()
	if index < len(oldList) {
		oldList = append(oldList[:index], oldList[index+1:]...)
		newList, _ := json.Marshal(oldList)
		ioutil.WriteFile("./subscribe.info.json", newList, 0777)
	} else {
		fmt.Println("invalid index")
	}

}
