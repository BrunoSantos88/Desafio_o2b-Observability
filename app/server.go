package main

import (
	"fmt"
	"net/http"

   "github.com/oschwald/geoip2-golang"
   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "ping_request_count",
       Help: "No of request handled by Ping handler",
   },
)

var (
	geoipDatabase *geoip2.Reader

	geoipCityHits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "geoip_city_hits_total",
			Help: "Total number of GeoIP city hits",
		},
		[]string{"city", "country"},
	)
)

func init() {
	prometheus.MustRegister(geoipCityHits)
}
func ping(w http.ResponseWriter, req *http.Request) {
   pingCounter.Inc()
   fmt.Fprintf(w, "pong")
}

func main() {
   prometheus.MustRegister(pingCounter)

   http.HandleFunc("/ping", ping)
   http.Handle("/metrics", promhttp.Handler())
   http.ListenAndServe(":8090", nil)
}
