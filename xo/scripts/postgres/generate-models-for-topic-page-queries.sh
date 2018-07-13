#!/usr/bin/env bash

function xo_generate() {
    local query="$( cat ${1} )"
    local queryType=$2

    bin/xo postgres://${ENM_POSTGRES_DATABASE_USERNAME}:${ENM_POSTGRES_DATABASE_PASSWORD}@localhost/${ENM_POSTGRES_DATABASE}?sslmode=disable \
        --query-mode \
        --query-trim \
	    --query-strip \
        --query "${query}" \
        --query-type $queryType \
        --out db/postgres/models
}

xo_generate xo/queries/postgres/epubs-for-topic-with-number-of-matched-pages.sql EpubsForTopicWithNumberOfMatchedPages
xo_generate xo/queries/postgres/epubs-number-of-pages.sql EpubsNumberOfPages
xo_generate xo/queries/postgres/external-relations-for-topic.sql ExternalRelationsForTopic
xo_generate xo/queries/postgres/page.sql Page
xo_generate xo/queries/postgres/related-topic-names-for-topic-with-number-of-occurrences.sql RelatedTopicNamesForTopicWithNumberOfOccurrences
xo_generate xo/queries/postgres/topic-alternate-name.sql TopicAlternateName
xo_generate xo/queries/postgres/topic-names-for-page.sql TopicNamesForPage
xo_generate xo/queries/postgres/topic-number-of-occurrences.sql TopicNumberOfOccurrences
