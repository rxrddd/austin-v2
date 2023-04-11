FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

RUN GOPROXY=https://goproxy.cn && go mod tidy
RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

VOLUME /data/conf

EXPOSE 8000

CMD ["./server", "-conf", "/data/conf"]
