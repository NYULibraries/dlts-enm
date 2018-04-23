// Copyright © 2017 NYU
//
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
	"github.com/spf13/cobra"

	"fmt"
)

var ReloadSource string

// reloadCmd represents the reload command
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload database from TCT",
	Long: `Truncates all database tables and reload from TCT API responses.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Disabled.  See comments below.  Keeping around for now until we are 100%
		// we'll never need to extract from the TCT JSON API again and will never
		// revive the original prototype MySQL ENM database.
		fmt.Println("`reload` is obsolete.  The last (probably) " +
			"working version was https://github.com/NYULibraries/dlts-enm/releases/tag/final-commit-before-moving-to-postgres.")
		fmt.Println("For more details, see https://jira.nyu.edu/jira/browse/NYUP-413.")

		// This is what this command used to do:
		// tct.Source = ReloadSource
        //
		// mysql.ClearTables()
		// mysql.Reload()
	},
}

func init() {
	RootCmd.AddCommand(reloadCmd)

	reloadCmd.PersistentFlags().StringVarP(
		&ReloadSource,
		"source",
		"s",
		"tct-api",
		"Specify data source: tct-api, cache",
	)
}
