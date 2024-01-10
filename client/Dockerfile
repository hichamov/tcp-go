FROM golang:1.21.5
WORKDIR /app
COPY main.go /app
RUN go mod init client
RUN go mod tidy
CMD "go run /app/main.go"
