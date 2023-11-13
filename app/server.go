package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()

	// Adiciona rastreamento ao manipulador "ping"
	span, ctx := opentracing.StartSpanFromContext(req.Context(), "ping-handler")
	defer span.Finish()

	fmt.Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)

	// Configuração do Jaeger
	cfg := &config.Configuration{
		ServiceName: "meu-app-go",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Metrics(metrics.NullFactory),
	)
	if err != nil {
		fmt.Printf("Erro ao criar o tracer: %v\n", err)
		os.Exit(1)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8090", nil)
}
