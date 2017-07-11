package subscriber

import (
	"encoding/json"
	"io/ioutil"
)

var list []interface{}

type SubscribeInfo struct {
	keyword string
	episode int
}

func GetListSubscriber() []interface{} {
	raw := getSubscribeInfoFile()
	json.Unmarshal(raw, &list)
	return list
}

func getSubscribeInfoFile() []byte {
	raw, _ := ioutil.ReadFile("./subscribe.info.json")
	return raw
}

func GetSubscribeInfo(index int) (string, int) {
	detail := list[index].(map[string]interface{})
	return detail["keyword"].(string), int(detail["episode"].(float64))
}

func UpdateSubscribeEpisode(index int) {
	detail := list[index].(map[string]interface{})
	prevEps := int(detail["episode"].(float64))
	nextEps := prevEps + 1
	detail["episode"] = float64(nextEps)
	newList, _ := json.Marshal(list)
	ioutil.WriteFile("./subscribe.info.json", newList, 0777)
}
