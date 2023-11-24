O que é Prometeus?
Prometheus é um kit de ferramentas de alerta e monitoramento de sistemas de código aberto originalmente desenvolvido no SoundCloud . Desde a sua criação em 2012, muitas empresas e organizações adotaram o Prometheus, e o projeto tem uma comunidade de desenvolvedores e usuários muito ativa . Agora é um projeto independente de código aberto e mantido de forma independente por qualquer empresa. Para enfatizar isso e esclarecer a estrutura de governança do projeto, o Prometheus juntou-se à Cloud Native Computing Foundation em 2016 como o segundo projeto hospedado, depois do Kubernetes .

# Clona repository

``` Acessar pasta Desafio_o2b
git clone https://github.com/BrunoSantos88/Desafio_o2b-Observability.git
cd Desafio_o2b-Observability
``` 

# Install
``` Dependencias caso seja feito local
 pip install Flask prometheus_client
```
# Execute 
````
python aplication/app.py
````
docker.compose.yml
``` yml
version: '3'

services:

  prometheus:
    image: prom/prometheus
    ports:
      - 9090:9090
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
      - ./prometheus/rules.yml:/etc/prometheus/rules.yml
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
    network_mode: "host"
 
  alertmanager:
    image: prom/alertmanager
    container_name: alertmanager
    ports:
     - '9093:9093'
    volumes:
      - ./alertmanager.yml:/etc/alertmanager/alertmanager.yml
    network_mode: "host"

  cadvisor:
    image: gcr.io/cadvisor/cadvisor:latest
    container_name: cadvisor
    ports:
    - 8080:8080
    volumes:
    - /:/rootfs:ro
    - /var/run:/var/run:rw
    - /sys:/sys:ro
    - /var/lib/docker/:/var/lib/docker:ro
    network_mode: "host"

  grafana:
    image: grafana/grafana
    #image: grafana/grafana:8.4.11
    #image: grafana/grafana:8.5.27
    #image: grafana/grafana:9.0.9
    #image: grafana/grafana:9.1.8
    #image: grafana/grafana:9.2.20
    #image: grafana/grafana:9.3.16
    #image: grafana/grafana:9.4.13
    #image: grafana/grafana:9.5.6
   #image: grafana/grafana:10.0.3
    container_name: grafana
    ports:
      - 3000:3000
    network_mode: "host"

volumes:
  grafana-data:

````
# killerKoda
- Linux (baseado no Ubuntu)
- Docker
- Prometheus
- Grafana
- Alertmanager
- python
  
# Link Guias

- Setup observability
- Link: https://dev.to/danielfavour/container-monitoring-ensuring-application-performance-and-health-kcj

- Alert Manager
- Link: https://samber.github.io/awesome-prometheus-alerts/rules.html#docker-containers

# Promethues 
- Link https://prometheus.io/docs/introduction/overview/

# prometheus.yml
  ```
global:
  scrape_interval: 5s
  evaluation_interval: 10s

  
rule_files:
  - rules.yml
alerting:
  alertmanagers:
  - static_configs:
    - targets:
       - localhost:9093
       
scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
  - job_name: 'node-exporter'
    static_configs:
      - targets: ['localhost:9100']
  - job_name: python_server
    scrape_interval: 5s
    metrics_path: /metrics
    static_configs:
      - targets: ["localhost:3001"]
  - job_name: cadvisor
    scrape_interval: 5s
    static_configs:
    - targets: ["localhost:8080"]
  - job_name: 'node'
    static_configs:
      - targets: ['localhost:9100']

