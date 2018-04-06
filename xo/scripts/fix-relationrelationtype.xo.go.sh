#!/usr/bin/env bash

# Remove duplicate Rtype field declaration from relationrelationtype.xo.go
# https://jira.nyu.edu/jira/browse/NYUP-397?focusedCommentId=94979&page=com.atlassian.jira.plugin.system.issuetabpanels:comment-tabpanel#comment-94979
# Sed command from https://stackoverflow.com/questions/45147094/sed-replace-multiline
# Added "// Duplicate Rtype field declaration removed" replacement text because
# otherwise an empty line is left in the file, which looks weird.
sed -i '' '/^[[:space:]]*Rtype       string         `json:"rtype"`       \/\/ rtype/{n;s/^[[:space:]]*Rtype       string         `json:"rtype"`       \/\/ rtype/\/\/ Duplicate Rtype field declaration removed/;}' \
    db/postgres/models/relationrelationtype.xo.go
