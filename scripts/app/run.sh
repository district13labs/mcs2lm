#!/usr/bin/env bash
source ${PWD}/.env

COMMAND=$1
case $COMMAND in
    "mcli")
        docker run \
        -u $(id -u):$(id -g) \
        -ti \
        -v ${VOLUME_NAME}:/${BIN_PATH} \
        -v ${PWD}:${APP_PATH} \
        --rm ${DOCKER_IMAGE_NAME} \
        ${BIN_PATH}/${BIN_NAME} "${@:2}"
        ;;

    "test")
        docker run \
        -u $(id -u):$(id -g) \
        -ti \
        -v ${VOLUME_NAME}:/${BIN_PATH} \
        -v ${PWD}:${APP_PATH} \
        --rm ${DOCKER_IMAGE_NAME} \
        go test "${@:2}"
        ;;

    "shell")
        docker run \
        -u $(id -u):$(id -g) \
        -ti \
        -v ${VOLUME_NAME}:/${BIN_PATH} \
        -v ${PWD}:${APP_PATH} \
        --rm ${DOCKER_IMAGE_NAME} \
        ${SHELL}
        ;;
        
    *)
        echo "
The purpose of this script is to make the development and manual testing easier.
You can either pass arguments to the previously built application, run tests and 
go inside the container opening a shell.

        
    Usage:

        Arguments:

            mcli     Pass arguments to the built application.
            test     Run tests. Basically prefixes your arguments with \"go test\" .
            shell    Opens a shell.


        Examples:

        1) Run all test in verbose mode
        
        ./scripts/app/run.sh test ./... -v


        2) Run top-level tests matching \"Foo\", such as \"TestFooBar\":
        
        ./scripts/app/run.sh test -run Foo 


        3) Pass arguments to the mcli

        ./scripts/app/run.sh cli arg1 arg2


        4) Open a shell

        ./scripts/app/run.sh shell
        "
esac

