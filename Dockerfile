FROM golang:latest
RUN mkdir /app
RUN mkdir /public
ADD ./src /app/
ADD ./public /app/public/
WORKDIR /app
RUN go get github.com/gorilla/websocket
RUN go build -o main .
CMD ["/app/main"]
