FROM golang:1.19 as builder

WORKDIR /go/src/gindemo

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct go build -ldflags="-w -s" -o gindemo .


FROM alpine:latest
WORKDIR /usr/local/demo
COPY --from=builder /go/src/gindemo/gindemo .
EXPOSE 5000
CMD ["./gindemo"]