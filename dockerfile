FROM golang:1.18
WORKDIR /go/src/app
COPY * main.go
RUN go mod init prom_example
RUN go mod tidy
RUN go build main.go
CMD ["./main"] 
EXPOSE 8080
