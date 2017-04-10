package cache

import (
	"os"
)

// TODO:
//   * Figure out if there's a better place to put this
//   * Allow user to specify cache path
// Tried using os.TempDir(), but it was returning
// /var/folders/dh/48wd7vnj3xqd1w_f126tcnvh0000gn/T/, which was not as convenient.
var Cache = "/tmp/enm-cache"

func init() {
	if _, err := os.Stat(Cache); os.IsNotExist(err) {
		os.Mkdir(Cache, 0700)
	} else if err != nil {
		panic(err.Error())
	}
}

func CacheFile(request string, id string) (cacheFile string) {
	cacheFile = Cache + "/" + request;
	if id != "" {
		cacheFile += "-" + id + ".json"
	} else {
		cacheFile += ".json"
	}

	return
}
