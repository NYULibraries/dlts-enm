// Copyright © 2017 NYU
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
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"strconv"

	"github.com/nyulibraries/dlts-enm/cache"
	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/util"
	"sort"
)

type EPUBMatch struct{
	Title string
	Authors string
	Publisher string
	ISBN string
	NumberOfOccurrences int
}

type ExternalRelation struct{
	Relationship string
	URL string
	Vocabulary string
}

type Link struct{
	Source int `json:"source"`
	Target int `json:"target"`
}

type Node struct{
	ID int `json:"id"`
	Name string `json:"name"`
	OCount int `json:"ocount"`
	Path string `json:"path"`
}

type RelatedTopic struct{
	ID int
	Name string
	NumberOfOccurrences int
}

type TopicPageData struct{
	AlternateNames []string
	DisplayName string
	EPUBMatches []EPUBMatch
	ExternalRelations []ExternalRelation
	Paths Paths
	RelatedTopics []RelatedTopic
	TopicID int
	VisualizationData VisualizationData
}

type VisualizationData struct{
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

var tpl *template.Template

var TopicIDs []string
var TopicPagesDir string

func GenerateTopicPages(destination string) {
	TopicPagesDir = destination + "/topic-pages"
	if _, err := os.Stat(TopicPagesDir); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(TopicPagesDir, os.FileMode(0755))
		if mkdirErr != nil {
			panic(mkdirErr)
		}
	}

	// Caching the template for speed
	tpl = NewTemplate()

	if Source == "database" {
		GenerateTopicPagesFromDatabase()
	} else if Source == "cache" {
		GenerateTopicPagesFromCache()
	} else {
		// Should never get here
	}
}

func NewTemplate() (tpl *template.Template){
	funcs := template.FuncMap{
		"stringsJoin": func (arrayOfStrings []string) template.HTML{
			return template.HTML(strings.Join(arrayOfStrings, "&nbsp;&bull;&nbsp;"))
		},
		"lastIndex": func (s []ExternalRelation) int {
			return len(s) - 1;
		},
	}
	tpl = template.New("index.html").Funcs(funcs)
	tpl, err := tpl.ParseFiles(
		TopicPageTemplateDirectory + "/index.html",
		TopicPageTemplateDirectory + "/epub.html",
		SharedTemplateDirectory + "/banner.html",
	)
	if err != nil {
		panic(err)
	}

	return
}

func GenerateTopicPagesFromCache(){

	var cacheFiles []string

	if (len(TopicIDs) > 0) {
		for _, topicID := range TopicIDs {
			intTopicID, err := strconv.Atoi(topicID)
			if (err != nil) {
				panic(err)
			}
			cacheFiles = append(cacheFiles, cache.SitegenTopicpagesCacheFile(intTopicID))
		}
	} else {
		err := filepath.Walk(cache.SitegenTopicpagesCache, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) == ".json" {
				cacheFiles = append(cacheFiles, path)
			}

			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	var topicPageData TopicPageData
	for _, cacheFile := range cacheFiles {
		jsonBytes, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonBytes,&topicPageData)
		if err != nil {
			panic(err)
		}

		err = WritePage(topicPageData)
		if err != nil {
			panic(err)
		}
	}
}

func GenerateTopicPagesFromDatabase() {
	var topicsWithAlternateNames []*db.TopicAlternateName
	if (len(TopicIDs) > 0) {
		topicsWithAlternateNames = db.GetTopicsWithAlternateNamesByTopicIDs(TopicIDs)
	} else {
		topicsWithAlternateNames = db.GetTopicsWithAlternateNamesAll()
	}

	SortTopicsWithAlternateNames(topicsWithAlternateNames)

	var alternateNames []string
	inProgressTopicID := topicsWithAlternateNames[0].TctID
	inProgressTopicName := topicsWithAlternateNames[0].DisplayName
	for _, topicWithAlternateNames := range topicsWithAlternateNames{
		// Finished getting alternate names for in-process topic.
		if (topicWithAlternateNames.TctID != inProgressTopicID) {
			// Generate the topic page.
			err := GenerateTopicPage(inProgressTopicID, inProgressTopicName, alternateNames)
			if err != nil {
				panic(fmt.Sprintf("ERROR: couldn't generate topic page for %d: \"%s\"\n", inProgressTopicID, err.Error()))
			}

			// Clear alternate names and set new in progress topic.
			alternateNames = []string{}
			inProgressTopicID = topicWithAlternateNames.TctID
			inProgressTopicName = topicWithAlternateNames.DisplayName
		}

		if (topicWithAlternateNames.Name != inProgressTopicName) {
			alternateNames = append(alternateNames, topicWithAlternateNames.Name)
		}
	}

	// Generate final page
	err := GenerateTopicPage(inProgressTopicID, inProgressTopicName, alternateNames)
	if err != nil {
		panic(fmt.Sprintf("ERROR: couldn't generate topic page for %d: \"%s\"\n", inProgressTopicID, err.Error()))
	}
}

