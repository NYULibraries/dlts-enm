package cache

import (
	"github.com/nyulibraries/dlts-enm/util"
	"os"
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
const SitegenTopicpagesCache = cache + "/sitegen-topicpages"
const SolrLoadCache = cache + "/solr-load"

func init() {
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		err := os.MkdirAll(SitegenTopicpagesCache, 0700)
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

func SolrLoadCacheFile(isbn string, page string) (cacheFile string) {
	cacheFile = SolrLoadCache + "/" + isbn + "/" + page + ".json"

	return
}

func SitegenTopicpagesCacheFile(topicID int) (cacheFile string) {
	cacheFile = SitegenTopicpagesCache +
		        "/" +
		        util.GetRelativeFilepathInLargeDirectoryTree("", topicID, ".json")

	return
}
