FROM golang:1.25.10-bookworm

ARG APP_HOME=/app
ARG APP_USER=app

WORKDIR ${APP_HOME}

RUN apt-get update && apt-get install -y \
    curl \
    sudo \
    vim \
    git \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

RUN groupadd --gid 1000 ${APP_USER} \
    && useradd --uid 1000 --gid ${APP_USER} --shell /bin/bash --create-home ${APP_USER}

RUN apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false \
    && rm -rf /var/lib/apt/lists/*

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./.air.toml ./.air.toml
COPY ./src/ ./src

COPY ./compose/backend/start /start
RUN chmod +x /start

USER ${APP_USER}
