package solr

import (
	"encoding/json"
	"fmt"
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

type solrStub struct {
	URL     string
	Version []int
}

var goldenFileLocationIDs map[int]bool
var solrGoldenFilesDirectory string
var solrLoadTestTmpDirectory string

func (s solrStub) Update(m map[string]interface{}, commit bool) (*solr.UpdateResponse, error) {
	locationID := m[ "add" ].([]interface{})[0].(interface{}).(map[string]interface{})["id"].(int)

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

// Based on Go Solr solr.Init()
func newSolrStub(host string, port int, core string) (*solrStub, error) {

	if len(host) == 0 {
		return nil, fmt.Errorf("Invalid hostname (must be length >= 1)")
	}

	if port <= 0 || port > 65535 {
		return nil, fmt.Errorf("Invalid port (must be 1..65535)")
	}

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

func TestLoad(t *testing.T) {
	solrStub, err := newSolrStub("fakesolr.dlib.nyu.edu", 8983, "enm-pages")
	if err != nil {
		t.Fatal(err)
	}

	Init("", 0,  solrStub)

	Load()
}