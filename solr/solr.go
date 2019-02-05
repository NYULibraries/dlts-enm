package solr

import (
	"github.com/rtt/Go-Solr"
)

type solrConnection interface {
	Update(m map[string]interface{}, commit bool) (*solr.UpdateResponse, error)
}

var conn solrConnection

func Init(server string, port int, injectedConn solrConnection) error {
	if injectedConn != nil {
		conn = injectedConn
	} else {
		var err error
		conn, err = solr.Init(server, port, "enm-pages")
		if err != nil {
			return err
		}
	}

	return nil
}
