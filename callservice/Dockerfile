FROM golang:1.16.7-alpine3.13

EXPOSE 9000

RUN apk add --update --no-cache git openssh
RUN echo -e "[url \"ssh://git@100.100.103.5/\"]\n\tinsteadOf = https://git.bridgetec.com/" >> /root/.gitconfig
RUN mkdir -p -m 0600 /root/.ssh && \
    echo "StrictHostKeyChecking no " > /root/.ssh/config &&

COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 ~/root/.ssh/id_rsa