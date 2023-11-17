#!bin/bash
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update -y
sudo apt install golang-1.18 -y
export PATH=$PATH:/usr/local/go/bin
go mod init prom_example
go mod tidy
go build server.go
./server