```

# Rules
- Link https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/

  # rules.yml
````
  groups:
  #Request contagem acesso
 - name: Count greater than 5
   rules:
   - alert: CountGreaterThan5
     expr: ping_request_count > 5
     for: 5s
     
# ALERT UP TIME
 - name: Aplication
   rules:
   - alert: InstanceDown
     expr: up{instance="localhost:3001", job="pythom_server"} == 0
     for: 5s
     annotations:
       title: 'Instance {{ $labels.instance }} down'
       description: '{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 1 minute.'
     labels:
       severity: 'critical'
 
  # Atingir espaço abaixo de 1GB
 - name: Uso de disk % maximo
   rules:
   - alert: LowInodesAlert
     expr: node_filesystem_avail_bytes{mountpoint="/", instance="localhost:9200", job="node-exporter"} / node_filesystem_size_bytes{mountpoint="/", instance="localhost:9200", job="node-exporter"} * 100 < 10
     labels:
       severity: warning
     annotations:
      summary: "Low Inodes"
      description: "The number of free inodes on /dev/vda1 is below 1000."

  # Link de Dados
 - name: NetworkDown
   rules:
   - alert: NetworkDownAlert
     expr: node_network_up{device="enp1s0", instance="localhost:9200", job="node-exporter"}
     for: 5s
     labels:
       severity: critical
     annotations:
      summary: "Network Interface Down"
      description: "The network interface 'enp1s0' is down."

 - name: Memory Usage
   rules:
   - alert: MemoryUsageAlert
     expr: node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes
     for: 5s
     labels:
       severity: critical
     annotations:
      summary: "High Memory Usage"
      description: "Memory usage is above 90%."
  ````

# AlertMaanager

- Link: https://prometheus.io/docs/alerting/latest/alertmanager/

  # alertmanager.yml

````
route:
  group_by: ['alertname', 'job']

  group_wait: 5s
  group_interval: 5s
  repeat_interval: 5s

  receiver: discord

receivers:
- name: discord
  discord_configs:
  - webhook_url: "https://discord.com/api/webhooks/#############################"
  ````
  

# Webhook
- Link https://webhook.site/#!/49627f3d-1930-47af-8c80-2a63f9378bcd/1140d711-fad3-4189-8e9e-f1a6b706d7a3/1

# install node_exported
Link: https://developer.couchbase.com/tutorial-node-exporter-setup

# Grafana Dashbord
Link: https://grafana.com/grafana/dashboards/

# Install Stress

````
sudo apt-get install stress-ng -y
stress-ng --cpu 4 --io 2 --vm 1 --vm-bytes 256M --timeout 60s
````
# Grafana DashBord

<img src="https://private-user-images.githubusercontent.com/91704169/285541482-360bb4dc-9dd7-49d3-a4f7-a22a266dc733.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTEiLCJleHAiOjE3MDA4NjAxNzIsIm5iZiI6MTcwMDg1OTg3MiwicGF0aCI6Ii85MTcwNDE2OS8yODU1NDE0ODItMzYwYmI0ZGMtOWRkNy00OWQzLWE0ZjctYTIyYTI2NmRjNzMzLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFJV05KWUFYNENTVkVINTNBJTJGMjAyMzExMjQlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjMxMTI0VDIxMDQzMlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWMzY2U2Mzk2ZmI5NTFkZjE2MDI0MjY5MzU2MTc3NTgxZTYwZWU1ZWExMjZiMTUxNGQ0NWEzMDIyNTA0ODQ1NzImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.XXqBLE33SNm_XSI7nmLjkcMRlZusIWQIiQqcPbwwwKo" min-width="300px" max-width="900px" width="900px" align="center" alt="Computador illustration">

 # DISCORD ALERT
<img src="https://private-user-images.githubusercontent.com/91704169/285544410-b2b6f5af-bdcc-4ce3-88c2-9eb2a69ab124.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTEiLCJleHAiOjE3MDA4NjA1NTAsIm5iZiI6MTcwMDg2MDI1MCwicGF0aCI6Ii85MTcwNDE2OS8yODU1NDQ0MTAtYjJiNmY1YWYtYmRjYy00Y2UzLTg4YzItOWViMmE2OWFiMTI0LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFJV05KWUFYNENTVkVINTNBJTJGMjAyMzExMjQlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjMxMTI0VDIxMTA1MFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTI4ZmViN2YyMTg0ZjY1N2YzZmIxZTJjZmNlYjcxMzJjNzRjZDljOTlhOGM0OWM0MGRmZTYxNzdjYjk3ZTlkMmUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.WEQptAefwxiAj_y86QnJzdde_H72_l8UpJGnYsgqa1Y" min-width="300px" max-width="900px" width="900px" align="rigth" alt="Computador illustration">

# Link video: https://www.youtube.com/watch?v=8s0hic3XusA
