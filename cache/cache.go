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

package cache

import (
	"os"
	"github.com/nyulibraries/dlts-enm/util"
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

func init() {
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		os.MkdirAll(SitegenTopicpagesCache, 0700)
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

func SitegenTopicpagesCacheFile(topicID int) (cacheFile string) {
	cacheFile = SitegenTopicpagesCache +
		        "/" +
		        util.GetRelativeFilepathInLargeDirectoryTree("", topicID, ".json")

	return
}
