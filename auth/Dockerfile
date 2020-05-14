FROM golang:1.13.4 as builder

WORKDIR ${GOPATH}/src/app
RUN PATH=${PATH}:${GOPATH}

ENV GO111MODULE=on
COPY . .

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /bin/app

FROM alpine:3.9 as production

RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["/bin/app"]