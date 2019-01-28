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
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/nyulibraries/dlts-enm/util"
)

func TestGenerateSitePagesNoGoogleAnalytics(t *testing.T) {
	GoogleAnalytics = false

	_, err := testGenerateSitePages()
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateSitePagesWithGoogleAnalytics(t *testing.T) {
	GoogleAnalytics = true

	_, err := testGenerateSitePages()
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func testGenerateSitePages() (bool, error) {
	wd, err := os.Getwd()
	if (err != nil) {
		return false, errors.New("os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	SitePagesGoldenFilesDirectory := rootDirectory + "/sitegen/testdata/golden/site-pages" +
		"/" + getGoldenFileSubdirectory()

	outputDir := rootDirectory + "/sitegen/testdata/tmp/site-pages"
	// Delete all the contents, but don't delete the directory itself
	err = os.RemoveAll(outputDir)
	err = os.MkdirAll(outputDir, 0700)

	if (err != nil) {
		return false, errors.New("os.Remove(" + outputDir + ") failed: " + err.Error())
	}

	// Only do this if another sitegen test hasn't already
	if ! strings.HasPrefix(BrowseTopicListsTemplateDirectory, rootDirectory) {
		BrowseTopicListsTemplateDirectory = rootDirectory + "/" + BrowseTopicListsTemplateDirectory
		SharedTemplateDirectory = rootDirectory + "/" + SharedTemplateDirectory
		SitePagesTemplateDirectory = rootDirectory + "/" + SitePagesTemplateDirectory
		TopicPageTemplateDirectory = rootDirectory + "/" + TopicPageTemplateDirectory
	}

	Source = "database"

	GenerateSitePages(outputDir)

	diffOutput, err := util.Diff(SitePagesGoldenFilesDirectory, outputDir)
	if (err != nil) {
		return false, errors.New("Diff of " + SitePagesGoldenFilesDirectory + " and " +
			outputDir + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		return false, errors.New(diffOutput)
	}

	return true, nil
}
