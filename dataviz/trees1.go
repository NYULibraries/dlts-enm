package dataviz

import(
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"

	"github.com/nyulibraries/dlts-enm/db"
)

const VisualizationBoilerplateDir = "dataviz/visualization-boilerplate"

var OutputDir string

var HtmlFileContents []byte
var CssFileContents []byte
var D3FileContents []byte
var D3LayoutFileContents []byte

type Node struct {
	TopicId int `json:"id"`
	Children []*Node `json:"children,omitempty"`
	Name string `json:"name"`
}

func (node *Node) AddChildNodes(ancestors []int) (depth int, err error){
	depth = len(ancestors)

	//fmt.Printf("topic %d (%v)\n", node.TopicId, ancestors)

	subentryTopics := db.GetTopicSubEntries(node.TopicId)
	for _, subentryTopic := range subentryTopics {

		//fmt.Printf("subentry %d\n", subentryTopic.TctID)

		for _, element := range ancestors {
			if element == subentryTopic.TctID {
				return depth, errors.New(fmt.Sprintf("loop: %v -> %d", ancestors, subentryTopic.TctID))
			}
		}

		node.Children = append(node.Children, &Node{
			TopicId: subentryTopic.TctID,
			Children: nil,
			Name: subentryTopic.DisplayNameDoNotUse,
		})
	}

	for _, child := range node.Children {
		newDepth, err := child.AddChildNodes(append(ancestors, child.TopicId))
		if newDepth > depth {
			depth = newDepth
		}

		if err != nil {
			return depth, err
		}
	}

	return depth, nil
}

func Trees1() {
	var err error
	HtmlFileContents, err = ioutil.ReadFile(VisualizationBoilerplateDir + "/index.html")
	if err != nil {
		panic(err)
	}
	CssFileContents, err = ioutil.ReadFile(VisualizationBoilerplateDir + "/css/style.css")
	if err != nil {
		panic(err)
	}
	D3FileContents, err = ioutil.ReadFile(VisualizationBoilerplateDir + "/js/d3.js")
	if err != nil {
		panic(err)
	}
	D3LayoutFileContents, err = ioutil.ReadFile(VisualizationBoilerplateDir + "/js/d3.layout.js")
	if err != nil {
		panic(err)
	}

	trees1Dir := OutputDir + "/trees1"
	os.Mkdir(trees1Dir, os.FileMode(0755))

	topics := db.GetTopicsAll()

	var treeAnchorTags []string
	for _, topic := range topics {
		rootNode := &Node{
			TopicId: topic.TctID,
			Children: nil,
			Name: topic.DisplayNameDoNotUse,
		}

		// Need to keep track of ancestor node IDs in order to detect infinite loops
		var ancestors = []int{topic.TctID}
		depth, err := rootNode.AddChildNodes(ancestors)
		// Loop detected, skip this topic
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if rootNode.Children == nil {
			continue
		}

		bytes, err := json.MarshalIndent(rootNode, "", "\t")
		if err != nil {
			panic(err)
		}

		CreateTrees1Visualization(trees1Dir + "/" + strconv.Itoa(topic.TctID), bytes)

		treeAnchorTags = append(treeAnchorTags,
			fmt.Sprintf("<p><a id=\"%d\" name=\"%d\" href=\"%d/\">%s (%d)</a></p>\n", topic.TctID, topic.TctID, topic.TctID, topic.DisplayNameDoNotUse, depth))
	}

	indexHtml := `<!DOCTYPE html><html lang="en">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>NYUP-234: Collapsible tree visualizations for topics with sub-entries</title>
</head>
<body>
`
	indexHtml += strings.Join(treeAnchorTags, "\n")

	indexHtml += `</body>
</html>
`
	err = ioutil.WriteFile(trees1Dir + "/index.html", []byte(indexHtml), 0644)
}

func CreateTrees1Visualization(dir string, json []byte) {
	fmt.Println(dir)

	os.Mkdir(dir, os.FileMode(0755))
	os.Mkdir(dir + "/css", os.FileMode(0755))
	os.Mkdir(dir + "/js", os.FileMode(0755))

	var err error
	err = ioutil.WriteFile(dir + "/topic-tree-data.json", json, 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir + "/index.html", HtmlFileContents, 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir + "/css/style.css", CssFileContents, 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir + "/js/d3.js", D3FileContents, 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir + "/js/d3.layout.js", D3LayoutFileContents, 0644)
	if err != nil {
		panic(err)
	}
}