FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../../service/rpc-video-service .
COPY service/rpc-video-service/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/videoService service/rpc-video-service/videoService.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/videoService /app/videoService
COPY --from=builder /app/etc /app/etc

CMD ["./videoService", "-f", "etc/videoService.yaml"]
