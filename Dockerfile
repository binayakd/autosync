FROM golang:latest AS builder

COPY . /go/src/github.com/binayakd/autosync
WORKDIR /go/src/github.com/binayakd/autosync

RUN go build

FROM alpine:latest

RUN apk --no-cache add curl unzip 

RUN addgroup -g 1000 autosync && adduser -u 1000 --disabled-password autosync -G autosync
USER 1000
WORKDIR /home/autosync
COPY --from=builder --chown=1000:1000 /go/src/github.com/binayakd/autosync/autosync .
RUN mkdir data configs bin && \
  curl -O https://downloads.rclone.org/rclone-current-linux-amd64.zip && \
  unzip rclone-current-linux-amd64.zip && \
  cp rclone-*-linux-amd64/rclone ./bin/ && \
  chmod +x ./bin/rclone && \
  chmod +x autosync && \
  rm -r ./rclone-*-linux-amd64 && \
  rm rclone-current-linux-amd64.zip 

ENV PATH=/home/autosync/bin:${PATH}
ENV XDG_CONFIG_HOME=/home/autosync/configs

VOLUME [ "/home/autosync/data", "/home/autosync/config" ]

RUN rclone version 

CMD [ "./autosync" ]