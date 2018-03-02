FROM golang:1.10

RUN mkdir -p /go/src/github.com/seslattery/pwnchk
WORKDIR /go/src/github.com/seslattery/pwnchk
RUN apt update && \
    apt install ca-certificates -y
RUN go get golang.org/x/crypto/ssh/terminal
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a src/main.go

FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/github.com/seslattery/pwnchk/main /main
CMD ["/main"]
