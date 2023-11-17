#!bin/bash
#nodexport
wget https://github.com/prometheus/node_exporter/releases/download/v*/node_exporter-*.*-amd64.tar.gz
tar xvfz node_exporter-*.*-amd64.tar.gz
sudo mv node_exporter-*.*-amd64/node_exporter /usr/local/bin/
#alertmaanager
wget https://github.com/prometheus/alertmanager/releases/download/v0.20.0/alertmanager-0.20.0.linux-amd64.tar.gz
tar xvfz alertmanager-0.20.0.linux-amd64.tar.gz
sudo cp alertmanager-0.20.0.linux-amd64/alertmanager /usr/local/bin/
sudo chown alertmanager:alertmanager /usr/local/bin/alertmanager
sudo mkdir -p /etc/alertmanager
sudo cp alertmanager-0.20.0.linux-amd64/alertmanager.yml /etc/alertmanager
sudo chown -R alertmanager:alertmanager /etc/alertmanager
sudo mkdir -p /var/lib/alertmanager
sudo chown alertmanager:alertmanager /var/lib/alertmanager
#promethues
wget https://github.com/prometheus/prometheus/releases/download/v*/prometheus-*.*-amd64.tar.gz
tar xvf prometheus-*.*-amd64.tar.gz
cd prometheus-*.*
./prometheus --config.file=./prometheus.yml
#godependecias
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update -y
sudo apt install golang-1.18 -y
export PATH=$PATH:/usr/local/go/bin
go mod init prom_example
go mod tidy
go build server.go
./server