package sitegen

import (
	"encoding/json"
	"github.com/nyulibraries/dlts-enm/cache"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/util"
)

type BrowseTopicsListPageData struct{
	ActiveNavbarTab string
	NavbarTabs []BrowseTopicsListsEntry
	Paths struct{
		WebRoot string
	}
	Topics []BrowseTopicsListPageDataTopic
}

type BrowseTopicsListPageDataTopic struct{
	Name string
	URL string
}

type BrowseTopicsListsEntry struct{
	Label string
	FileBasename string
	Regexp string
}

var BrowseTopicsListsDir string

func GenerateBrowseTopicsLists(destination string) {
	BrowseTopicsListsDir = destination + "/browse-topics-lists"

	if _, err := os.Stat(BrowseTopicsListsDir); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(BrowseTopicsListsDir, os.FileMode(0755))
		if mkdirErr != nil {
			panic(mkdirErr)
		}
	}

	if Source == "database" {
		db.InitDB()
		GenerateDynamicBrowseTopicsListsFromDatabase()
	} else if Source == "cache" {
		GenerateDynamicBrowseTopicsListsFromCache()
	} else {
		// Should never get here
	}

	// These browse topics lists are currently unchanging, with all topics
	// hardcoded into the templates.
	WriteStaticBrowseTopicsListsPage("enm-picks")
}

func getBrowseTopicsListsCategories() (browseTopicsListsCategories []BrowseTopicsListsEntry){
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabet); i++ {
		letter := string(alphabet[i])
		browseTopicsListsEntry := BrowseTopicsListsEntry{
			Label : strings.ToUpper(letter),
			FileBasename : letter,
			Regexp : letter,
		}
		browseTopicsListsCategories = append(browseTopicsListsCategories, browseTopicsListsEntry)
	}

	browseTopicsListsCategories = append(browseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "0-9",
		FileBasename : "0-9",
		Regexp : "[0-9]",
	})
	browseTopicsListsCategories = append(browseTopicsListsCategories, BrowseTopicsListsEntry{
		Label : "?#@",
		FileBasename : "non-alphanumeric",
		Regexp : "[^a-z0-9]",
	})

	return
}

func GenerateDynamicBrowseTopicsListsFromCache() {
	jsonBytes, err := ioutil.ReadFile(cache.SitegenBrowseTopicListsCategoriesCacheFile())
	if err != nil {
		panic(err)
	}

	var browseTopicsListsCategories []BrowseTopicsListsEntry
	err = json.Unmarshal(jsonBytes,&browseTopicsListsCategories)
	if err != nil {
		panic(err)
	}

	for _, browseTopicsListsCategory := range browseTopicsListsCategories {
		jsonBytes, err := ioutil.ReadFile(cache.SitegenBrowseTopicListsCategoryCacheFile(browseTopicsListsCategory.FileBasename))
		if err != nil {
			panic(err)
		}

		var browseTopicsListPageData BrowseTopicsListPageData
		err = json.Unmarshal(jsonBytes,&browseTopicsListPageData)
		if err != nil {
			panic(err)
		}

		filename := browseTopicsListsCategory.FileBasename + ".html"

		err = WriteDynamicBrowseTopicsListPage(filename, browseTopicsListPageData)
		if (err != nil) {
			panic(err)
		}
	}

}

func GenerateDynamicBrowseTopicsListsFromDatabase() {
	var topics []db.Topic

	browseTopicsListsCategories := getBrowseTopicsListsCategories()

	WriteCategoriesCacheFile(browseTopicsListsCategories)

	for _, browseTopicsListsCategory := range browseTopicsListsCategories {
		topics = db.GetTopicsWithSortKeysByFirstSortableCharacterRegexp(browseTopicsListsCategory.Regexp)

		// Must use SliceStable to ensure deterministic sort for tests.
		sort.SliceStable(topics, func(i, j int) bool {
			// Note that i is the higher index, it is j+1
			return util.CompareUsingEnglishCollation(
				topics[i].DisplayNameSortKey, topics[j].DisplayNameSortKey) == -1
		} )

		filename := browseTopicsListsCategory.FileBasename + ".html"
		browseTopicsListPageData :=
			CreateBrowseTopicsListPageData(topics, browseTopicsListsCategory.Label, browseTopicsListsCategories)

		WriteCategoryCacheFile(browseTopicsListsCategory.FileBasename, browseTopicsListPageData)

		err := WriteDynamicBrowseTopicsListPage(filename, browseTopicsListPageData)
		if (err != nil) {
			panic(err)
		}
	}

}

