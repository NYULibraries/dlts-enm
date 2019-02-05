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
	// Do not use collate.IgnoreCase or collate.Loose options.  We need to be able
	// to distinguish between topic names like Culture and culture
	cl := collate.New(language.English, collate.IgnoreDiacritics, collate.IgnoreWidth)
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
	// Retain case-sensitivity, as topic names can differ only by case.
	// Example: "Culture" and "culture" are distinct topic names.
	return strings.TrimPrefix(topicName, "\"")
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
