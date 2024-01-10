FROM golang:1.21.3
WORKDIR /app
COPY main.go /app
RUN go mod init server
RUN go mod tidy
EXPOSE 2222
EXPOSE 3333
CMD go run /app/main.go
