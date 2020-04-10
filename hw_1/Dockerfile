FROM golang:1.13.4 as builder

WORKDIR ${GOPATH}/src/hw_1
RUN PATH=${PATH}:${GOPATH}
COPY . .

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /bin/hw_1

FROM alpine:3.9 as production

RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/hw_1 /bin/hw_1

ENTRYPOINT ["/bin/hw_1"]