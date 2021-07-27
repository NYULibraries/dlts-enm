#!/usr/bin/env bash

ROOT=$(cd "$(dirname "$0")" ; cd ..; pwd -P )

DEPLOY_SCRIPT=${ROOT}/bin/$(basename $0 _interactive.sh).sh

if [ ! -x "$(command -v aws)" ]
then
    cat >&2 <<EOF

The [AWS Command Line Interface](https://aws.amazon.com/cli/) does not appear to
be installed.  You will need to install it and put the executable `aws` in your PATH
in order to run this script.
EOF

    exit 1
fi

function usage() {
    script_name=$(basename $0)

    cat <<EOF

usage: ${script_name} environment
    environment: "dev", "stage", or "prod"

EOF
}

while getopts h opt
do
    case $opt in
        h) usage; exit 0 ;;
        *) echo >&2 "Options not set correctly."; usage; exit 1 ;;
    esac
done

shift $((OPTIND-1))

deploy_to_environment=$1

if [ -z "$deploy_to_environment" ]
then
    usage

    exit 0
fi

echo -n 'Do complete regeneration of the site before copying to server? [y/n] '
read generate_site

case "$generate_site" in
    [yY][eE][sS]|[yY])
        generate_site_flag=-g
        ;;
    *)
        generate_site_flag=
        ;;
esac

echo -n 'Use the cache for regenerating the site? [y/n] '
read use_cache

case "$use_cache" in
    [yY][eE][sS]|[yY])
        use_cache_flag=-c
        ;;
    *)
        use_cache_flag=
        ;;
esac

${DEPLOY_SCRIPT} ${use_cache_flag} ${generate_site_flag} ${deploy_to_environment}
