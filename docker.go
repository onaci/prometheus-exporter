package main

import (
	"context"

	"github.com/docker/docker/api/types"
	//"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"

	"github.com/prometheus/client_golang/prometheus"
)

var name = "docker"

// DockerCollectorGathererCollector implements the Collector interface.
type DockerCollectorGathererCollector struct {
}

// Descriptors used by the DockerCollectorGathererCollector below.
var (
	dockerSwarmNodesDesc = prometheus.NewDesc(
		"Docker_swarm_nodes",
		"Number of Docker Swarm nodes.",
		[]string{
			"host",
			"role",
			"availability",
			"state",
			"version",
			"reachability",
			"leader",
		}, nil,
	)
)

// Describe is implemented with DescribeByCollect. That's possible because the
// Collect method will always return the same two metrics with the same two
// descriptors.
func (cc DockerCollectorGathererCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(cc, ch)
}

// Collect first triggers the ReallyExpensiveAssessmentOfTheSystemState. Then it
// creates constant metrics for each host on the fly based on the returned data.
//
// Note that Collect could be called concurrently, so we depend on
// ReallyExpensiveAssessmentOfTheSystemState to be concurrency-safe.
func (cc DockerCollectorGathererCollector) Collect(ch chan<- prometheus.Metric) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		leader := "false"
		reachable := "worker"
		if node.ManagerStatus != nil {
			if node.ManagerStatus.Leader {
				leader = "true"
			}
			reachable = string(node.ManagerStatus.Reachability)
		}
		ch <- prometheus.MustNewConstMetric(
			dockerSwarmNodesDesc,
			prometheus.GaugeValue,
			1,
			node.Description.Hostname,
			string(node.Spec.Role),
			string(node.Spec.Availability),
			string(node.Status.State),
			node.Description.Engine.EngineVersion,
			reachable,
			leader,
		)
	}
}

// NewDockerCollectorGatherer first creates a Prometheus-ignorant DockerCollectorGatherer
// instance. Then, it creates a DockerCollectorGathererCollector for the just created
// DockerCollectorGatherer. Finally, it registers the DockerCollectorGathererCollector with a
// wrapping Registerer that adds the zone as a label. In this way, the metrics
// collected by different DockerCollectorGathererCollectors do not collide.
func NewDockerCollectorGatherer(zone string, reg prometheus.Registerer) {
	cc := DockerCollectorGathererCollector{}
	prometheus.WrapRegistererWith(prometheus.Labels{"type": "docker"}, reg).MustRegister(cc)
}
