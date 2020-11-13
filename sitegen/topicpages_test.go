package sitegen

import (
	"errors"
	"github.com/nyulibraries/dlts-enm/cache"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nyulibraries/dlts-enm/util"
)

func TestGenerateTopicPagesNoGoogleAnalyticsFromCache(t *testing.T) {
	GoogleAnalytics = false

	cache.SetCacheLocation(testCache)

	_, err := testGenerateTopicPages("cache")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateTopicPagesWithGoogleAnalyticsFromCache(t *testing.T) {
	GoogleAnalytics = true

	cache.SetCacheLocation(testCache)

	_, err := testGenerateTopicPages("cache")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateTopicPagesNoGoogleAnalyticsFromDatabase(t *testing.T) {
	GoogleAnalytics = false

	cache.SetCacheLocation(cache.DefaultCache)

	_, err := testGenerateTopicPages("database")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func TestGenerateTopicPagesWithGoogleAnalyticsFromDatabase(t *testing.T) {
	GoogleAnalytics = true

	cache.SetCacheLocation(cache.DefaultCache)

	_, err := testGenerateTopicPages("database")
	if (err != nil) {
		t.Fatal(err.Error())
	}
}

func testGenerateTopicPages(source string) (bool, error) {
	wd, err := os.Getwd()
	if (err != nil) {
		return false, errors.New("os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	TopicPagesGoldenFilesDirectory := rootDirectory + "/sitegen/testdata/golden/topic-pages" +
		"/" + getGoldenFileSubdirectory()

	destination := rootDirectory + "/sitegen/testdata/tmp/topic-pages/" +
		getGoldenFileSubdirectory()
	outputDir := destination + "/topic-pages"
	err = os.RemoveAll(outputDir)
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

	Source = source

	goldenFiles := []string{}
	err = filepath.Walk(TopicPagesGoldenFilesDirectory, func(path string, f os.FileInfo, err error) error {
		if (f.IsDir()) {
			return nil
		} else if (strings.HasSuffix(path, ".html")) {
			goldenFiles = append(goldenFiles, path)
		} else {
			return errors.New("Unexpected file encountered: " + path)
		}

		return nil
	})
	if err != nil {
		return false, errors.New(err.Error())
	}

	for _, goldenFile := range goldenFiles {
		TopicIDs = append(TopicIDs, util.GetTopicIDFromTopicPagePath(goldenFile))
	}

	GenerateTopicPages(destination)

	diffOutput, err := util.Diff(TopicPagesGoldenFilesDirectory, outputDir)
	if (err != nil) {
		return false, errors.New("Diff of " + TopicPagesGoldenFilesDirectory + " and " +
			destination + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		return false, errors.New(diffOutput)
	}

	return true, nil
}
