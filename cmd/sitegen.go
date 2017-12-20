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
	"os"

	"github.com/spf13/cobra"
)

var Destination string

// sitegenCmd represents the sitegen command
var sitegenCmd = &cobra.Command{
	Use:   "sitegen",
	Short: "Generate ENM website components",
	Long: `Generate the following ENM website components:

	* Topics browse lists
	* Topic pages
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("sitegen called")
	},
}

func init() {
	RootCmd.AddCommand(sitegenCmd)

	defaultDestination, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	sitegenCmd.PersistentFlags().StringVarP(
		&Destination,
		"destination",
		"d",
		defaultDestination,
		"Filesystem path to write files to",
	)
}
