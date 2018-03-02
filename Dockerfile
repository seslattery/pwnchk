FROM golang:1.10

RUN mkdir -p /go/src/github.com/seslattery/pwnchk
WORKDIR /go/src/github.com/seslattery/pwnchk
RUN apt update && \
    apt install ca-certificates wget -y
RUN go get golang.org/x/crypto/ssh/terminal
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static -s -w" -a src/main.go
RUN wget https://github.com/lalyos/docker-upx/releases/download/v3.91/upx -O /bin/upx
RUN chmod +x /bin/upx && \
    upx -f --brute -o /main ./main


FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /main /
CMD ["/main"]
