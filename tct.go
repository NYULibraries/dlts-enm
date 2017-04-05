package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ENM_TCT_BASE_URL = "https://nyuapi.infoloom.nyc"

type Topic struct {
	Id int64 `json:"id"`
	DisplayName string `json:"display_name"`
}
func GetResponseBody(url string) (body []byte) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetTopicsAll() (topicsList []Topic) {
	responseBody := GetResponseBody(ENM_TCT_BASE_URL + "/api/hit/basket/all/")

	err := json.Unmarshal(responseBody, &topicsList)
	if err != nil {
		panic(err.Error())
	}

	return
}
