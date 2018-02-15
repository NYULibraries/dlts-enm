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

// Tricky...this assumes that location of the templates relative to working directory
// matches what the location would be if calling `go run main.go` from root of this
// repo.  This is brittle: user could call `go run ../main.go` from sitegen/, or
// could be calling the built binary from anywhere relative to templates.
// Consider including the templates in the binary using something like
// https://github.com/jteeuwen/go-bindata
const BrowseTopicListsTemplateDirectory = "sitegen/templates/browse-topics-lists"

type BrowseTopicsListPageData struct{

}

type BrowseTopicsListsEntry struct{
	Label string
	FileBasename string
	Regexp string
}

var BrowseTopicsListsDir string

func GenerateBrowseTopicsLists(destination string) {
	BrowseTopicsListsDir = destination + "/browse-topics-lists"

	if _, err := os.Stat(BrowseTopicsListsDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(BrowseTopicsListsDir, os.FileMode(0755))
		if mkdirErr != nil {
			panic(mkdirErr)
		}
	}

	browseTopicsListsCategories := browseTopicsListsCategories()

	if Source == "database" {
		GenerateBrowseTopicsListsFromDatabase(browseTopicsListsCategories)
	} else if Source == "cache" {
		// Don't know if will be implementing this
		fmt.Println("Generation of topic browse lists pages has not yet been implemented.")
	} else {
		// Should never get here
	}
}

func browseTopicsListsCategories() (browseTopicsListsCategories []BrowseTopicsListsEntry) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabet); i++ {
		letter := string(alphabet[i])
		browseTopicsListsEntry := BrowseTopicsListsEntry{
			Label : strings.ToUpper(letter),
			FileBasename : letter,
			Regexp : "^" + letter,
		}
		browseTopicsListsCategories = append(browseTopicsListsCategories, browseTopicsListsEntry)
	}

	browseTopicsListsCategories = append(browseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "0-9",
		FileBasename : "0-9",
		Regexp : "^[0-9]",
	})
	browseTopicsListsCategories = append(browseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "?#@",
		FileBasename : "non-alphanumeric",
		Regexp : "^[^a-z0-9]",
	})

	return
}

func GenerateBrowseTopicsListsFromDatabase(browseTopicsListsCategories []BrowseTopicsListsEntry) {
	var topicsWithSortKeys []models.TopicsWithSortKeys

	for _, browseTopicsListsCategory := range browseTopicsListsCategories {
		topicsWithSortKeys = db.GetTopicsWithSortKeysByFirstSortableCharacterRegexp(browseTopicsListsCategory.Regexp)
		filename := browseTopicsListsCategory.Label + ".html"
		browseTopicsListPageData := CreateBrowseTopicsListPageData(topicsWithSortKeys)
		WriteBrowseTopicsListPage(filename, browseTopicsListPageData)
	}

}

func CreateBrowseTopicsListPageData(topicsWithSortKeys []models.TopicsWithSortKeys) BrowseTopicsListPageData {
	return nil
}

func WriteBrowseTopicsListPage(filename string, browseTopicsListPageData BrowseTopicsListPageData) (err error){
	tpl := template.New("a-to-z.html")
	tpl, err = tpl.ParseFiles(
		TemplateDirectory + "/a-to-z.html",
		TemplateDirectory + "/a-to-z-content.html",
		TemplateDirectory + "/banner.html",
	)

	if err != nil {
		return err
	}

	file := TopicPagesDir + "/" + filename
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
