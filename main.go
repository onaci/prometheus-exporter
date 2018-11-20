package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	cpuTemp.Set(65.3)
	hdFailures.Inc()

	reg := prometheus.NewPedanticRegistry()

	NewEnvCollectorGatherer("environment", reg)
	NewDockerCollectorGatherer(name, reg)

	reg.MustRegister(cpuTemp)
	reg.MustRegister(hdFailures)

	// The built in process and golang metrics
	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	endpoint := "/metrics"
	fmt.Printf("Listening on %s\n", endpoint)
	http.Handle(endpoint, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
