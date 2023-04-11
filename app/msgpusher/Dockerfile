FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src/app/msgpusher

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
RUN make build

FROM debian:stable-slim

RUN  sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/msgpusher/bin /app

WORKDIR /app

EXPOSE 7000
EXPOSE 8001
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]