package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

   "github.com/gin-gonic/gin"
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

router.GET("/", func(c *gin.Context) {
	// Get the approximate location of the client's IP address
	ipAddress := c.ClientIP()
	location, err := getLocation(ipAddress)
	if err != nil {
		c.String(500, "Error fetching GeoIP information")
		return
	}

	// Increment GeoIP city hits metric
	labels := prometheus.Labels{
		"city":   location.City,
		"country": location.Country,
	}
	geoipCityHits.With(labels).Inc()

	// Respond with Hello and GeoIP location
	c.String(200, fmt.Sprintf("Hello from %s!", location))
})

func main() {
   prometheus.MustRegister(pingCounter)

   http.HandleFunc("/ping", ping)
   http.Handle("/metrics", promhttp.Handler())
   http.ListenAndServe(":8090", nil)
}

type locationInfo struct {
	City    string
	Country string
}

func logRequest(c *gin.Context) {
	// Log the incoming request
	log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL)
	c.Next()
}

func getLocation(ipAddress string) (locationInfo, error) {
	// Query GeoIP information
	ip := net.ParseIP(ipAddress)
	record, err := geoipDatabase.City(ip)
	if err != nil {
		return locationInfo{}, err
	}

	// Construct a location struct
	location := locationInfo{
		City:    record.City.Names["en"],
		Country: record.Country.Names["en"],
	}
	return location, nil
}