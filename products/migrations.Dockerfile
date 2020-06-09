FROM golang:1.13 as builder

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

ENTRYPOINT ["/bin/goose"]
