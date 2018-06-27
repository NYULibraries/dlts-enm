package solr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/rtt/Go-Solr"

	"github.com/nyulibraries/dlts-enm/util"
)

type solrStub struct {
	URL     string
	Version []int
}

var solrGoldenFilesDirectory string
var solrLoadTestTmpDirectory string

func (s solrStub) Update(m map[string]interface{}, commit bool) (*solr.UpdateResponse, error) {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return nil, err
	}

	locationId := m[ "add" ].([]interface{})[0].(interface{}).(map[string]interface{})["id"].(int)
	err = ioutil.WriteFile(solrLoadTestTmpDirectory+ "/" + strconv.Itoa(locationId) + ".json", b, 0644)
	if err != nil {
		panic(err)
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

func TestLoad(t *testing.T) {
	solrStub, err := newSolrStub("fakesolr.dlib.nyu.edu", 8983, "enm-pages")
	if err != nil {
		t.Fatal(err)
	}

	Init("", 0,  solrStub)

	Load()
}