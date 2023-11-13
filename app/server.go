package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	pingCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ping_request_count",
			Help: "Number of requests handled by the Ping handler",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(pingCounter)

	// Configure Jaeger tracer
	cfg := &config.Configuration{
		ServiceName: "my-go-app",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:      true,
			AgentHostPort: "jaeger-agent:6831", // Adjust this to your Jaeger agent's address
		},
	}

	metricsFactory := prometheus.New()
	tracer, _, err := cfg.New(
		"jaeger-go-app",
		config.Logger(jaeger.StdLogger),
		config.Metrics(metricsFactory),
	)
	if err != nil {
		fmt.Println("Failed to initialize Jaeger tracer:", err)
		return
	}

	opentracing.SetGlobalTracer(tracer)
}

func ping(w http.ResponseWriter, req *http.Request) {
	// Start a new span for the ping operation
	span := opentracing.StartSpan("ping")
	defer span.Finish()

	pingCounter.WithLabelValues("success").Inc()

	// Your logic here

	fmt.Fprintf(w, "pong")
}

func main() {
	// Register Prometheus metrics handler
	http.Handle("/metrics", promhttp.Handler())

	// Register your ping handler
	http.HandleFunc("/ping", ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
