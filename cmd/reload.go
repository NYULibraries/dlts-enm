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

	_ "github.com/nyulibraries/dlts/enm/db"
	"github.com/nyulibraries/dlts/enm/db"
	"github.com/nyulibraries/dlts/enm/tct"
)

// reloadCmd represents the reload command
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload database from TCT",
	Long: `Truncates all database tables and reload from TCT API responses.`,
	Run: func(cmd *cobra.Command, args []string) {
		tct.Source = Source

		db.ClearTables()
		db.Reload()
	},
}

func init() {
	RootCmd.AddCommand(reloadCmd)

	reloadCmd.PersistentFlags().StringVarP(
		&Source,
		"source",
		"s",
		"tct-api",
		"Specify data source: tct-api, cache",
	)
}
