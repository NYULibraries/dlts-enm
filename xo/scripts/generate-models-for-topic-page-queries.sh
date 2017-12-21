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
        --out db/models/

}

xo_generate xo/queries/topics-alternate-names.sql TopicAlternateName
xo_generate xo/queries/epubs-for-topic-with-number-of-matched-pages.sql EpubsForTopicWithNumberOfMatchedPages
xo_generate xo/queries/related-topic-names-for-topic-with-number-of-occurrences.sql RelatedTopicNamesForTopicWithNumberOfOccurrences
xo_generate xo/queries/external-relations-for-topic.sql ExternalRelationsForTopic