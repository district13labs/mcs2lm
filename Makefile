.PHONY: setup-githooks setup build purge

include .env

DOCKER=true

setup-githooks:
	@DOCKER=${DOCKER} ${PWD}/scripts/githooks/precommit/setup-precommit

setup:
	docker build -t ${DOCKER_IMAGE_NAME} -f ${PWD}/Dockerfile --build-arg USER_ID=$(shell id -u) --build-arg GROUP_ID=$(shell id -g) --build-arg BIN_PATH=${BIN_PATH} .
	docker volume create ${VOLUME_NAME}

build:
	@DOCKER=${DOCKER} ${PWD}/scripts/app/build

purge:
	docker rmi ${DOCKER_IMAGE_NAME} && \
	docker volume rm ${VOLUME_NAME}
