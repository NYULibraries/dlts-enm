#!/usr/bin/env bash

ROOT=$(cd "$(dirname "$0")" ; cd ..; pwd -P )
DIST=$ROOT/dist

source $ROOT/bin/$(basename $0 .sh)_common.sh

function usage() {
    script_name=$(basename $0)

    cat <<EOF

usage: ${script_name} [-g] [-u username] environment
    -g:          generate static in dist/
    -u username: username on bastion host and web server
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

    if [ ! -x "$(command -v rsync)" ]
    then
        echo >&2 '`rsync` must be available in $PATH in order to run this script.'
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
    go run main.go sitegen sitepages --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi

    echo 'Generating browse topics lists...'
    go run main.go sitegen browsetopicslists --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi

    echo 'Generating topic pages...'
    go run main.go sitegen topicpages --destination=$dist $google_analytics_flag
    if [ $? -ne 0 ]
    then
        exit 1
    fi
}

function copy_files() {
    local username=$1
    local bastion_host=$2
    local server=$3
    local static_site_path=$4

    cp -pr $ROOT/sitegen/static $DIST/shared

    for f in about.html index.html browse-topics-lists/ shared/ topic-pages/
    do
        rsync --archive --compress --delete --human-readable --verbose \
                -e "ssh -o ProxyCommand='ssh -W %h:%p ${username}@${bastion_host}'" \
                $ROOT/dist/$f \
                ${username}@${server}:${static_site_path}/$f
    done
}

check_prerequisites

STATIC_SITE_PATH=/www/sites/enm

generate_site=false

while getopts gu: opt
do
    case $opt in
        g) generate_site=true ;;
        u) username=$OPTARG ;;
        *) echo >&2 "Options not set correctly."; usage; exit 1 ;;
    esac
done

if [ -z $username ]; then
    echo >&2 'You must provide a username.'
    usage
    exit 1
fi

shift $((OPTIND-1))

deploy_to_environment=$1

validate_environment_arg $deploy_to_environment

google_analytics_flag=$( get_google_analytics_flag $deploy_to_environment )

server=$( get_server $deploy_to_environment )

if $generate_site
then
    clean_dist $DIST
    generate_site $DIST $google_analytics_flag
fi

copy_files $username $BASTION_HOST $server $STATIC_SITE_PATH

# This string tells the expect script wrapper that refresh run has completed.
echo $SCRIPT_RUN_COMPLETE
