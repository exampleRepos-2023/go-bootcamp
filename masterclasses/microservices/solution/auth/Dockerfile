FROM golang:1.18.0-alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build ./main/server.go
CMD ["/app/server"]
