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

}