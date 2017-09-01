FROM golang:1.8.1

WORKDIR /go/src/github.com/mizalewski/docker_stats_monitor
RUN go get -u github.com/kardianos/govendor

COPY . /go/src/github.com/mizalewski/docker_stats_monitor
RUN govendor sync

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=0 /go/src/github.com/mizalewski/docker_stats_monitor .

CMD ["./app"]
