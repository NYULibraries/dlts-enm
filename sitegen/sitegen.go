// Copyright © 2017 NYU
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sitegen

import (
	"fmt"
	"html/template"
	"os"

	"github.com/nyulibraries/dlts-enm/db"
)

// Tricky...this assumes that location of the templates relative to working directory
// matches what the location would be if calling `go run main.go` from root of this
// repo.  This is brittle: user could call `go run ../main.go` from sitegen/, or
// could be calling the built binary from anywhere relative to templates.
// Consider including the templates in the binary using something like
// https://github.com/jteeuwen/go-bindata
const TemplateDirectory =  "sitegen/templates/topic-page"

type ExternalRelation struct{
	Relationship string
	URL string
	Vocabulary string
}

type Paths struct{
	WebRoot string
}

type RelatedTopic struct{
	ID int
	Name string
	NumberOfOccurrences int
}

type TopicPageData struct{
	AlternateNames []string
	DisplayName string
	EPUBMatches []string
	ExternalRelations []ExternalRelation
	Paths Paths
	RelatedTopics []RelatedTopic
	VisualizationData string
}

var Destination string

func GenerateTopicPages(destination string) {
	Destination = destination
	topicPagesDir := Destination + "/topic-pages"
	if _, err := os.Stat(topicPagesDir); os.IsNotExist(err) {
		os.Mkdir(topicPagesDir, os.FileMode(0755))
		if err != nil {
			panic(err)
		}
	}

	topicsWithAlternateNames := db.GetTopicsWithAlternateNamesAll()
	var alternateNames []string
	inProgressTopicID := topicsWithAlternateNames[0].TctID
	inProgressTopicName := topicsWithAlternateNames[0].DisplayNameDoNotUse
	for _, topicWithAlternateNames := range topicsWithAlternateNames{
		// Finished getting alternate names for in-process topic.
		if (topicWithAlternateNames.TctID != inProgressTopicID) {
			// Generate the topic page.
			err := GenerateTopicPage(inProgressTopicID, inProgressTopicName, alternateNames)
			if err != nil {
				panic(fmt.Sprintf("ERROR: couldn't generate topic page for %d: \"%s\"\n", inProgressTopicID, err.Error()))
			}

			// Clear alternate names and set new in progress topic.
			alternateNames = []string{}
			inProgressTopicID = topicWithAlternateNames.TctID
			inProgressTopicName = topicWithAlternateNames.DisplayNameDoNotUse
		}

		if (topicWithAlternateNames.Name != inProgressTopicName) {
				alternateNames = append(alternateNames, topicWithAlternateNames.Name)
		}
	}
}

func GenerateTopicPage(topicID int, topicDisplayName string, alternateNames []string) error {
	relatedTopics := []RelatedTopic{}
	relatedTopicNamesWithNumberOfOccurrences := db.GetRelatedTopicNamesForTopicWithNumberOfOccurrences(topicID)
	for _, relatedTopic := range relatedTopicNamesWithNumberOfOccurrences {
		relatedTopics = append(relatedTopics, RelatedTopic{
			ID: relatedTopic.Topic2ID,
			Name: relatedTopic.DisplayNameDoNotUse,
			NumberOfOccurrences: int(relatedTopic.NumberOfOccurrences),
		})
	}
	fmt.Println("%s (%d): %v", topicDisplayName, topicID, alternateNames)
	fmt.Println("%v", relatedTopics)
	return nil
}

func Test() {
	funcs := template.FuncMap{
		"lastIndex": func (s []ExternalRelation) int {
			return len(s) - 1;
		},
	}

	tpl := template.New("index.html").Funcs(funcs)
	tpl, err := tpl.ParseFiles(
		TemplateDirectory + "/index.html",
		TemplateDirectory + "/epub.html",
		TemplateDirectory + "/banner.html",
	)

	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, TopicPageData{
		AlternateNames: []string{"alt1", "alt2", "alt3"},
		DisplayName: "topic!",
		EPUBMatches: []string{"epub1", "epub2", "epub3"},
		ExternalRelations: []ExternalRelation{
			{
				Relationship: "exactMatch",
				URL: "http://nowhere.net",
				Vocabulary: "VIAF",
			},
			{
				Relationship: "exactMatch",
				URL: "http://nowhere2.net",
				Vocabulary: "VIAF2",
			},
			{
				Relationship: "exactMatch",
				URL: "http://nowhere2.net",
				Vocabulary: "VIAF2",
			},
		},
		Paths: Paths{
			WebRoot: "/webroot",
		},
		VisualizationData: "[VISUALIZATION DATA]",
	})
	if err != nil {
		panic(err)
	}
}