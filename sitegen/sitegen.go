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

package sitegen

import (
	"html/template"
	"os"
)

type ExternalRelation struct{
	Relationship string
	URL string
	Vocabulary string
}

type ExternalRelations struct{
	AlternateNames []string
	DisplayName string
	EPUBMatches []string
	ExternalRelations []ExternalRelation
}

type Paths struct{
	WebRoot string
}

func Test() {
	funcs := template.FuncMap{
		"lastIndex": func (s []interface{}) int {
			return len(s) - 1;
		},
	}

	var tpl, err = template.New("index.html").Funcs(funcs).ParseFiles("templates/topic-page/index.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, ExternalRelations{
		AlternateNames: []string{"alt1", "alt2", "alt3"},
		DisplayName: "topic!",
		EPUBMatches: []string{"epub1", "epub2", "epub3"},
		ExternalRelations: []ExternalRelation{
			{
				Relationship: "exactMatch",
				URL: "http://nowhere.net",
				Vocabulary: "VIAF",
			},
			{
				Relationship: "exactMatch",
				URL: "http://nowhere2.net",
				Vocabulary: "VIAF2",
			},
			{
				Relationship: "exactMatch",
				URL: "http://nowhere2.net",
				Vocabulary: "VIAF2",
			},
		},
		Paths: Paths{
			WebRoot: "webroot",
		},
	})
	if err != nil {
		panic(err)
	}
}