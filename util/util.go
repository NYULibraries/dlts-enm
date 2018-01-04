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

package util

import (
	"fmt"
	"strings"
)

func GetMapKeys(m map[string]string) (keys []string) {
	keys = make([]string, len(m))

	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return
}

func GetRelativeFilepathInLargeDirectoryTree(prefix string, ID int, extension string) string {
	zeroPaddedString := fmt.Sprintf("%010d", ID)
	filename := prefix + zeroPaddedString + extension

	return zeroPaddedString[0:2] +
		"/" +
		zeroPaddedString[2:4] +
		"/" +
		zeroPaddedString[4:6] +
		"/" +
		zeroPaddedString[6:8] +
		"/" +
		filename
}

func SnakeToCamelCase(snakeCaseString string) (camelCaseString string){
	tokens := strings.Split(snakeCaseString, "_")

	for _, token := range tokens {
		camelCaseString += strings.Title(token)
	}

	return
}
