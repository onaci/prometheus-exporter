# prometheus-exporter

A quick little golang program that exports some information from the docker daemon to prometheus

```
curl -q http://localhost:8080/metrics/docker 
# HELP Docker_swarm_nodes Number Docker Swarm nodes.
# TYPE Docker_swarm_nodes gauge
Docker_swarm_nodes{availability="active",docker="host",host="TOWER-SL",leader="true",reachability="reachable",role="manager",state="ready",version="18.06.1-ce"} 1
Docker_swarm_nodes{availability="active",docker="host",host="oa-29-mel",leader="false",reachability="worker",role="worker",state="ready",version="18.06.1-ce"} 1
Docker_swarm_nodes{availability="active",docker="host",host="oa-30-mel",leader="false",reachability="worker",role="worker",state="ready",version="18.06.1-ce"} 1
```

and

```
curl -q http://localhost:8080/metrics
# HELP cpu_temperature_celsius Current temperature of the CPU.
# TYPE cpu_temperature_celsius gauge
cpu_temperature_celsius 65.3
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
...
```