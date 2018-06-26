package solr

import (
	"fmt"

	"github.com/rtt/Go-Solr"
)

type solrStub struct {
	URL     string
	Version []int
}

func (s solrStub) Update(m map[string]interface{}, commit bool) (*solr.UpdateResponse, error) {
	fmt.Printf("%v", m)

	return &solr.UpdateResponse{ Success: true }, nil
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
