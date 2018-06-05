package sitegen

import (
	"path/filepath"
	"testing"
	"strings"
	"os"
	"path"

	"github.com/nyulibraries/dlts-enm/util"
)

const TopicPagesGoldenFilesDirectory = "testdata/golden/topic-pages"

func TestGenerateTopicPages(t *testing.T) {
	wd, err := os.Getwd()
	if (err != nil) {
		t.Fatal( "os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	destination := rootDirectory + "/sitegen/testdata/tmp"
	err = os.RemoveAll(destination)
	if (err != nil) {
		t.Fatal( "os.Remove(" + destination + ") failed: " + err.Error())
	}

	Source = "cache"
	BrowseTopicListsTemplateDirectory = rootDirectory + "/" + BrowseTopicListsTemplateDirectory
	SharedTemplateDirectory = rootDirectory + "/" + SharedTemplateDirectory
	SitePagesTemplateDirectory = rootDirectory + "/" + SitePagesTemplateDirectory
	TopicPageTemplateDirectory = rootDirectory + "/" + TopicPageTemplateDirectory

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
}
