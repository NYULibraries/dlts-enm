package solr

import (
	"github.com/rtt/Go-Solr"
)

var Port int
var Server string
var conn *solr.Connection

func Init(server string, port int) error {
	var err error
	conn, err = solr.Init(server, port, "enm-pages")
	if err != nil {
		return err
	}

	return nil
}
