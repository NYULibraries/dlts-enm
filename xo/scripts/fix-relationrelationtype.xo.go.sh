#!/usr/bin/env bash

# Remove duplicate Rtype field declaration from relationrelationtype.xo.go
# https://jira.nyu.edu/jira/browse/NYUP-397?focusedCommentId=94979&page=com.atlassian.jira.plugin.system.issuetabpanels:comment-tabpanel#comment-94979
# Sed command from https://stackoverflow.com/questions/45147094/sed-replace-multiline
# Added "// Duplicate Rtype field declaration removed" replacement text because
# otherwise an empty line is left in the file, which looks weird.
# Note that there doesn't seem to be a way to inline sed without backup file
# that works on both Mac OS X (BSD sed) and Linux (GNU sed), so just make a
# backup file and delete it.
sed -i.bak '/^[[:space:]]*Rtype       string         `json:"rtype"`       \/\/ rtype/{n;s/^[[:space:]]*Rtype       string         `json:"rtype"`       \/\/ rtype/\/\/ Duplicate Rtype field declaration removed/;}' \
    db/postgres/models/relationrelationtype.xo.go

rm db/postgres/models/relationrelationtype.xo.go.bak
