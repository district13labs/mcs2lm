.PHONY: setup-githooks setup build purge test shell

include .env

DOCKER_RUN_CMD_PREFIX=docker run \
-u $(shell id -u):$(shell id -g) \
-ti \
-v ${VOLUME_NAME}:${BIN_PATH} \
-v ${PWD}:${APP_PATH} \
--rm \
${DOCKER_IMAGE_NAME} \
/bin/bash -c

setup-githooks:
	eval ${PWD}/scripts/githooks/precommit/setup-precommit

setup:
	docker build -t ${DOCKER_IMAGE_NAME} -f ${PWD}/Dockerfile --build-arg USER_ID=$(shell id -u) --build-arg GROUP_ID=$(shell id -g) --build-arg BIN_PATH=${BIN_PATH}. 
	docker volume create ${VOLUME_NAME}

build:
	${DOCKER_RUN_CMD_PREFIX} "go build -o ${BIN_PATH}/${BIN_NAME} ./cmd"

purge:
	docker rmi ${DOCKER_IMAGE_NAME} && \
	docker volume rm ${VOLUME_NAME}
