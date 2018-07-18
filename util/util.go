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
	"os"
	"os/exec"
	"path/filepath"
	"path"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func CompareUsingEnglishCollation(a, b string) (int) {
	cl := collate.New(language.English, collate.Loose)
	return cl.CompareString(a, b)
}

func CreateFileWithAllParentDirectories(file string) (f *os.File, err error) {
	f, err = os.Create(file)
	if err != nil {
		// Create the subdirectories and try again if "no such file or directory" error
		if err.(*os.PathError).Err.Error() == "no such file or directory" {
			os.MkdirAll(filepath.Dir(file), 0755)
			f, err = os.Create(file)
		}
	}

	return
}

func Diff(path1 string, path2 string) (string, error) {
	diffCmd := "diff"

	outputBytes, err := exec.Command(diffCmd, "-r", path1, path2).CombinedOutput()
	if err != nil {
		switch err.(type) {
			case *exec.ExitError:
				// `diff` ran successfully with non-zero exit code.  Report the
				// differences.
			default:
				// `diff` command failed to run.
				return "", err
		}
	}

	return string(outputBytes), nil
}

func GetNormalizedTopicNameForSorting(topicName string) string {
	return strings.ToLower(strings.TrimPrefix(topicName, "\""))
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

func GetTopicIDFromTopicPagePath(topicPagePath string) string {
	filename := path.Base(topicPagePath)
	basename := strings.TrimSuffix(filename, ".html")

	return strings.TrimLeft(basename, "0")
}
