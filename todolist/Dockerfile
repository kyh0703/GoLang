FROM golang
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod tidy && \
    go build
CMD ["/app/todo"]

EXPOSE 8000