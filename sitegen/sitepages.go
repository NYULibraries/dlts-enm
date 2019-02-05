package sitegen

import (
	"html/template"
	"os"
)

var sitePagesDir string

func GenerateSitePages(destination string) {
	sitePagesDir = destination + "/"

	if _, err := os.Stat(sitePagesDir); os.IsNotExist(err) {
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
		SharedTemplateDirectory    + "/google-analytics.html",
	)

	if err != nil {
		return err
	}

	file := sitePagesDir + "/" + filename
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	pageData := struct{
		GoogleAnalytics bool
		Paths
	}{
		GoogleAnalytics: GoogleAnalytics,
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
