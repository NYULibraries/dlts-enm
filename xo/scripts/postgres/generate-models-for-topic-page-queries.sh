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

xo_generate xo/queries/postgres/related-topic-names-for-topic-with-number-of-occurrences.sql RelatedTopicNamesForTopicWithNumberOfOccurrences
