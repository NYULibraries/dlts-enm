package cache

import (
	"github.com/nyulibraries/dlts-enm/util"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

// TODO:
//   * Figure out if there's a better place to put this
// Tried using os.TempDir(), but it was returning
// /var/folders/dh/48wd7vnj3xqd1w_f126tcnvh0000gn/T/, which was not as convenient.
var cache = "/tmp/enm-cache"
var SitegenBrowseTopicListsCache string
var SitegenBrowseTopicListsCategoriesCache string
var SitegenTopicpagesCache string
var SolrLoadCache string
var TCTResponseCache string

func init() {
	// User can override the default cache using environment variable
	cacheEnvVar := os.Getenv("ENM_CACHE")
	if cacheEnvVar != "" {
		// Need to declare this otherwise cache var is shadowed through use of :=
		var err error
		cache, err = filepath.Abs(cacheEnvVar)
		if (err != nil) {
			panic(err)
		}
	}

	SitegenBrowseTopicListsCache = cache + "/sitegen-browsetopiclists"
	SitegenBrowseTopicListsCategoriesCache = SitegenBrowseTopicListsCache + "/categories"
	SitegenTopicpagesCache = cache + "/sitegen-topicpages"
	SolrLoadCache = cache + "/solr-load"
	TCTResponseCache = cache + "/tct"

	if _, err := os.Stat(cache); os.IsNotExist(err) {
		// This will create SitegenBrowseTopicListsCache at the same time.
		err := os.MkdirAll(SitegenBrowseTopicListsCategoriesCache, 0700)
		if (err != nil) {
			panic(err)
		}

		err = os.MkdirAll(SitegenTopicpagesCache, 0700)
		if (err != nil) {
			panic(err)
		}

		err = os.MkdirAll(SolrLoadCache, 0700)
		if (err != nil) {
			panic(err)
		}
	} else if err != nil {
		panic(err.Error())
	}
}

func SitegenBrowseTopicListsCategoriesCacheFile() (cacheFile string) {
	cacheFile = path.Join(SitegenBrowseTopicListsCache, "categories.json")

	return
}

func SitegenBrowseTopicListsCategoryCacheFile(categoryFileBasename string) (cacheFile string) {
	cacheFile = path.Join(SitegenBrowseTopicListsCategoriesCache, categoryFileBasename+ ".json")

	return
}

func SitegenTopicpagesCacheFile(topicID int) (cacheFile string) {
	cacheFile = path.Join(SitegenTopicpagesCache,
		util.GetRelativeFilepathInLargeDirectoryTree("", topicID, ".json"))

	return
}

func SolrLoadCacheFile(isbn string, page string, id int) (cacheFile string) {
	cacheFile = path.Join(SolrLoadCache, isbn, page+"__"+strconv.Itoa(id)+".json")

	return
}

func TCTCacheFile(request string, id string) (cacheFile string) {
	cacheFile = TCTResponseCache + "/" + request;
	if id != "" {
		cacheFile += "-" + id + ".json"
	} else {
		cacheFile += ".json"
	}

	return
}
