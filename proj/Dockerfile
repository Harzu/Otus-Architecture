FROM golang:1.13.4 as builder

WORKDIR ${GOPATH}/src/hw
RUN PATH=${PATH}:${GOPATH}

COPY . .
ENV GO111MODULE=on

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /bin/hw

FROM alpine:3.9 as production

RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/hw /bin/hw

ENTRYPOINT ["/bin/hw"]