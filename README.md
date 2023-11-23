O que é Prometeus?
Prometheus é um kit de ferramentas de alerta e monitoramento de sistemas de código aberto originalmente desenvolvido no SoundCloud . Desde a sua criação em 2012, muitas empresas e organizações adotaram o Prometheus, e o projeto tem uma comunidade de desenvolvedores e usuários muito ativa . Agora é um projeto independente de código aberto e mantido de forma independente por qualquer empresa. Para enfatizar isso e esclarecer a estrutura de governança do projeto, o Prometheus juntou-se à Cloud Native Computing Foundation em 2016 como o segundo projeto hospedado, depois do Kubernetes .

# Clona repository

``` Acessar pasta Desafio_o2b
docker-compose up -d
``` 

# Install
``` Dependencias
 pip install Flask prometheus_client
``` 
``` Excutar
  python python-app/app.py
```   
# killerKoda
- Linux (baseado no Ubuntu)
- Docker
- Golang
- Prometheus
- Grafana
- Alertmanager
- python
  
# Link Guias

- Setup observability
- Link: https://dev.to/danielfavour/container-monitoring-ensuring-application-performance-and-health-kcj

- Alert Manager
- Link: https://samber.github.io/awesome-prometheus-alerts/rules.html#docker-containers

# Node_exported install

``` install node_exporter

curl -LO https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz
mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/
useradd -rs /bin/false node_exporter

nano /etc/systemd/system/node_exporter.service

[Unit]
Description=Node Exporter
After=network.target

[Service] 
User=node_exporter  
Group=node_exporter 
Type=simple 
ExecStart=/usr/local/bin/node_exporter </p>

[Install] 
WantedBy=multi-user.target

systemctl daemon-reload
systemctl start node_exporter
systemctl enable node_exporter
curl http://<server-IP>:9100/metrics
``` 

# Promethues 
- Link https://prometheus.io/docs/introduction/overview/

# Rules
- Link https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/

# AlertMaanager

- Link: https://prometheus.io/docs/alerting/latest/alertmanager/

# Webhook
- Link https://webhook.site/#!/49627f3d-1930-47af-8c80-2a63f9378bcd/1140d711-fad3-4189-8e9e-f1a6b706d7a3/1
