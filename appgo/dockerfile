FROM golang:1.18
WORKDIR /go/src/app
COPY . .
RUN go mod init prom_example
RUN go mod tidy
RUN go build server.go
CMD ["./server"] 
EXPOSE 8090
