package sitegen

import (
	"testing"
	"os"
	"path"

	"github.com/nyulibraries/dlts-enm/util"
)

func TestGenerateBrowseTopicsLists(t *testing.T) {
	wd, err := os.Getwd()
	if (err != nil) {
		t.Fatal( "os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	BrowseTopicsListsGoldenFilesDirectory := rootDirectory + "/sitegen/testdata/golden/browse-topics-lists"

	destination := rootDirectory + "/sitegen/testdata/tmp"
	outputDir := destination + "/browse-topics-lists"
	err = os.RemoveAll(outputDir)
	if (err != nil) {
		t.Fatal( "os.RemoveAll(" + destination + ") failed: " + err.Error())
	}

	BrowseTopicListsTemplateDirectory = rootDirectory + "/" + BrowseTopicListsTemplateDirectory
	SharedTemplateDirectory = rootDirectory + "/" + SharedTemplateDirectory
	SitePagesTemplateDirectory = rootDirectory + "/" + SitePagesTemplateDirectory
	TopicPageTemplateDirectory = rootDirectory + "/" + TopicPageTemplateDirectory

	Source = "database"

	GenerateBrowseTopicsLists(destination)

	diffOutput, err := util.Diff(BrowseTopicsListsGoldenFilesDirectory, outputDir)
	if (err != nil) {
		t.Fatal("Diff of " + BrowseTopicsListsGoldenFilesDirectory + " and " +
			destination + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		t.Errorf("%s", diffOutput)
	}
}
