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