func GenerateTopicPage(topicID int, topicDisplayName string, alternateNames []string) error {
	visualizationData := VisualizationData{}
	relatedTopics := []RelatedTopic{}
	relatedTopicNamesWithNumberOfOccurrences := db.GetRelatedTopicNamesForTopicWithNumberOfOccurrences(topicID)

	SortRelatedTopicNamesWithNumberOfOccurrences(relatedTopicNamesWithNumberOfOccurrences)

	for _, relatedTopic := range relatedTopicNamesWithNumberOfOccurrences {
		relatedTopics = append(relatedTopics, RelatedTopic{
			ID: relatedTopic.ID,
			Name: relatedTopic.DisplayName,
			NumberOfOccurrences: int(relatedTopic.NumberOfOccurrences),
		})

		visualizationData.Links = append(visualizationData.Links, Link{
			Source: topicID,
			Target: relatedTopic.ID,
		})

		visualizationData.Nodes = append(visualizationData.Nodes, Node{
			Name: relatedTopic.DisplayName,
			ID: relatedTopic.ID,
			OCount: int(relatedTopic.NumberOfOccurrences),
			Path: GetRelativeFilepathForTopicPage(relatedTopic.ID),
		})
	}

	visualizationData.Nodes = append(visualizationData.Nodes, Node{
		Name: topicDisplayName,
		ID: topicID,
		OCount: db.GetTopicNumberOfOccurrencesByTopicId(topicID),
		Path: GetRelativeFilepathForTopicPage(topicID),
	})

	if (visualizationData.Links == nil) {
		visualizationData.Links = []Link{}
	}

	epubMatches := []EPUBMatch{}
	for _, epubForTopicWithNumberOfMatchedPages := range db.GetEpubsForTopicWithNumberOfMatchedPages(topicID) {
		epubMatches = append(epubMatches, EPUBMatch{
			Title:               epubForTopicWithNumberOfMatchedPages.Title,
			Authors:             epubForTopicWithNumberOfMatchedPages.Author,
			Publisher:           epubForTopicWithNumberOfMatchedPages.Publisher,
			ISBN:                epubForTopicWithNumberOfMatchedPages.Isbn,
			NumberOfOccurrences: int(epubForTopicWithNumberOfMatchedPages.NumberOfOccurrences),
		})
	}

	externalRelations := []ExternalRelation{}
	for _, externalRelationsForTopic := range db.GetExternalRelationsForTopics(topicID) {
		externalRelations = append(externalRelations, ExternalRelation{
			Relationship: externalRelationsForTopic.Relationship,
			URL: externalRelationsForTopic.URL,
			Vocabulary: externalRelationsForTopic.Vocabulary,
		})
	}

	topicPageData := TopicPageData{
		AlternateNames: alternateNames,
		DisplayName: topicDisplayName,
		EPUBMatches: epubMatches,
		ExternalRelations: externalRelations,
		Paths: Paths{
			WebRoot: "../../../../..",
		},
		RelatedTopics: relatedTopics,
		TopicID: topicID,
		VisualizationData: visualizationData,
	}

	err := WriteCacheFile(topicPageData)
	if err != nil {
		panic(err)
	}

	err = WritePage(topicPageData)
	if err != nil {
		panic(err)
	}

	return nil
}

func GetRelativeFilepathForTopicPage(topicID int) string {
	return util.GetRelativeFilepathInLargeDirectoryTree("", topicID, ".html")
}

func WriteCacheFile(topicPageData TopicPageData) (err error){
	topicPageDataJSON, err := json.MarshalIndent(topicPageData,"","    ")
	if err != nil {
		panic(err)
	}

	cacheFile := cache.SitegenTopicpagesCacheFile(topicPageData.TopicID)
	f, err := util.CreateFileWithAllParentDirectories(cacheFile)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(topicPageDataJSON)
	if err != nil {
		panic(err)
	}

	return nil
}

func WritePage(topicPageData TopicPageData) (err error){
	filename := TopicPagesDir + "/" + GetRelativeFilepathForTopicPage(topicPageData.TopicID)
	f, err := util.CreateFileWithAllParentDirectories(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tpl.Execute(f, topicPageData)
	if err != nil {
		return err
	}

	return nil
}

func SortRelatedTopicNamesWithNumberOfOccurrences(relatedTopicNamesWithNumberOfOccurrences []*db.RelatedTopicNamesForTopicWithNumberOfOccurrences) {
	sort.Slice(relatedTopicNamesWithNumberOfOccurrences, func(i, j int) bool {
		return util.CompareUsingEnglishCollation(
			util.GetNormalizedTopicNameForSorting(relatedTopicNamesWithNumberOfOccurrences[i].DisplayName),
			util.GetNormalizedTopicNameForSorting(relatedTopicNamesWithNumberOfOccurrences[j].DisplayName)) == -1
	} )
}

func SortTopicsWithAlternateNames(topicAlternateNames []*db.TopicAlternateName) {
	sort.Slice(topicAlternateNames, func(i, j int) bool {
		return util.CompareUsingEnglishCollation(
			util.GetNormalizedTopicNameForSorting(topicAlternateNames[i].Name),
			util.GetNormalizedTopicNameForSorting(topicAlternateNames[j].Name)) == -1
	} )

	sort.Slice(topicAlternateNames, func(i, j int) bool {
		return util.CompareUsingEnglishCollation(
			util.GetNormalizedTopicNameForSorting(topicAlternateNames[i].DisplayName),
			util.GetNormalizedTopicNameForSorting(topicAlternateNames[j].DisplayName)) == -1
	} )
}