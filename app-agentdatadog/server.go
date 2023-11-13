package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "Number of requests handled by the Ping handler",
	},
	[]string{"status"},
)

func init() {
	prometheus.MustRegister(pingCounter)

	// Configure Datadog tracer
	tracer.Start(tracer.WithAgentAddr("localhost:8126")) // Adjust this to your Datadog Agent's address
}

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.WithLabelValues("success").Inc()

	// Your logic here

	fmt.Fprintf(w, "pong")
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
