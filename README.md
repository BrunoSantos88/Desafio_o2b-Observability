O que é Prometeus?
Prometheus é um kit de ferramentas de alerta e monitoramento de sistemas de código aberto originalmente desenvolvido no SoundCloud . Desde a sua criação em 2012, muitas empresas e organizações adotaram o Prometheus, e o projeto tem uma comunidade de desenvolvedores e usuários muito ativa . Agora é um projeto independente de código aberto e mantido de forma independente por qualquer empresa. Para enfatizar isso e esclarecer a estrutura de governança do projeto, o Prometheus juntou-se à Cloud Native Computing Foundation em 2016 como o segundo projeto hospedado, depois do Kubernetes .

O Prometheus coleta e armazena suas métricas como dados de séries temporais, ou seja, as informações das métricas são armazenadas com o carimbo de data e hora em que foram registradas, juntamente com pares de valores-chave opcionais chamados rótulos.

Para visões gerais mais elaboradas do Prometheus, consulte os recursos vinculados na seção de mídia .

Características
As principais características do Prometheus são:

um modelo de dados multidimensional com dados de série temporal identificados por nome de métrica e pares chave/valor
PromQL, uma linguagem de consulta flexível para aproveitar essa dimensionalidade
nenhuma dependência de armazenamento distribuído; nós de servidor único são autônomos
a coleta de série temporal acontece por meio de um modelo pull sobre HTTP
o envio de séries temporais é suportado por meio de um gateway intermediário
os alvos são descobertos por meio de descoberta de serviço ou configuração estática
vários modos de suporte a gráficos e painéis
O que são métricas?
Métricas são medidas numéricas em termos leigos. O termo série temporal refere-se ao registro de mudanças ao longo do tempo. O que os usuários desejam medir difere de aplicativo para aplicativo. Para um servidor web, podem ser tempos de solicitação; para um banco de dados, pode ser o número de conexões ativas ou de consultas ativas e assim por diante.

As métricas desempenham um papel importante na compreensão de por que seu aplicativo está funcionando de determinada maneira. Vamos supor que você esteja executando uma aplicação web e descubra que ela está lenta. Para saber o que está acontecendo com seu aplicativo, você precisará de algumas informações. Por exemplo, quando o número de solicitações é alto, a aplicação pode ficar lenta. Se você tiver a métrica de contagem de solicitações, poderá determinar a causa e aumentar o número de servidores para lidar com a carga.

Componentes
O ecossistema Prometheus consiste em vários componentes, muitos dos quais são opcionais:

o principal servidor Prometheus que coleta e armazena dados de séries temporais
bibliotecas de cliente para instrumentação de código de aplicativo
um gateway push para apoiar empregos de curta duração
exportadores para fins especiais de serviços como HAProxy, StatsD, Graphite, etc.
um gerenciador de alertas para lidar com alertas
diversas ferramentas de suporte
A maioria dos componentes do Prometheus são escritos em Go , tornando-os fáceis de construir e implantar como binários estáticos.


# Clona repository

- Acessar pasta Prometheus-Dockerizado
- commando docker-compose up -d
- python /python-app app.py

# Install

- pip install Flask prometheus_client
- python python-app/app.py
  
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


# Promethues 
- Link https://prometheus.io/docs/introduction/overview/

# Webhook
- Link https://webhook.site/#!/49627f3d-1930-47af-8c80-2a63f9378bcd/1140d711-fad3-4189-8e9e-f1a6b706d7a3/1
