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
