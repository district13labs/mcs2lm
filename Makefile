.PHONY: setup-githooks setup build purge test shell

APP_PATH=/app
BIN_PATH=${APP_PATH}-bin
BIN_NAME=mcs2lm
DOCKER_IMAGE_NAME=mcs2lm
VOLUME_NAME=mcs2lm
DOCKER_RUN_CMD_PREFIX=docker run \
-u $(shell id -u):$(shell id -g) \
-ti \
-v ${VOLUME_NAME}:${BIN_PATH} \
-v ${PWD}:${APP_PATH} \
--rm \
${DOCKER_IMAGE_NAME} \
/bin/bash -c

setup-githooks:
	@ln -s ${PWD}/scripts/githooks/precommit/pre-commit ${PWD}/.git/hooks/pre-commit

setup: 
	docker build -t ${DOCKER_IMAGE_NAME} -f ${PWD}/Dockerfile --build-arg USER_ID=$(shell id -u) --build-arg GROUP_ID=$(shell id -g) .
	docker volume create ${VOLUME_NAME}

build:
	${DOCKER_RUN_CMD_PREFIX} "go build -o ${BIN_PATH}/${BIN_NAME}"

purge:
	docker rmi ${DOCKER_IMAGE_NAME} && \
	docker volume rm ${VOLUME_NAME}
