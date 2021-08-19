FROM golang:1.16.7-alpine3.13 as builder
LABEL maintainer="kyh0703@bridgetec.co.kr"

RUN apk add --update --no-cache git openssh
RUN echo -e "[url \"ssh://git@100.100.103.5/\"]\n\tinsteadOf = https://git.bridgetec.com/" >> /root/.gitconfig
RUN mkdir -p -m 0600 /root/.ssh && \
    echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN go env -w GO111MODULE=on
RUN go env -w GOPRIVATE=github.com/kyh0703

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./