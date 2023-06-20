#!/bin/bash

# By default, we're assuming you created the extended k6 image as "k6-extended:latest".
# If not, override the name on the command-line with `IMAGE_NAME=...`.

# Each execution is provided a unique `testid` tag to differentiate discrete test runs.
# (Not required, but provided for convenience)

# This is a basic wrapper to run a clean docker container
#   -v   : we're allowing scripts to be located in the current directory, or any of its children
#   --rm : once the script completes, the container will be removed (good housekeeping, you'll thank me)
for SCRIPT_NAME in "$(pwd)"/scripts/*
do
    if [ $(find . -type f | wc -l) -eq 0 ] ; then
        echo "scripts directory do not contain any scripts."
        exit 0
    fi
    SCRIPT_BASE_NAME=$(basename -s .js $SCRIPT_NAME)
    TAG_NAME="$SCRIPT_BASE_NAME-$(date +%F_%T)"
    docker-compose run --rm -v /$PWD/scripts:/scripts k6 run //scripts/$SCRIPT_BASE_NAME.js --tag testid=$TAG_NAME
done
