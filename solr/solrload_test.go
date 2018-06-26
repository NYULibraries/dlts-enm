package solr

import (
	"testing"
)

func TestLoad(t *testing.T) {
	solrStub, err := newSolrStub("fakesolr.dlib.nyu.edu", 8983, "enm-pages")
	if err != nil {
		t.Fatal(err)
	}

	Init("", 0,  solrStub)

	Load()
}