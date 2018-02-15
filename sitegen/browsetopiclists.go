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
	_ "html/template"
	_ "io/ioutil"
	"fmt"
	"os"

	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/db/models"
)

// Tricky...this assumes that location of the templates relative to working directory
// matches what the location would be if calling `go run main.go` from root of this
// repo.  This is brittle: user could call `go run ../main.go` from sitegen/, or
// could be calling the built binary from anywhere relative to templates.
// Consider including the templates in the binary using something like
// https://github.com/jteeuwen/go-bindata
const BrowseTopicListsTemplateDirectory = "sitegen/templates/browse-topic-lists"

var BrowseTopicListsDir string

func GenerateBrowseTopicLists(destination string) {
	BrowseTopicListsDir = destination + "/browse-topic-lists"

	if _, err := os.Stat(BrowseTopicListsDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(BrowseTopicListsDir, os.FileMode(0755))
		if mkdirErr != nil {
			panic(mkdirErr)
		}
	}

	if Source == "database" {
		GenerateBrowseTopicListsFromDatabase()
	} else if Source == "cache" {
		// Don't know if will be implementing this
		fmt.Println("Generation of topic browse lists pages has not yet been implemented.")
	} else {
		// Should never get here
	}
}

func GenerateBrowseTopicListsFromDatabase() {
	var topicsWithSortKeys []models.TopicsWithSortKeys

	topicsWithSortKeys = db.GetTopicsWithSortKeysByFirstSortableCharacterRegexp("^a")
	fmt.Println(topicsWithSortKeys)
}