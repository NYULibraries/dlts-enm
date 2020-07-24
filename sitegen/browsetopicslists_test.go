package sitegen

import (
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/nyulibraries/dlts-enm/util"
)

func TestGenerateBrowseTopicsListsNoGoogleAnalyticsFromCache(t *testing.T) {
	GoogleAnalytics = false

	_, err := testGenerateBrowseTopicsLists("cache")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateBrowseTopicsListsWithGoogleAnalyticsFromCache(t *testing.T) {
	GoogleAnalytics = true

	_, err := testGenerateBrowseTopicsLists("cache")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateBrowseTopicsListsNoGoogleAnalyticsFromDatabase(t *testing.T) {
	GoogleAnalytics = false

	_, err := testGenerateBrowseTopicsLists("database")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateBrowseTopicsListsWithGoogleAnalyticsFromDatabase(t *testing.T) {
	GoogleAnalytics = true

	_, err := testGenerateBrowseTopicsLists("database")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func testGenerateBrowseTopicsLists(source string) (bool, error) {
	wd, err := os.Getwd()
	if (err != nil) {
		return false, errors.New("os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	BrowseTopicsListsGoldenFilesDirectory := rootDirectory + "/sitegen/testdata/golden/browse-topics-lists" +
		"/" + getGoldenFileSubdirectory()

	destination := rootDirectory + "/sitegen/testdata/tmp/browse-topics-lists/" +
		getGoldenFileSubdirectory()
	outputDir := destination + "/browse-topics-lists"
	err = os.RemoveAll(outputDir)
	if (err != nil) {
		return false, errors.New("os.RemoveAll(" + outputDir + ") failed: " + err.Error())
	}

	// Only do this if another sitegen test hasn't already
	if ! strings.HasPrefix(BrowseTopicListsTemplateDirectory, rootDirectory) {
		BrowseTopicListsTemplateDirectory = rootDirectory + "/" + BrowseTopicListsTemplateDirectory
		SharedTemplateDirectory = rootDirectory + "/" + SharedTemplateDirectory
		SitePagesTemplateDirectory = rootDirectory + "/" + SitePagesTemplateDirectory
		TopicPageTemplateDirectory = rootDirectory + "/" + TopicPageTemplateDirectory
	}

	Source = source

	GenerateBrowseTopicsLists(destination)

	diffOutput, err := util.Diff(BrowseTopicsListsGoldenFilesDirectory, outputDir)
	if (err != nil) {
		return false, errors.New("Diff of " + BrowseTopicsListsGoldenFilesDirectory +
			" and " +  destination + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		return false, errors.New(diffOutput)
	}

	return true, nil
}
