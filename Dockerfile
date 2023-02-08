FROM golang:1.19-alpine AS builder

ARG GIT_COMMIT=unspecified
ARG BUILD_DATE=unspecified
ARG SERVICE_NAME=unspecified

LABEL GIT_COMMIT=$GIT_COMMIT
LABEL BUILD_DATE=$BUILD_DATE
LABEL SERVICE_NAME=$SERVICE_NAME


ENV GOPATH=/go

RUN apk add --no-cache make gcc musl-dev linux-headers git gettext

ADD . /workspace

WORKDIR /workspace

RUN go build -mod vendor -ldflags "-s -w -X main.version=${BUILD_DATE}:${GIT_COMMIT} -X main.service=${SERVICE_NAME}" -o /app ./cmd/server

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app /app

CMD ["/app"]
