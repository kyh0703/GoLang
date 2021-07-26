FROM golang:1.6-alpine
RUN mkdir /app
ADD ./tcp.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
CMD ["/app/main"]

EXPOSE 8000