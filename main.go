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

package main

import (
	"fmt"
	"os"

	"github.com/nyulibraries/dlts-enm/cmd"
)

//go:generate bin/xo postgres://${ENM_POSTGRES_DATABASE_USERNAME}:${ENM_POSTGRES_DATABASE_PASSWORD}@localhost/${ENM_POSTGRES_DATABASE}?sslmode=disable -o db/postgres/models/ --template-path $PWD/xo/templates/

//go:generate bin/xo postgres://${ENM_POSTGRES_DATABASE_USERNAME}:${ENM_POSTGRES_DATABASE_PASSWORD}@localhost/${ENM_POSTGRES_DATABASE}?sslmode=disable --query-mode --query-trim --query-strip --query "SELECT hb.id, hb.display_name, hh.name FROM hit_basket hb INNER JOIN hit_hit hh ON hb.id = hh.basket_id ORDER BY display_name, name" --query-type TopicAlternateNames --out db/postgres/models

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}