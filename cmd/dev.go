package cmd
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

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Source string

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Temporary command for initial development",
	Long: `Just a temporary command to use during initial development for
	running new code and doing sundry experiments`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev command called")
	},
}

func init() {
	RootCmd.AddCommand(devCmd)

	devCmd.PersistentFlags().StringVarP(
		&Source,
		"source",
		"s",
		"tct-api",
		"Specify data source: tct-api, cache",
	)
}
