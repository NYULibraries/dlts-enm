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

// Created using github.com/ChimeraCoder/gojson/gojson
// `cat json-api-samples/topic-detail.json | gojson -name=TopicDetail`
type TopicDetail struct {
	Basket struct {
		DisplayName string `json:"display_name"`
		ID          int64  `json:"id"`
		Occurs      []struct {
			ID       int64 `json:"id"`
			Location struct {
				Document struct {
					Author string `json:"author"`
					Title  string `json:"title"`
				} `json:"document"`
				ID      int64  `json:"id"`
				Localid string `json:"localid"`
			} `json:"location"`
		} `json:"occurs"`
		TopicHits []struct {
			Bypass    bool   `json:"bypass"`
			Hidden    bool   `json:"hidden"`
			ID        int64  `json:"id"`
			Name      string `json:"name"`
			Preferred bool   `json:"preferred"`
			Scope     struct {
				ID    int64  `json:"id"`
				Scope string `json:"scope"`
			} `json:"scope"`
		} `json:"topic_hits"`
	} `json:"basket"`
	Relations []struct {
		Basket struct {
			DisplayName string `json:"display_name"`
			ID          int64  `json:"id"`
		} `json:"basket"`
		Direction    string `json:"direction"`
		ID           int64  `json:"id"`
		Relationtype struct {
			ID          int64  `json:"id"`
			RoleFrom    string `json:"role_from"`
			RoleTo      string `json:"role_to"`
			Rtype       string `json:"rtype"`
			Symmetrical bool   `json:"symmetrical"`
		} `json:"relationtype"`
	} `json:"relations"`
}
