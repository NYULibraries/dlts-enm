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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/nyulibraries/dlts-enm/solr"
)

// solrLoadCmd represents the add command
var solrLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads the Solr index",
	Long: `Loads the Solr index from views pages and page_topics`,
	Run: func(cmd *cobra.Command, args []string) {
		err := solr.Init(Server, Port)
		if err != nil {
			panic(fmt.Sprintf("ERROR: couldn't initialize solr: %s\n", err.Error()))
		}

		solr.Load()
	},
}

func init() {
	solrCmd.AddCommand(solrLoadCmd)
}
