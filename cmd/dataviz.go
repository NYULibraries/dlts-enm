// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var OutputDir string

// datavizCmd represents the dataviz command
var datavizCmd = &cobra.Command{
	Use:   "dataviz",
	Short: "Create data visualizations",
	Long: `Create data visualizations.  Gallery:

* trees1: collapsible topic trees - see NYUP-234`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dataviz with outputdir = " + OutputDir)
	},
}

func init() {
	RootCmd.AddCommand(datavizCmd)

	wd, err := os.Getwd()
	if err != nil {
		panic("Can't get working directory")
	}

	datavizCmd.PersistentFlags().StringVarP(
		&OutputDir,
		"outputdir",
		"o",
		wd,
		"Specify output directory",
	)
}
