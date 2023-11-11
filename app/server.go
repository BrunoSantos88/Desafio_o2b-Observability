package main

import (
   "fmt"
   "net/http"
   "fmt"
   "log"
   "net"

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

func ping(w http.ResponseWriter, req *http.Request) {
   pingCounter.Inc()
   fmt.Fprintf(w, "pong")
}

func loadGeoIPDatabase() (*geoip2.Reader, error) {
   db, err := geoip2.Open("path/to/GeoIP2-City.mmdb")
   if err != nil {
       return nil, err
   }
   return db, nil
}

func main() {
   prometheus.MustRegister(pingCounter)

   http.HandleFunc("/ping", ping)
   http.Handle("/metrics", promhttp.Handler())
   http.ListenAndServe(":8090", nil)
}

func getGeoIPInfo(db *geoip2.Reader, ipAddress string) {
   ip := net.ParseIP(ipAddress)
   record, err := db.City(ip)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Printf("City: %s\n", record.City.Names["en"])
   fmt.Printf("Country: %s\n", record.Country.Names["en"])
   fmt.Printf("Latitude: %f\n", record.Location.Latitude)
   fmt.Printf("Longitude: %f\n", record.Location.Longitude)
}

func closeGeoIPDatabase(db *geoip2.Reader) {
   db.Close()
}