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
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nyulibraries/dlts-enm/util"
)

func TestGenerateTopicPages(t *testing.T) {
	wd, err := os.Getwd()
	if (err != nil) {
		t.Fatal( "os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	TopicPagesGoldenFilesDirectory := rootDirectory + "/sitegen/testdata/golden/topic-pages"

	destination := rootDirectory + "/sitegen/testdata/tmp"
	outputDir := destination + "/topic-pages"
	err = os.RemoveAll(outputDir)
	if (err != nil) {
		t.Fatal( "os.Remove(" + destination + ") failed: " + err.Error())
	}

	// Only do this if another sitegen test hasn't already
	if ! strings.HasPrefix(BrowseTopicListsTemplateDirectory, rootDirectory) {
		BrowseTopicListsTemplateDirectory = rootDirectory + "/" + BrowseTopicListsTemplateDirectory
		SharedTemplateDirectory = rootDirectory + "/" + SharedTemplateDirectory
		SitePagesTemplateDirectory = rootDirectory + "/" + SitePagesTemplateDirectory
		TopicPageTemplateDirectory = rootDirectory + "/" + TopicPageTemplateDirectory
	}

	Source = "database"

	goldenFiles := []string{}
	err = filepath.Walk(TopicPagesGoldenFilesDirectory, func(path string, f os.FileInfo, err error) error {
		if (f.IsDir()) {
			return nil
		} else if (strings.HasSuffix(path, ".html")) {
			goldenFiles = append(goldenFiles, path)
		} else {
			t.Fatal("Unexpected file encountered: " + path)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, goldenFile := range goldenFiles {
		TopicIDs = append(TopicIDs, util.GetTopicIDFromTopicPagePath(goldenFile))
	}

	GenerateTopicPages(destination)

	diffOutput, err := util.Diff(TopicPagesGoldenFilesDirectory, outputDir)
	if (err != nil) {
		t.Fatal("Diff of " + TopicPagesGoldenFilesDirectory + " and " +
			destination + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		t.Errorf("%s", diffOutput)
	}
}
