package main

import (
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

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
