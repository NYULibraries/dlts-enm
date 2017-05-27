package dataviz

import(
	"encoding/json"
	"fmt"

	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/db/models"
)

type Node struct {
	TopicId int `json:"id"`
	Children []*Node `json:"children,omitempty"`
	Name string `json:"name"`
}

func (node *Node) AddChildNodes() {
	subentryTopics := db.GetTopicSubEntries(node.TopicId)
	for _, subentryTopic := range subentryTopics {
		node.Children = append(node.Children, &Node{
			TopicId: subentryTopic.TctID,
			Children: nil,
			Name: subentryTopic.DisplayNameDoNotUse,
		})
	}

	for _, child := range node.Children {
		child.AddChildNodes()
	}
}

func Trees1(topicId int) {
	topic, err := models.TopicByTctID(db.DB, topicId)
	if err != nil {
		panic(err)
	}

	rootNode := &Node{
		TopicId: topicId,
		Children: nil,
		Name: topic.DisplayNameDoNotUse,
	}
	rootNode.AddChildNodes()

	bytes, err := json.MarshalIndent(rootNode, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}