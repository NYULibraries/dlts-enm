// Copyright © 2018 NYU
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
	"html/template"
	"fmt"
	"os"
	"strings"

	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/db/models"
)

type BrowseTopicsListPageData struct{
	ActiveNavbarTab string
	NavbarTabs []BrowseTopicsListsEntry
	Paths struct{
		WebRoot string
	}
	Topics []BrowseTopicsListPageDataTopic
}

type BrowseTopicsListPageDataTopic struct{
	Name string
	URL string
}

type BrowseTopicsListsEntry struct{
	Label string
	FileBasename string
	Regexp string
}

var BrowseTopicsListsCategories []BrowseTopicsListsEntry

var BrowseTopicsListsDir string

func GenerateBrowseTopicsLists(destination string) {
	BrowseTopicsListsDir = destination + "/browse-topics-lists"

	if _, err := os.Stat(BrowseTopicsListsDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(BrowseTopicsListsDir, os.FileMode(0755))
		if mkdirErr != nil {
			panic(mkdirErr)
		}
	}

	fillBrowseTopicsListsCategories()

	if Source == "database" {
		GenerateDynamicBrowseTopicsListsFromDatabase()
	} else if Source == "cache" {
		// Don't know if will be implementing this
		fmt.Println("Generation of topic browse lists pages has not yet been implemented.")
	} else {
		// Should never get here
	}

	// These browse topics lists are currently unchanging, with all topics
	// hardcoded into the templates.
	WriteStaticBrowseTopicsListsPage("enm-picks")
}

func fillBrowseTopicsListsCategories() {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabet); i++ {
		letter := string(alphabet[i])
		browseTopicsListsEntry := BrowseTopicsListsEntry{
			Label : strings.ToUpper(letter),
			FileBasename : letter,
			Regexp : "^" + letter,
		}
		BrowseTopicsListsCategories = append(BrowseTopicsListsCategories, browseTopicsListsEntry)
	}

	BrowseTopicsListsCategories = append(BrowseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "0-9",
		FileBasename : "0-9",
		Regexp : "^[0-9]",
	})
	BrowseTopicsListsCategories = append(BrowseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "?#@",
		FileBasename : "non-alphanumeric",
		Regexp : "^[^a-z0-9]",
	})

	return
}

func GenerateDynamicBrowseTopicsListsFromDatabase() {
	var topicsWithSortKeys []models.TopicsWithSortKeys

	for _, browseTopicsListsCategory := range BrowseTopicsListsCategories {
		topicsWithSortKeys = db.GetTopicsWithSortKeysByFirstSortableCharacterRegexp(browseTopicsListsCategory.Regexp)
		filename := browseTopicsListsCategory.FileBasename + ".html"
		browseTopicsListPageData := CreateBrowseTopicsListPageData(topicsWithSortKeys, browseTopicsListsCategory.Label)
		err := WriteDynamicBrowseTopicsListPage(filename, browseTopicsListPageData)
		if (err != nil) {
			panic(err)
		}
	}

}

func CreateBrowseTopicsListPageData(topicsWithSortKeys []models.TopicsWithSortKeys, browseTopicsListsCategoryLabel string) (browseTopicsListPageData BrowseTopicsListPageData) {
	for _, topicWithSortKeys := range topicsWithSortKeys {

		topic := BrowseTopicsListPageDataTopic{
			Name: topicWithSortKeys.DisplayNameDoNotUse,
			URL: "../topic-pages/" + GetRelativeFilepathForTopicPage(topicWithSortKeys.TctID),
		}
		browseTopicsListPageData.Topics = append(browseTopicsListPageData.Topics, topic)
	}

	browseTopicsListPageData.ActiveNavbarTab = browseTopicsListsCategoryLabel
	browseTopicsListPageData.NavbarTabs = BrowseTopicsListsCategories
	browseTopicsListPageData.Paths = Paths{ WebRoot: ".." }

	return
}

func WriteStaticBrowseTopicsListsPage(listname string) (err error){
	var filename = listname + ".html"
	var templateFile = listname + ".html"

	tpl := template.New(filename)
	tpl, err = tpl.ParseFiles(
		BrowseTopicListsTemplateDirectory + "/" + templateFile,
		SharedTemplateDirectory    + "/banner.html",
	)

	if err != nil {
		return err
	}

	file := BrowseTopicsListsDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	pageData := struct{
		Paths
	}{
		Paths: Paths{
			WebRoot: "..",
		},
	}

	err = tpl.Execute(f, pageData)
	if err != nil {
		return err
	}

	return nil
}

func WriteDynamicBrowseTopicsListPage(filename string, browseTopicsListPageData BrowseTopicsListPageData) (err error){
	tpl := template.New("a-to-z.html")
	tpl, err = tpl.ParseFiles(
		BrowseTopicListsTemplateDirectory + "/a-to-z.html",
		BrowseTopicListsTemplateDirectory + "/a-to-z-content.html",
		SharedTemplateDirectory           + "/banner.html",
	)

	if err != nil {
		return err
	}

	file := BrowseTopicsListsDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	err = tpl.Execute(f, browseTopicsListPageData)
	if err != nil {
		return err
	}

	return nil
}
