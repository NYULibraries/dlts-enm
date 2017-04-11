package main

import (
	"fmt"
	"os"

	"github.com/nyulibraries/dlts/enm/cmd"
	_ "github.com/nyulibraries/dlts/enm/db"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}