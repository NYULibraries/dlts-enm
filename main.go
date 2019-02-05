package main

import (
	"fmt"
	"os"

	"github.com/nyulibraries/dlts-enm/cmd"
)

//go:generate bin/xo postgres://${ENM_POSTGRES_DATABASE_USERNAME}:${ENM_POSTGRES_DATABASE_PASSWORD}@localhost/${ENM_POSTGRES_DATABASE}?sslmode=disable -o db/postgres/models/ --template-path $PWD/xo/templates/

//go:generate xo/scripts/postgres/generate-models-for-topic-page-queries.sh

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
