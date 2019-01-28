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

package sitegen

type Paths struct{
	WebRoot string
}

var GoogleAnalytics bool
var Source string

// Tricky...this assumes that location of the templates relative to working directory
// matches what the location would be if calling `go run main.go` from root of this
// repo.  This is brittle, breaking in multiple cases:
//
// * Doing `go run [PATH TO main.go` or calling the binary from somewhere relative
//   to the templates directory other than what's assumed below.
// * `go test` on tests in a directory that would set working directory to some
//   place other than what's assumed below.
//
// Consider including the templates in the binary using something like
// https://github.com/jteeuwen/go-bindata
var BrowseTopicListsTemplateDirectory = "sitegen/templates/browse-topics-lists"
var SharedTemplateDirectory = "sitegen/templates/shared"
var SitePagesTemplateDirectory = "sitegen/templates"
var TopicPageTemplateDirectory =  "sitegen/templates/topic-page"
