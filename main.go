package main

import (
    "log"
    "net/http"
    httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
    "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
    rules := []tracer.SamplingRule{tracer.RateRule(1)}
    tracer.Start(
        tracer.WithSamplingRules(rules),
        tracer.WithService("service"),
        tracer.WithEnv("env"),
    )
    defer tracer.Stop()

    // Create a traced mux router
    mux := httptrace.NewServeMux()
    // Continue using the router as you normally would.
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
 }

 func main() {
    prometheus.MustRegister(pingCounter)
 
    http.HandleFunc("/ping", ping)
 }
 