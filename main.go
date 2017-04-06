package main

import (
	"fmt"
)

func main() {
	topicsList := GetTopicsAll()
	for i, v := range topicsList {
		fmt.Printf("Topic #%d: %v\n", i, v)
	}

	topicDetail := GetTopicDetail(7399)
	fmt.Println(topicDetail)

	namesList := GetNamesAll()
	for i, v := range namesList {
		fmt.Printf("Name #%d: %v\n", i, v)
	}

	epubsList := GetEpubsAll()
	for i, v := range epubsList {
		fmt.Printf("epub #%d: %v\n", i, v)
	}

	epubDetail := GetEpubDetail(12)
	fmt.Println(epubDetail)

	location := GetLocation(2410)
	fmt.Println(location)
}