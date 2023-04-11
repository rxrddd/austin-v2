FROM golang:1.17 AS builder

ARG APP_RELATIVE_PATH

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
RUN make build

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

RUN  sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 8000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]