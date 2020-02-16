.PHONY: setup-githooks setup build purge

include .env

setup-githooks:
	@DOCKER=${DOCKER} ${PWD}/scripts/githooks/precommit/setup-precommit

setup:
	docker build -t ${DOCKER_IMAGE_NAME} -f ${PWD}/Dockerfile --build-arg USER_ID=$(id -u) --build-arg GROUP_ID=$(id -g) --build-arg BIN_PATH=${BIN_PATH}. 
	docker volume create ${VOLUME_NAME}

build:
	docker run \
	-u $(id -u):$(id -g) \
	-ti \
	-v ${VOLUME_NAME}:${BIN_PATH} \
	-v ${PWD}:${APP_PATH} \
	--rm \
	${DOCKER_IMAGE_NAME} \
	go build -o ${BIN_PATH}/${BIN_NAME} ./cmd

purge:
	docker rmi ${DOCKER_IMAGE_NAME} && \
	docker volume rm ${VOLUME_NAME}
