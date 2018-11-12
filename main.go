package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	hdFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	})
)

func init() {
	// Registered to the /metrics-default handler
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	cpuTemp.Set(65.3)
	hdFailures.Inc()

	endpoint := "/metrics"
	fmt.Printf("Listening on %s\n", endpoint)
	http.Handle(endpoint, prometheus.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
