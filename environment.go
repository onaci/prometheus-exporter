package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	reg := prometheus.NewPedanticRegistry()

	NewEnvCollectorGatherer(name, reg)

	name := "environment"
	endpoint := "/metrics/" + name
	fmt.Printf("Listening on %s\n", endpoint)
	http.Handle(endpoint, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
}

// EnvCollectorGathererCollector implements the Collector interface.
type EnvCollectorGathererCollector struct {
}

func (cc EnvCollectorGathererCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(cc, ch)
}

var (
	envVarDesc = prometheus.NewDesc(
		"environment_variables",
		"Environment variables.",
		[]string{
			"name",
			"value",
		}, nil,
	)
)

// TODO: this exporter is actually static, so doesn't need to be done this way..

func (cc EnvCollectorGathererCollector) Collect(ch chan<- prometheus.Metric) {
	envStrings := os.Environ()
	for _, env := range envStrings {
		e := strings.Split(env, "=")

		ch <- prometheus.MustNewConstMetric(
			envVarDesc,
			prometheus.GaugeValue,
			1,
			e[0],
			e[1],
		)
	}
}

func NewEnvCollectorGatherer(zone string, reg prometheus.Registerer) {
	cc := EnvCollectorGathererCollector{}
	prometheus.WrapRegistererWith(prometheus.Labels{"type": "environment"}, reg).MustRegister(cc)
}
