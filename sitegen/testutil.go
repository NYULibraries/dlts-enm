package sitegen

const (
	noGoogleAnalyticsSubdirectory   = "no-google-analytics"
	withGoogleAnalyticsSubdirectory = "with-google-analytics"
)

func getGoldenFileSubdirectory() (string) {
	if ( GoogleAnalytics ) {
		return withGoogleAnalyticsSubdirectory
	} else {
		return noGoogleAnalyticsSubdirectory
	}
}