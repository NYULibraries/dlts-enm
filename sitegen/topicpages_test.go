package sitegen

import (
	"path/filepath"
	"testing"
	"strings"
)

const TopicPagesGoldenFilesDirectory = "testdata/topicpages"

func TestGenerateTopicPages(t *testing.T) {
	Source = "cache"
	goldenFiles, err := filepath.Glob(TopicPagesGoldenFilesDirectory + "/*.golden")
	if err != nil {
		t.Fatal(err)
	}

	for _, goldenFile := range goldenFiles {
		pieces := strings.SplitN(goldenFile, ".", 2)
		TopicIDs = append(TopicIDs, strings.TrimLeft(pieces[0],TopicPagesGoldenFilesDirectory + "/"))
	}
}