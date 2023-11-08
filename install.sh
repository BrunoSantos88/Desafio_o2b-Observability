#!bin/bash
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update -y
sudo apt install golang-1.18 -y
sudo export PATH=$PATH:/usr/local/go/bin
