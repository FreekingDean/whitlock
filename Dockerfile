############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/client
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/whitlock-client

############################
# STEP 2 build a small image
############################
FROM scratch
COPY --from=builder /go/bin/whitlock-client /client
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV PORT 3000
EXPOSE 3000
ENV PROXY_TARGET whitlock.io 
ENTRYPOINT ["/client"]
