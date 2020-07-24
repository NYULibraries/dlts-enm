package cache

import (
	"github.com/nyulibraries/dlts-enm/util"
	"os"
	"path"
	"strconv"
)

// TODO:
//   * Figure out if there's a better place to put this
//   * Allow user to specify cache path
// Tried using os.TempDir(), but it was returning
// /var/folders/dh/48wd7vnj3xqd1w_f126tcnvh0000gn/T/, which was not as convenient.
//   * Use subdirectories:
//       * reload command
//       * sitegentopicpages command
const cache = "/tmp/enm-cache"
const SitegenBrowseTopicListsCache = cache + "/sitegen-browsetopiclists"
const SitegenBrowseTopicListsCategoriesCache = SitegenBrowseTopicListsCache + "/categories"
const SitegenTopicpagesCache = cache + "/sitegen-topicpages"
const SolrLoadCache = cache + "/solr-load"

func init() {
	if _, err := os.Stat(cache); os.IsNotExist(err) {
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

func CacheFile(request string, id string) (cacheFile string) {
	cacheFile = cache + "/" + request;
	if id != "" {
		cacheFile += "-" + id + ".json"
	} else {
		cacheFile += ".json"
	}

	return
}

func SolrLoadCacheFile(isbn string, page string, id int) (cacheFile string) {
	cacheFile = path.Join(SolrLoadCache, isbn, page + "__" + strconv.Itoa(id) + ".json")

	return
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
