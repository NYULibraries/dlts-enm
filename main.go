package main

import (
	"fmt"
)

func main() {
	topicsList := GetTopicsAll()

	for i, v := range topicsList {
		fmt.Printf("Topic #%d: %v\n",i , v)
	}
}