#!/usr/bin/env bash

function xo_generate() {
    local query="$( cat ${1} )"
    local queryType=$2

    xo mysql://${ENM_DATABASE_USERNAME}:${ENM_DATABASE_PASSWORD}@localhost/${ENM_DATABASE} \
        --query-mode \
        --query-trim \
        --query-strip \
        --query "${query}" \
        --query-type $queryType \
        --out db/mysql/models/

}

xo_generate xo/queries/topic-number-of-occurrences.sql TopicNumberOfOccurrences
xo_generate xo/queries/topics-alternate-names.sql TopicAlternateName
xo_generate xo/queries/epubs-for-topic-with-number-of-matched-pages.sql EpubsForTopicWithNumberOfMatchedPages
# NOTE: This model was originally created by this script, but in order to fix a
# bug which caused zero-occurrence topics to be omitted, the SQL file had to be
# updated to UNION the omitted topics.
# Due to a bug in mysql versions below 8.0.0, `xo` is unable to generate mysql
# models based on UNION queries.
# For details, see:
#     https://jira.nyu.edu/jira/browse/NYUP-414?focusedCommentId=99202&page=com.atlassian.jira.plugin.system.issuetabpanels:comment-tabpanel#comment-99202
#xo_generate xo/queries/related-topic-names-for-topic-with-number-of-occurrences.sql RelatedTopicNamesForTopicWithNumberOfOccurrences
xo_generate xo/queries/external-relations-for-topic.sql ExternalRelationsForTopic