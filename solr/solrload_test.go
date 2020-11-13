package solr

import (
	"encoding/json"
	"fmt"
	"github.com/nyulibraries/dlts-enm/cache"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/rtt/Go-Solr"

	"github.com/nyulibraries/dlts-enm/util"
)

const testCache = "./testdata/cache"

type solrStub struct {
	URL     string
	Version []int
}

var goldenFileLocationIDs map[int]bool
var solrGoldenFilesDirectory string
var solrLoadTestTmpDirectory string

func (s solrStub) Update(m map[string]interface{}, commit bool) (*solr.UpdateResponse, error) {
	locationID := m[ "add" ].([]interface{})[0].(interface{}).(map[string]interface{})["id"].(int)

	// Write out files for later comparison with golden files
	if (goldenFileLocationIDs[locationID] == true) {
		b, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(solrLoadTestTmpDirectory+ "/" + strconv.Itoa(locationID) + ".json", b, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		// Do nothing
	}

	return nil, nil
}

// Use same signature as Go Solr solr.Init()
func newSolrStub(host string, port int, core string) (*solrStub, error) {
	url := fmt.Sprintf("http://%s:%d/solr/%s", host, port, core)

	return &solrStub{URL: url}, nil
}

func getGoldenFileLocationIDs() {
	files, err := ioutil.ReadDir(solrGoldenFilesDirectory)
	if err != nil {
		log.Fatal(err)
	}

	goldenFileLocationIDs = make(map[int]bool)

	for _, file := range files {
		var locationID, err = strconv.Atoi(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
		if err != nil {
			log.Fatal(err)
		}

		goldenFileLocationIDs[locationID] = true
	}
}

func TestLoadFromCache(t *testing.T) {
	cache.SetCacheLocation(testCache)

	_Load(t, "cache")
}

func TestLoadFromDatabase(t *testing.T) {
	cache.SetCacheLocation(cache.DefaultCache)

	_Load(t, "database")
}

func _Load(t *testing.T, source string) {
	solrStub, err := newSolrStub("fakesolr.dlib.nyu.edu", 8983, "enm-pages")
	if err != nil {
		t.Fatal(err)
	}

	Init("", 0,  solrStub)

	wd, err := os.Getwd()
	if (err != nil) {
		t.Fatal( "os.Getwd() failed: " + err.Error())
	}

	rootDirectory := path.Dir(wd)

	solrGoldenFilesDirectory = rootDirectory + "/solr/testdata/golden"
	getGoldenFileLocationIDs()

	solrLoadTestTmpDirectory = rootDirectory + "/solr/testdata/tmp/load/" + source
	err = os.RemoveAll(solrLoadTestTmpDirectory)
	if (err != nil) {
		t.Fatal( "os.RemoveAll(" + solrLoadTestTmpDirectory + ") failed: " + err.Error())
	}

	mkdirErr := os.MkdirAll(solrLoadTestTmpDirectory, os.FileMode(0755))
	if mkdirErr != nil {
		t.Fatal(mkdirErr)
	}

	Source = source

	// Load only pages for which there are golden files
	pageIDsToTest := make([]int, len(goldenFileLocationIDs))
	i := 0
	for pageID := range goldenFileLocationIDs {
		pageIDsToTest[i] = pageID
		i++
	}
	PageIDs = pageIDsToTest

	Load()

	diffOutput, err := util.Diff(solrGoldenFilesDirectory, solrLoadTestTmpDirectory)
	if (err != nil) {
		t.Fatal("Diff of " + solrGoldenFilesDirectory + " and " +
			solrLoadTestTmpDirectory + " failed to run: " + err.Error())
	}

	if (diffOutput != "") {
		t.Errorf("%s", diffOutput)
	}
}
