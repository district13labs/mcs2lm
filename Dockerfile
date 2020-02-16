FROM district13labs/golang:1.13

ARG USER_ID
ARG GROUP_ID
ARG BIN_PATH

RUN mkdir /.cache ${BIN_PATH} && \
    chown -R ${USER_ID}:${GROUP_ID} /.cache ${BIN_PATH}
