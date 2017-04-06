package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ENM_TCT_BASE_URL = "https://nyuapi.infoloom.nyc"

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

func GetEpubDetail(epubId int) (epubDetail EpubDetail) {
	epubIdString := strconv.Itoa(epubId)
	responseBody := GetResponseBody(ENM_TCT_BASE_URL + "/api/epub/document/" + epubIdString)

	err := json.Unmarshal(responseBody, &epubDetail)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetEpubsAll() (epubsList []Epub) {
	responseBody := GetResponseBody(ENM_TCT_BASE_URL + "/api/epub/document/all/")

	err := json.Unmarshal(responseBody, &epubsList)
	if err != nil {
		panic(err.Error())
	}

	return
}

func GetNamesAll() (namesList []Name) {
	responseBody := GetResponseBody(ENM_TCT_BASE_URL + "/api/hit/hits/all/")

	err := json.Unmarshal(responseBody, &namesList)
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

func GetTopicDetail(topicId int) (topicDetails TopicDetail) {
	topicIdString := strconv.Itoa(topicId)
	responseBody := GetResponseBody(ENM_TCT_BASE_URL + "/api/hit/basket/" + topicIdString)

	err := json.Unmarshal(responseBody, &topicDetails)
	if err != nil {
		panic(err.Error())
	}

	return
}

// Created with the help of github.com/ChimeraCoder/gojson/gojson
// `cat json-api-samples/epubs-all.txt | gojson -name=EpubsAll`
type Epub struct {
	Author    string `json:"author"`
	ID        int64  `json:"id"`
	Isbn      string `json:"isbn"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
}

// Created by github.com/ChimeraCoder/gojson/gojson
// `cat json-api-samples/epub-detail.json | gojson -name=EpubDetail`
type EpubDetail struct {
	Author    string `json:"author"`
	Isbn      string `json:"isbn"`
	Locations []struct {
		Content        string `json:"content"`
		ID             int64  `json:"id"`
		Localid        string `json:"localid"`
		SequenceNumber int64  `json:"sequence_number"`
	} `json:"locations"`
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
}

// Created with the help of github.com/ChimeraCoder/gojson/gojson
// `cat json-api-samples/names-all.json | gojson -name=NamesAll`
type Name struct {
	Basket    int64  `json:"basket"`
	Hidden    bool   `json:"hidden"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Preferred bool   `json:"preferred"`
	Scope     string `json:"scope"`
}

type Topic struct {
	Id int64 `json:"id"`
	DisplayName string `json:"display_name"`
}

// Created by github.com/ChimeraCoder/gojson/gojson
// `cat json-api-samples/topic-detail.json | gojson -name=TopicDetail`
type TopicDetail struct {
	Basket struct {
		DisplayName string `json:"display_name"`
		ID          int64  `json:"id"`
		Occurs      []struct {
			Basket   int64 `json:"basket"`
			ID       int64 `json:"id"`
			Location struct {
				Document struct {
					Author string `json:"author"`
					Title  string `json:"title"`
				} `json:"document"`
				ID             int64  `json:"id"`
				Localid        string `json:"localid"`
				SequenceNumber int64  `json:"sequence_number"`
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
