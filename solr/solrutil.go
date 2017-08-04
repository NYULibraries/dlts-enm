package solrutil

import (
	"errors"
	"fmt"

	"github.com/nyulibraries/dlts-enm/db/models"
	solr "github.com/rtt/Go-Solr"
)

var Port int
var Server string

func Init(server string, port int) error {
	s, err := solr.Init(server, port, "enm-topics")
	if err != nil {
		return err
	}

	fmt.Printf("solr.Init(): %v\n", s)

	return nil
}

func AddPage(page *models.Page) error {
	return errors.New("Error!")
}