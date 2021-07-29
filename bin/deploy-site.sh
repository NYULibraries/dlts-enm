#!/usr/bin/env bash

ROOT=$( cd "$(dirname "$0")" ; cd ..; pwd -P )
DIST=$ROOT/dist

MINIMUM_BASH_VERSION=4

declare -A ENVIRONMENT

CLOUDFRONT_DISTRIBUTION_ID_KEY_SUFFIX='-cloudfront-distribution-id'
GA_KEY_SUFFIX='-google-analytics'
S3_BUCKET_KEY_SUFFIX='-s3-bucket'

ENVIRONMENT[dev${GA_KEY_SUFFIX}]=
ENVIRONMENT[dev${CLOUDFRONT_DISTRIBUTION_ID_KEY_SUFFIX}]=E2DL5S1BQ4HW26
ENVIRONMENT[dev${S3_BUCKET_KEY_SUFFIX}]=dlts-enm-dev

ENVIRONMENT[stage${GA_KEY_SUFFIX}]=
ENVIRONMENT[stage${CLOUDFRONT_DISTRIBUTION_ID_KEY_SUFFIX}]=E1D91RKPQMQWM8
ENVIRONMENT[stage${S3_BUCKET_KEY_SUFFIX}]=dlts-enm-stage

ENVIRONMENT[prod${GA_KEY_SUFFIX}]='--google-analytics'
ENVIRONMENT[prod${CLOUDFRONT_DISTRIBUTION_ID_KEY_SUFFIX}]=E1TWMHHWMKVLXN
ENVIRONMENT[prod${S3_BUCKET_KEY_SUFFIX}]=dlts-enm

function usage() {
    script_name=$(basename $0)

    cat <<EOF

usage: ${script_name} [-c] [-g] [-h] environment
    -c:          generate static site from cache instead of database
    -g:          generate static site in dist/
    -h:          print this usage message
    environment: "dev", "stage", or "prod"

EOF
}

function check_prerequisites() {
    if [ -z "${BASH_VERSINFO}" ] || [ -z "${BASH_VERSINFO[0]}" ] || [ ${BASH_VERSINFO[0]} -lt $MINIMUM_BASH_VERSION ]
    then
        echo >&2 "Bash version must be >= $MINIMUM_BASH_VERSION"
        exit 1
    fi

    if [ ! -x "$(command -v go)" ]
    then
        echo >&2 '`go` must be available in $PATH in order to run this script.'
        exit 1
    fi

    if [ ! -x "$(command -v aws)" ]
    then
        echo >&2 '`aws` must be available in $PATH in order to run this script.'
        exit 1
    fi
}

function clean_dist() {
    local dist=$1
    rm -fr $dist/*
}

function generate_site() {
    local dist=$1
    local google_analytics_flag=$2

    echo 'Generating site pages...'
    go run main.go sitegen sitepages --source=$source --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi

    echo 'Generating browse topics lists...'
    go run main.go sitegen browsetopicslists --source=$source --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi

    echo 'Generating topic pages...'
    go run main.go sitegen topicpages --source=$source --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi
}

function get_cloudfront_distribution_id() {
    local deploy_to_environment=$1

    echo ${ENVIRONMENT[${deploy_to_environment}${CLOUDFRONT_DISTRIBUTION_ID_KEY_SUFFIX}]}
}

function get_google_analytics_flag() {
    local deploy_to_environment=$1

    echo ${ENVIRONMENT[${deploy_to_environment}${GA_KEY_SUFFIX}]}
}

function get_s3_bucket() {
    local deploy_to_environment=$1

    echo ${ENVIRONMENT[${deploy_to_environment}${S3_BUCKET_KEY_SUFFIX}]}
}

function invalidate_cloudfront_paths() {
    local cloudfront_distribution_id=$1

    # https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Invalidation.html#invalidation-specifying-objects
    aws cloudfront create-invalidation \
        --distribution-id ${cloudfront_distribution_id} \
        --paths '/' '/about.html' '/browse-topics-lists*' '/index.html' '/shared*' '/topic-pages*'
}

function sync_s3_bucket() {
    local s3_bucket=$1

    aws s3 sync $DIST s3://${s3_bucket} \
        --delete \
        --exact-timestamps \
        --exclude '.commit-empty-directory' \
        --exclude 'search/*'
}

function validate_environment_arg() {
    local deploy_to_environment=$1

    if ! [ ${ENVIRONMENT[${deploy_to_environment}${S3_BUCKET_KEY_SUFFIX}]} ]
    then
        echo >&2 "\"${deploy_to_environment}\" is not a recognized deployment environment."

        usage

        exit 1
    fi
}

check_prerequisites

source='database'
generate_site=false

while getopts cgh: opt
do
    case $opt in
        c) source='cache' ;;
        g) generate_site=true ;;
        h) usage; exit 0 ;;
        *) echo >&2 "Options not set correctly."; usage; exit 1 ;;
    esac
done

shift $((OPTIND-1))

deploy_to_environment=$1

validate_environment_arg $deploy_to_environment

google_analytics_flag=$( get_google_analytics_flag $deploy_to_environment )

if $generate_site
then
    clean_dist $DIST

    cp -pr $ROOT/sitegen/static $DIST/shared

    generate_site $DIST $google_analytics_flag
fi

s3_bucket=$( get_s3_bucket $deploy_to_environment )
sync_s3_bucket $s3_bucket

cloudfront_distribution_id=$( get_cloudfront_distribution_id $deploy_to_environment )
invalidate_cloudfront_paths $cloudfront_distribution_id
