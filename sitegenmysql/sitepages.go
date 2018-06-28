// Copyright © 2018 NYU
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

package sitegenmsyql

import (
	"html/template"
	"os"
)

var SitePagesDir string

func GenerateSitePages(destination string) {
	SitePagesDir = destination + "/"

	if _, err := os.Stat(SitePagesDir); os.IsNotExist(err) {
		panic(err)
	}

	WriteSitePage("about")
	WriteSitePage("index")
}

func WriteSitePage(pageName string) (err error){
	var filename = pageName + ".html"
	var templateFile = pageName + ".html"

	tpl := template.New(filename)
	tpl, err = tpl.ParseFiles(
		SitePagesTemplateDirectory + "/" + templateFile,
		SharedTemplateDirectory    + "/banner.html",
	)

	if err != nil {
		return err
	}

	file := SitePagesDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	pageData := struct{
		Paths
	}{
		Paths: Paths{
			WebRoot: ".",
		},
	}

	err = tpl.Execute(f, pageData)
	if err != nil {
		return err
	}

	return nil
}
