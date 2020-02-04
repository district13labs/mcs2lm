FROM district13labs/golang:1.13

ARG USER_ID
ARG GROUP_ID

RUN mkdir /.cache /app-bin && \
    chown -R ${USER_ID}:${GROUP_ID} /.cache /app-bin