func CreateBrowseTopicsListPageData(
	topicsFromDatabase []db.Topic,
	browseTopicsListsCategoryLabel string,
	browseTopicsListsCategories []BrowseTopicsListsEntry) (browseTopicsListPageData BrowseTopicsListPageData) {
	for _, topicFromDatabase := range topicsFromDatabase {

		topic := BrowseTopicsListPageDataTopic{
			Name: topicFromDatabase.DisplayName,
			URL:  "../topic-pages/" + GetRelativeFilepathForTopicPage(topicFromDatabase.ID),
		}
		browseTopicsListPageData.Topics = append(browseTopicsListPageData.Topics, topic)
	}

	browseTopicsListPageData.ActiveNavbarTab = browseTopicsListsCategoryLabel
	browseTopicsListPageData.NavbarTabs = browseTopicsListsCategories
	browseTopicsListPageData.Paths = Paths{ WebRoot: ".." }

	return
}

func WriteCategoriesCacheFile(browseTopicsListsCategories []BrowseTopicsListsEntry) (err error){
	browseTopicsListsCategoriesJSON, err := json.MarshalIndent(browseTopicsListsCategories,"","    ")
	if err != nil {
		panic(err)
	}

	cacheFile := cache.SitegenBrowseTopicListsCategoriesCacheFile()
	f, err := util.CreateFileWithAllParentDirectories(cacheFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(browseTopicsListsCategoriesJSON)
	if err != nil {
		panic(err)
	}

	return nil
}

func WriteCategoryCacheFile(categoryFileBasename string, browseTopicsListPageData BrowseTopicsListPageData) (err error){
	browseTopicsListsCategoryJSON, err := json.MarshalIndent(browseTopicsListPageData,"","    ")
	if err != nil {
		panic(err)
	}

	cacheFile := cache.SitegenBrowseTopicListsCategoryCacheFile(categoryFileBasename)
	f, err := util.CreateFileWithAllParentDirectories(cacheFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(browseTopicsListsCategoryJSON)
	if err != nil {
		panic(err)
	}

	return nil
}

func WriteStaticBrowseTopicsListsPage(listname string) (err error){
	var filename = listname + ".html"
	var templateFile = listname + ".html"

	tpl := template.New(filename)
	tpl, err = tpl.ParseFiles(
		BrowseTopicListsTemplateDirectory + "/" + templateFile,
		SharedTemplateDirectory    + "/banner.html",
		SharedTemplateDirectory    + "/google-analytics.html",
	)

	if err != nil {
		return err
	}

	file := BrowseTopicsListsDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	pageData := struct{
		GoogleAnalytics bool
		Paths
	}{
		GoogleAnalytics: GoogleAnalytics,
		Paths: Paths{
			WebRoot: "..",
		},
	}

	err = tpl.Execute(f, pageData)
	if err != nil {
		return err
	}

	return nil
}

func WriteDynamicBrowseTopicsListPage(filename string, browseTopicsListPageData BrowseTopicsListPageData) (err error){
	tpl := template.New("a-to-z.html")
	tpl, err = tpl.ParseFiles(
		BrowseTopicListsTemplateDirectory + "/a-to-z.html",
		BrowseTopicListsTemplateDirectory + "/a-to-z-content.html",
		SharedTemplateDirectory           + "/banner.html",
		SharedTemplateDirectory           + "/google-analytics.html",
	)

	if err != nil {
		return err
	}

	file := BrowseTopicsListsDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	pageData := struct{
		BrowseTopicsListPageData
		GoogleAnalytics bool
	}{
		browseTopicsListPageData,
		GoogleAnalytics,
	}
	err = tpl.Execute(f, pageData)
	if err != nil {
		return err
	}

	return nil
}
