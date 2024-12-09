module github.com/nyulibraries/dlts-enm

go 1.16

replace github.com/nyulibraries/dlts-enm/cmd => ./cmd

require (
	github.com/lib/pq v1.10.4
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/rtt/Go-Solr v0.0.0-20190512221613-64fac99dcae2
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.7
	gopkg.in/ini.v1 v1.66.4 // indirect
)
