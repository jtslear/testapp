FROM golang:1.11.4 as builder

RUN apt update && \
    apt install -y upx

COPY helloweb.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o testapp .
RUN upx -9 -o testapp.up testapp

FROM scratch

EXPOSE 8000

COPY --from=builder /go/testapp.up /

CMD ["/testapp.up"]
