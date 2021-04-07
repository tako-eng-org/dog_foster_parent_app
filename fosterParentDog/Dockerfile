FROM golang:1.15-alpine as dev

RUN apk update \
  && apk add --no-cache git \
  && apk add --no-cache tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  # エラー対策
  # cgo: exec gcc: exec: "gcc": executable file not found in $PATH
  && export CGO_ENABLED=0 \
  && GO111MODULE=off go get -u github.com/derekparker/delve/cmd/dlv \
  && GO111MODULE=off go get -u github.com/oxequa/realize

ARG APP_DIR=app

COPY ./ /${APP_DIR}
WORKDIR /${APP_DIR}

EXPOSE 8090 2345

ENTRYPOINT [ "/app/docker-entrypoint.sh" ]
CMD [ "/sbin/init" ]
