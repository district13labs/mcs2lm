.PHONY: setup-githooks setup test

APP_PATH=/app
BIN_PATH=${APP_PATH}/bin
BIN_NAME=mcs2lm
DOCKER_IMAGE_NAME=mcs2lm
VOLUME_NAME=mcs2lm
DOCKER_RUN_CMD_PREFIX=docker run \
-ti \
--mount source=${VOLUME_NAME},destination=/app \
--rm ${DOCKER_IMAGE_NAME} \
 /bin/sh -c

setup-githooks:
	@ln -s ${PWD}/scripts/githooks/precommit/pre-commit ${PWD}/.git/hooks/pre-commit

setup: 
	docker build . -t ${DOCKER_IMAGE_NAME} && \
	docker volume create ${VOLUME_NAME}

build:
	${DOCKER_RUN_CMD_PREFIX} "go build -o ${BIN_PATH}/${BIN_NAME}"

test:
	${DOCKER_RUN_CMD_PREFIX} "go test ./... -v"
