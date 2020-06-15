FROM golang:alpine as builder

RUN apk add --no-cache git ca-certificates build-base
RUN go get -u github.com/pressly/goose/cmd/goose

WORKDIR ${GOPATH}/src/app
RUN PATH=${PATH}:${GOPATH}
COPY . .

RUN cp -R ./migrations /tmp/migrations
RUN cp ${GOPATH}/bin/goose /bin/goose

ENTRYPOINT []

FROM alpine as migrations

COPY --from=builder /tmp/migrations /tmp/migrations
COPY --from=builder /bin/goose /bin/goose

RUN chmod +x /bin/goose

WORKDIR /goose
