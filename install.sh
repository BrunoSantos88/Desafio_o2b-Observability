#!bin/bash
wget https://github.com/prometheus/prometheus/releases/download/v*/prometheus-*.*-amd64.tar.gz
tar xvf prometheus-*.*-amd64.tar.gz
wget https://github.com/prometheus/alertmanager/releases/download/v0.26.0/alertmanager-0.26.0.linux-amd64.tar.gz
tar xvf alertmanager-*.*-amd64.tar.gz
mkdir -p /etc/alertmanager
mkdir -p /var/lib/alertmanager
mkdir -p /etc/amtool
cp alertmanager-0.25.0.linux-amd64/alertmanager /usr/local/bin/
cp alertmanager-0.25.0.linux-amd64/alertmanager.yml /etc/alertmanager
cp alertmanager-0.25.0.linux-amd64/amtool /usr/local/bin/
sudo useradd -M -r -s /bin/false alertmanager
chown alertmanager:alertmanager /usr/local/bin/alertmanager
chown -R alertmanager:alertmanager /etc/alertmanager
chown alertmanager:alertmanager /var/lib/alertmanager
rm -f /etc/systemd/system/alertmanager.service
mv alertmanager.service /etc/systemd/system/
systemctl daemon-reload
sudo systemctl enable alertmanager
sudo systemctl start alertmanager
#godependecias
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update -y
sudo apt install golang-1.18 -y
export PATH=$PATH:/usr/local/go/bin
go mod init prom_example
go mod tidy
go build server.go
./server