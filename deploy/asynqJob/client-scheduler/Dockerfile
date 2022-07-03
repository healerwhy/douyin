FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../../../service/asynqJob/client-scheduler .
COPY service/asynqJob/client-scheduler/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/scheduler service/asynqJob/client-scheduler/scheduler.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/scheduler /app/scheduler
COPY --from=builder /app/etc /app/etc

CMD ["./scheduler", "-f", "etc/scheduler.yaml"]
