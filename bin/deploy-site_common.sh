MINIMUM_BASH_VERSION=4

BASTION_HOST=b.dlib.nyu.edu

declare -A ENVIRONMENTS

GA_KEY_SUFFIX='-google-analytics'

ENVIRONMENTS[dev]=devweb1.dlib.nyu.edu
ENVIRONMENTS[dev${GA_KEY_SUFFIX}]=

ENVIRONMENTS[stage]=stageweb1.dlib.nyu.edu
ENVIRONMENTS[stage${GA_KEY_SUFFIX}]=

ENVIRONMENTS[prod]=web1.dlib.nyu.edu
ENVIRONMENTS[prod${GA_KEY_SUFFIX}]='--google-analytics'

# This string tells the expect script wrapper that refresh run has completed.
SCRIPT_RUN_COMPLETE='ENM site deployment completed.'

function validate_environment_arg() {
    local deploy_to_environment=$1

    if ! [ ${ENVIRONMENTS[${deploy_to_environment}]} ]
    then
        echo >&2 "\"${deploy_to_environment}\" is not a recognized deployment environment."

        usage

        exit 1
    fi
}

function get_google_analytics_flag() {
    local deploy_to_environment=$1

    echo ${ENVIRONMENTS[${deploy_to_environment}${GA_KEY_SUFFIX}]}
}

function get_server() {
    local deploy_to_environment=$1

    echo ${ENVIRONMENTS[${deploy_to_environment}]}
}

