package util

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var nyuPressPublisherNameRegex *regexp.Regexp
var standardIdentifierMap map[string]string

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

// The ISBNs used as EPUB identifiers in the TCT are now almost all obsolete --
// see https://jira.nyu.edu/browse/NYUP-753.
// Rather than change all the ISBNs in the Postgres database and the cache files, for
// now we just provide this mapping function.  This map will never change because
// re-mapping only needs to be done for EPUBs that were published to Open Square
// before we standardized on library ISBN as the Open Square identifier.  For simplicity,
// therefore, we just hardcode it in this function.
func GetOpenSquareStandardIdentifierForISBN(isbn string) string {
	if (standardIdentifierMap == nil) {
		// Map: https://raw.githubusercontent.com/NYULibraries/dlts-open-square-standard-identifiers/040b684b021ce33935123549d61ac828c2482ac4/map-of-nonstandard-isbns-to-standard-isbns.json
		standardIdentifierMap = map[string]string{
			"9780814706404": "9780814707821",
			"9780814706657": "9780814707517",
			"9780814711774": "9780814725078",
			"9780814712481": "9780814723418",
			"9780814712771": "9780814786086",
			"9780814712917": "9780814786123",
			"9780814713013": "9780814723425",
			"9780814713266": "9780814709108",
			"9780814714218": "9780814790168",
			"9780814714539": "9780814790144",
			"9780814715123": "9780814723715",
			"9780814715352": "9780814723708",
			"9780814715383": "9780814772195",
			"9780814715635": "9780814790175",
			"9780814718124": "9780814744147",
			"9780814718766": "9780814720981",
			"9780814718803": "9780814721100",
			"9780814726815": "9780814728901",
			"9780814726846": "9780814728048",
			"9780814730911": "9780814733486",
			"9780814731437": "9780814738573",
			"9780814731956": "9780814733158",
			"9780814735053": "9780814773130",
			"9780814735084": "9780814744758",
			"9780814735190": "9780814744772",
			"9780814735206": "9780814744819",
			"9780814735237": "9780814744840",
			"9780814735282": "9780814744871",
			"9780814735305": "9780814773215",
			"9780814735336": "9780814744789",
			"9780814735923": "9780814744765",
			"9780814742297": "9780814743980",
			"9780814742303": "9780814743959",
			"9780814742358": "9780814743973",
			"9780814746622": "9780814749234",
			"9780814746677": "9780814763582",
			"9780814746929": "9780814763551",
			"9780814747148": "9780814763520",
			"9780814750957": "9781479898626",
			"9780814751008": "9780814752685",
			"9780814751213": "9780814752715",
			"9780814755112": "9780814763223",
			"9780814755297": "9780814763148",
			"9780814755471": "9780814763179",
			"9780814755969": "9780814759714",
			"9780814757970": "9780814759271",
			"9780814761908": "9780814762622",
			"9780814766569": "9780814767917",
			"9780814774410": "9780814769447",
			"9780814774434": "9780814769485",
			"9780814774458": "9780814771518",
			"9780814774632": "9780814769423",
			"9780814774694": "9780814776636",
			"9780814774755": "9780814771501",
			"9780814774816": "9780814769461",
			"9780814774823": "9780814769409",
			"9780814779163": "9780814788745",
			"9780814779170": "9780814741498",
			"9780814779965": "9780814771037",
			"9780814780015": "9780814788806",
			"9780814780213": "9780814788707",
			"9780814780978": "9780814739617",
			"9780814782194": "9780814784488",
			"9780814787922": "9780814788462",
			"9780814793114": "9780814784600",
			"9780814793398": "9780814784891",
			"9781479804948": "9781479882281",
			"9781479820375": "9781479811908",
			"9781479824243": "9781479863570",
			"9781479835737": "9781479891672",
			"9781479849857": "9781479888788",
			"9781479852758": "9781479888900",
			"9781479868148": "9781479868148",
			"9781479892464": "9781479807185",
			"9781479899982": "9781479829712",
		}
	}

	return standardIdentifierMap[isbn]
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

func IsNYUPress(publisherName string) bool {
	if (nyuPressPublisherNameRegex == nil) {
		nyuPressPublisherNameRegex = regexp.MustCompile(`(?i).*new york university.*`)
	}

	return (nyuPressPublisherNameRegex.MatchString(publisherName)) ||
		// Teaching What You're Not: Identity Politics in Higher Education
		// by Katherine Mayberry had this incorrect <dc:publisher> value in its OPF
		// file when it was ingested into the TCT.
		(publisherName == "York University Press")
}
