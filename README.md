# enm - DLTS Enhanced Network Monographs

CLI application for generating the Enhanced Networked Monographs (ENM) static website
and Solr index.

## Overview

`enm` is a CLI application for performing various backend ENM functions:

* Create static pages
  * About
  * Home
  * Create browse topics lists: e.g. [featured topics](http://enm.dlib.nyu.edu/browse-topics-lists/enm-picks.html)
  * Create topic pages: e.g. [culture -- popular](http://enm.dlib.nyu.edu/topic-pages/00/00/00/78/0000007805.html)
* Load `enm-pages` Solr index
* Automatically create cached data files which can be used in place of the Postgres database in
subsequent jobs: [nyudlts/enm\-cache](https://github.com/nyudlts/enm-cache)

The `enm` program has not been productionized.  Enough development was done to
create an initial stable and correct demo site.  ENM data is currently frozen and
there are no active plans to add or change data at this time.

## Getting Started

### Prerequisities

* [AWS CLI v2](https://aws.amazon.com/cli/) - for deployment scripts.
* [Go](https://golang.org/) - version 1.16 or higher

#### Database

The Postgres database from [TCT](https://github.com/NYULibraries/dlts-enm-tct-backend)
should already be set up on devdb1.dlib.nyu.edu with the correct user credentials.
See https://jira.nyu.edu/jira/browse/NYUP-437 for details.

### Installation and setup

```shell
git clone git@github.com:NYULibraries/dlts-enm.git
cd dlts-enm/
go build
mv dlts-enm enm
```

#### Set environment variables

Set environment variables for database access:

```shell
export ENM_POSTGRES_DATABASE=enm
export ENM_POSTGRES_DATABASE_HOSTNAME=127.0.0.1
export ENM_POSTGRES_DATABASE_USERNAME=enm_readonly
export ENM_POSTGRES_DATABASE_PASSWORD=[password for devdb1.dlib.nyu.edu:enm database for user enm_readonly]
```

Note that we use 127.0.0.1 even though the database is remote because we'll need
to access the remote Postgres server through an SSH tunnel through the bastion
host.

These environment variables must be set before running any `enm` command that
requires database access.  Failure to do so will cause a `panic`.

Set location of the cache using ENM_CACHE:

```shell script
export ENM_CACHE=$HOME/enm-cache/
```

This environment variable is optional.  If it is set and the path that is pointed
to does not already exist, `enm` will create it, along with any needed intermediate
directories.
If `ENM_CACHE` is not set, the location of the cache defaults to `/tmp/enm-cache/`.

#### Set up SSH tunnel to devdb1

Set up SSH tunneling to port 5432 on Postgres host devdb1.dlib.nyu.edu through
bastion host by running this command in a separate terminal:

```shell
ssh -N -L 5432:devdb1.dlib.nyu.edu:5432 [USERNAME]@b.dlib.nyu.edu
```

This will allow remote database access via local port 5432.

A less verbose command can be run if the following is set up in `.ssh/config`:

```shell
Host devdb1
     Hostname devdb1.dlib.nyu.edu
     ProxyCommand ssh bastion -W %h:%p
     User     [USERNAME]

...

Host bastion
     Hostname b.dlib.nyu.edu
     User     [USERNAME]
```

In a separate terminal, run this command:

```shell
ssh -N -L 5432:devdb1:5432 bastion
```

### Usage

#### Deploy site

There is a deploy script `bin/deploy-site.sh` that generates the full website,
syncs it with the S3 bucket (without touching the /search/ path that contains
the search application built by [dlts\-enm\-search\-application](https://github.com/NYULibraries/dlts-enm-search-application)),
and invalidates the website paths in CloudFront so that the latest files are fetched
from S3.

There is also a convenience wrapper script `bin/deploy-site_interactive.sh` which
will call `bin/deploy-site.sh` with options set according to the user's responses to
interactive prompts:

* `bin/deploy-site_interactive.sh dev`
* `bin/deploy-site_interactive.sh stage`
* `bin/deploy-site_interactive.sh prod`

See [examples](#examples) for demonstrations of deployment for other use cases.

The deploy script runs all the `sitegen` commands (detailed below) with destination
set to directories in `dist/`.

Note that the deploy script does not perform Solr indexing. 

#### Get general help

`./enm help`

#### Get help on a command

`./enm help sitegen`

#### Get help on a subcommmand

`./enm help sitegen browsetopicslists`

#### Create browse topics lists using Postgres database

`./enm sitegen browsetopicslists --destination=[DESTINATION]`

#### Create browse topics lists using cache files

`./enm sitegen browsetopicslists --destination=[DESTINATION] --source=cache`

#### Create page About, Home, etc.

`./enm sitegen sitepages --destination=[DESTINATION]`

#### Create all topic pages using Postgres database

This automatically creates cache files in /tmp/enm-cache/sitegen-topicpages/ that
can be used as the data source for subsequent topic pages generation runs:

`./enm sitegen topicpages --source=database --destination=[DESTINATION]`

#### Create all topic pages using cache files in /tmp/enm-cache/sitegen-topicpages/

`./enm sitegen topicpages --source=cache --destination=[DESTINATION]`

#### Create topic pages for 2 topics only using Postgres database

`./enm sitegen topicpages --source=database --destination=[DESTINATION] [TOPIC ID 1] [TOPIC ID 2]`

#### Create topic pages for 2 topics only using cache files

`./enm sitegen topicpages --source=cache --destination=[DESTINATION] [TOPIC ID 1] [TOPIC ID 2]`

#### Load enm-pages Solr index using Postgres database

`./enm solr load --server=[SOLR SERVER] --port 8983`

#### Load enm-pages Solr index using cache files

`./enm solr load --server=[SOLR SERVER] --source=cache --port 8983`

### Examples

#### Create the dev website from latest cache files cloned to $HOME/enm-cache/ and deploy 

In the example below, it is assumed that the `dlts-enm` repo is located at
`$GOPATH/src/github.com/nyulibraries/dlts-enm/`, and the `https://github.com/nyudlts/enm-cache`
repo has already been cloned to $HOME.

```shell
$ export ENM_CACHE=$HOME/enm-cache/
$ bin/deploy-site_interactive.sh dev
Do complete regeneration of the site before copying to server? [y/n] y
Use the cache for regenerating the site? [y/n] y
Generating site pages...
Generating browse topics lists...
Generating topic pages...
upload: dist/about.html to s3://dlts-enm-dev/about.html                            
upload: dist/browse-topics-lists/0-9.html to s3://dlts-enm-dev/browse-topics-lists/0-9.html
upload: dist/browse-topics-lists/enm-picks.html to s3://dlts-enm-dev/browse-topics-lists/enm-picks.html
upload: dist/browse-topics-lists/g.html to s3://dlts-enm-dev/browse-topics-lists/g.html 
upload: dist/browse-topics-lists/d.html to s3://dlts-enm-dev/browse-topics-lists/d.html
upload: dist/browse-topics-lists/f.html to s3://dlts-enm-dev/browse-topics-lists/f.html
upload: dist/browse-topics-lists/c.html to s3://dlts-enm-dev/browse-topics-lists/c.html
upload: dist/browse-topics-lists/j.html to s3://dlts-enm-dev/browse-topics-lists/j.html
upload: dist/browse-topics-lists/i.html to s3://dlts-enm-dev/browse-topics-lists/i.html

[...SNIPPED...]

upload: dist/topic-pages/00/00/04/74/0000047490.html to s3://dlts-enm-dev/topic-pages/00/00/04/74/0000047490.html
upload: dist/topic-pages/00/00/04/74/0000047476.html to s3://dlts-enm-dev/topic-pages/00/00/04/74/0000047476.html
upload: dist/topic-pages/00/00/04/74/0000047485.html to s3://dlts-enm-dev/topic-pages/00/00/04/74/0000047485.html
upload: dist/topic-pages/00/00/04/74/0000047491.html to s3://dlts-enm-dev/topic-pages/00/00/04/74/0000047491.html
upload: dist/topic-pages/00/00/04/74/0000047492.html to s3://dlts-enm-dev/topic-pages/00/00/04/74/0000047492.html
{
    "Location": "https://cloudfront.amazonaws.com/2020-05-31/distribution/E2DL5S1BQ4HW26/invalidation/I3IVXF6N96CJXR",
    "Invalidation": {
        "Id": "I3IVXF6N96CJXR",
        "Status": "InProgress",
        "CreateTime": "2021-07-27T22:57:49.269000+00:00",
        "InvalidationBatch": {
            "Paths": {
                "Quantity": 5,
                "Items": [
                    "/about.html",
                    "/index.html",
                    "/browse-topics-lists*",
                    "/shared*",
                    "/topic-pages*"
                ]
            },
            "CallerReference": "cli-1627426668-935235"
        }
    }
}
You have new mail in /var/mail/david
$ 
```

#### Load prod Solr index from Postgres database

`./enm solr load --server=discovery1.dlib.nyu.edu --port 8983`

#### Load prod Solr index from cache files at $ENM_CACHE if set, or default cache location `/tmp/enm-cache/` 

`./enm solr load --server=discovery1.dlib.nyu.edu --source=cache --port 8983`

## Running the tests

Make sure to set up access to the Postgres database before running the tests.
See [Set environment variables](#set-environment-variables).

```shell
go test ./...
```

## Generation of code files in `db/postgres/models`

The Go code in `db/postgres/models` was generated automatically by [xo](https://github.com/xo/xo)
using custom `xo` templates.  If changes are made to the Postgres database,
the models can be updated by running `go generate` at the root of the project.

## Configuration

Configuration is done through command/subcommand options and
[environment variables](https://github.com/NYULibraries/dlts-enm/tree/nyup-536_write-readme-documentation-for-dlts-enm-and-dlts-enm-search-application#set-environment-variables).
Environment variables starting with ENM_POSTGRES_DATABASE must be set before running any `enm` command that
requires database access.  Failure to do so will cause a `panic`.
See [Set environment variables](#set-environment-variables).

## Future improvements

* Real error handling/recovery/messaging/logging instead of the liberal use of
`panic` calls
* (maybe) Embedding of `sitegen` templates into the `enm` binary using something like
https://github.com/jteeuwen/go-bindata (for motivation see comment in
`sitegen/sitegen.go`)
* Write more tests, and stub out Postgres in the tests in the `solr` and `sitegen` packages
(and in all future tests).

## ENM project Github repos

* [dlts-enm](https://github.com/NYULibraries/dlts-enm)
* [dlts-enm-search-application](https://github.com/NYULibraries/dlts-enm-search-application)
* [dlts-enm-tct-backend](https://github.com/NYULibraries/dlts-enm-tct-backend)
* [dlts-enm-tct-developer](https://github.com/NYULibraries/dlts-enm-tct-developer)
* [dlts-enm-tct-frontend](https://github.com/NYULibraries/dlts-enm-tct-frontend)
* [dlts-enm-verifier](https://github.com/NYULibraries/dlts-enm-verifier)
* [dlts-enm-web](https://github.com/NYULibraries/dlts-enm-web)

## Built With

* [cobra](https://github.com/spf13/cobra)
* [Go-Solr](https://github.com/rtt/Go-Solr)
* [xo](https://github.com/xo/xo)

## License

This project is licensed under the Apache License Version 2.0 - see the [LICENSE](LICENSE)
file for details.
