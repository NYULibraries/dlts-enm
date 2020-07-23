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

* [Go](https://golang.org/) (at least 1.10 recommended)
* [dep](https://golang.github.io/dep/)
* [Expect](https://core.tcl.tk/expect/index) (if using the deploy shell scripts)

#### Database

The Postgres database from [TCT](https://github.com/NYULibraries/dlts-enm-tct-backend)
should already be set up on devdb1.dlib.nyu.edu with the correct user credentials.
See https://jira.nyu.edu/jira/browse/NYUP-437 for details.

### Installation and setup

Installation using `go get`:

```shell
go get github.com/nyulibraries/dlts-enm
cd dlts-enm/
git remote rm origin
git remote add origin git@github.com:NYULibraries/dlts-enm.git
git fetch --all
dep ensure
go build
mv dlts-enm enm
```

Installation using `git clone`:

```shell
mkdir -p $GOPATH/src/github.com/NYULibraries/
cd $GOPATH/src/github.com/NYULibraries/
git clone git@github.com:NYULibraries/dlts-enm.git
cd dlts-enm/
dep ensure
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

These environment variables must be set before running any `enm` command, even
the ones that do not technically need database access.  Failure to do so will
cause a `panic`.

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

There is a deploy script that can generate the full website and copy to dev, stage,
or prod web servers.  The easiest way to use it is to run the wrapper script
that prompts for whether the generate the full site and also prompts for the
server username and password, then generates the site (if requested) and performs
the many necessary `rsync` commands using [`expect`](https://core.tcl.tk/expect/index)
to automatically the enter user credentials for the bastion and web server hosts.

* `bin/deploy-site_expect.sh dev`
* `bin/deploy-site_expect.sh stage`
* `bin/deploy-site_expect.sh prod`

The deploy script runs all the `sitegen` commands detailed below with destination
set to directories in `dist/`.

See [examples](#examples) for a full demonstration of how to use this wrapper script.

Note that the deploy scripts do not perform Solr indexing, only static page
generation and copying to server.

#### Get general help

`./enm help`

#### Get help on a command

`./enm help sitegen`

#### Get help on a subcommmand

`./enm help sitegen browsetopicslists`

#### Create browse topics lists

`./enm sitegen browsetopicslists --destination=[DESTINATION]`

#### Create page About, Home, etc.

`./enm sitegen sitepages --destination=[DESTINATION]`

#### Create all topic pages using Postgres database

This automatically creates cache files in /tmp/enm-cache/sitegen-topicpages/ that
can be used as the data source for subsequent topic pages generation runs:

`./enm sitegen topicpages --source=database --destination=[DESTINATION]`

#### Create all topic pages using cache files in /tmp/enm-cache/sitegen-topicpages/

`./enm sitegen topicpages --source=cache --destination=[DESTINATION]`

#### Create topic pages for 2 topics only using Postgres database

`./enm sitegen topicpages --source=cache --destination=[DESTINATION] [TOPIC ID 1] [TOPIC ID 2]`

#### Create topic pages for 2 topics only using cache files

`./enm sitegen topicpages --source=cache --destination=[DESTINATION] [TOPIC ID 1] [TOPIC ID 2]`

#### Load enm-pages Solr index

Data source is always the Postgres database.
There is currently no caching of data for Solr index operations.

`./enm solr load --server=[SOLR SERVER] --port 8983`

### Examples

#### Create the production website and deploy to the web server

In the example below, it is assumed that the `dlts-enm` repo is located
`$GOPATH/src/github.com/nyulibraries/dlts-enm/`

```shell
$ bin/deploy-site_expect.sh dev
Do complete regeneration of the site before copying to server? [y/n] y
Username for b.dlib.nyu.edu and devweb1.dlib.nyu.edu: SOMEUSER
Password for b.dlib.nyu.edu and devweb1.dlib.nyu.edu:
spawn /Users/someuser/go/src/github.com/nyulibraries/dlts-enm/bin/deploy-site.sh -g -u SOMEUSER dev
Generating site pages...
Generating browse topics lists...
Generating topic pages...
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@b.dlib.nyu.edu's password:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@devweb1.dlib.nyu.edu's password:
building file list ... done
about.html

sent 127 bytes  received 66 bytes  77.20 bytes/sec
total size is 2.46K  speedup is 12.76

rsync #1 completed successfully.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@b.dlib.nyu.edu's password:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@devweb1.dlib.nyu.edu's password:
building file list ... done
index.html

sent 127 bytes  received 66 bytes  128.67 bytes/sec
total size is 2.45K  speedup is 12.69

rsync #2 completed successfully.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@b.dlib.nyu.edu's password:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@devweb1.dlib.nyu.edu's password:
building file list ... done
./
0-9.html
a.html
b.html
c.html
d.html
e.html
enm-picks.html
f.html
g.html
h.html
i.html
j.html
k.html
l.html
m.html
n.html
non-alphanumeric.html
o.html
p.html
q.html
r.html
s.html
t.html
u.html
v.html
w.html
x.html
y.html
z.html

sent 1.76K bytes  received 69.40K bytes  28.46K bytes/sec
total size is 8.47M  speedup is 119.01

rsync #3 completed successfully.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@b.dlib.nyu.edu's password:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@devweb1.dlib.nyu.edu's password:
building file list ... done

sent 2.50K bytes  received 20 bytes  1.68K bytes/sec
total size is 27.13M  speedup is 10761.88

rsync #4 completed successfully.
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@b.dlib.nyu.edu's password:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           WARNING:  UNAUTHORIZED PERSONS ........ DO NOT PROCEED
           ~~~~~~~   ~~~~~~~~~~~~~~~~~~~~          ~~~~~~~~~~~~~~

[...SNIPPED...]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
SOMEUSER@devweb1.dlib.nyu.edu's password:
building file list ... done
./
00/
00/00/
00/00/00/
00/00/00/00/
00/00/00/00/0000000002.html
00/00/00/00/0000000003.html
00/00/00/00/0000000004.html
00/00/00/00/0000000005.html

[...SNIPPED...]

sent 2.74M bytes  received 3.08M bytes  125.19K bytes/sec
total size is 222.68M  speedup is 38.25

rsync #5 completed successfully.
ENM site deployment completed.
```

#### Load prod Solr index

`./enm solr load --server=discovery1.dlib.nyu.edu --port 8983`

## Running the tests

Make sure access to the Postgres database has already been set up before running
the tests.  See [Set environment variables](#set-environment-variables).

```shell
go test ./...
```

Ideally for tests Postgres would be stubbed out with a fake.  This might be done
as a future improvement. 

## Generation of code files in `db/postgres/models`

The Go code in `db/postgres/models` was generated automatically by [xo](https://github.com/xo/xo)
using custom `xo` templates.  If changes are made to the the Postgres database,
the models can be updated by running `go generate` at the root of the project.

## Configuration

Configuration is done through command/subcommand options and
[environment variables](https://github.com/NYULibraries/dlts-enm/tree/nyup-536_write-readme-documentation-for-dlts-enm-and-dlts-enm-search-application#set-environment-variables),
which must be set before running any commands, including those that do not require
database access.

## Future improvements

* Real error handling/recovery/messaging/logging instead of the liberal use of
`panic` calls
* `solr load` caching of Postgres data, for improved performance and for decoupling
from the TCT Postgres database (and TCT in general)
* (maybe) Embedding of `sitegen` templates into the `enm` binary using something like
https://github.com/jteeuwen/go-bindata (for motivation see comment in
`sitegen/sitegen.go`)
* Write more tests, and stub out Postgres in tests in the `solr` and `sitegen` packages
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
