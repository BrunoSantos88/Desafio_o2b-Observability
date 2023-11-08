#!bin/bash
add-apt-repository ppa:longsleep/golang-backports -y
apt update -y
apt install golang-1.18 -y
export PATH=$PATH:/usr/local/go/bin
