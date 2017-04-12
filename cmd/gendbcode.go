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

	"github.com/nyulibraries/dlts/enm/db"
)

// gendbcodeCmd represents the gendbcode command
var gendbcodeCmd = &cobra.Command{
	Use:   "gendbcode",
	Short: "Generate go code from table definitions",
	Long: `Generate tables-related go code for db package`,
	Run: func(cmd *cobra.Command, args []string) {
		db.GenerateDbCode()
	},
}

func init() {
	devCmd.AddCommand(gendbcodeCmd)
}
