package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const YARN_SCRAPER_NAME = "yarn"

var yarnServiceTypes = []string{"YARN"}
var yarnRoleTypes = []string{"RESOURCEMANAGER", "NODEMANAGER", "JOBHISTORY"}

var yarnSpec = serviceScraperSpec{
	Name:         YARN_SCRAPER_NAME,
	Help:         "YARN service metrics",
	ServiceTypes: yarnServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "apps_running",
			Help:  "YARN running applications.",
			Query: serviceMetricQuery("apps_running", yarnServiceTypes),
		},
		{
			Name:  "apps_pending",
			Help:  "YARN pending applications.",
			Query: serviceMetricQuery("apps_pending", yarnServiceTypes),
		},
		{
			Name:  "apps_completed_rate",
			Help:  "YARN completed applications per second.",
			Query: serviceMetricQuery("apps_completed_rate", yarnServiceTypes),
		},
		{
			Name:  "apps_failed_rate",
			Help:  "YARN failed applications.",
			Query: serviceMetricQuery("apps_failed_rate", yarnServiceTypes),
		},
		{
			Name:  "apps_killed_rate",
			Help:  "YARN killed applications.",
			Query: serviceMetricQuery("apps_killed_rate", yarnServiceTypes),
		},
		{
			Name:  "allocated_containers",
			Help:  "YARN allocated containers.",
			Query: serviceMetricQuery("allocated_containers", yarnServiceTypes),
		},
		{
			Name:  "pending_containers",
			Help:  "YARN pending containers.",
			Query: serviceMetricQuery("pending_containers", yarnServiceTypes),
		},
		{
			Name:  "reserved_containers",
			Help:  "YARN reserved containers.",
			Query: serviceScopedMetricQuery("reserved_containers", yarnServiceTypes),
		},
		{
			Name:  "allocated_memory_mb",
			Help:  "YARN allocated memory in MB.",
			Query: serviceMetricQuery("allocated_memory_mb", yarnServiceTypes),
		},
		{
			Name:  "available_memory_mb",
			Help:  "YARN available memory in MB.",
			Query: serviceMetricQuery("available_memory_mb", yarnServiceTypes),
		},
		{
			Name:  "pending_memory_mb",
			Help:  "YARN pending memory in MB.",
			Query: serviceMetricQuery("pending_memory_mb", yarnServiceTypes),
		},
		{
			Name:  "reserved_memory_mb",
			Help:  "YARN reserved memory in MB.",
			Query: serviceMetricQuery("reserved_memory_mb", yarnServiceTypes),
		},
		{
			Name:  "allocated_vcores",
			Help:  "YARN allocated vcores.",
			Query: serviceMetricQuery("allocated_vcores", yarnServiceTypes),
		},
		{
			Name:  "available_vcores",
			Help:  "YARN available vcores.",
			Query: serviceMetricQuery("available_vcores", yarnServiceTypes),
		},
		{
			Name:  "pending_vcores",
			Help:  "YARN pending vcores.",
			Query: serviceMetricQuery("pending_vcores", yarnServiceTypes),
		},
		{
			Name:  "reserved_vcores",
			Help:  "YARN reserved vcores.",
			Query: serviceMetricQuery("reserved_vcores", yarnServiceTypes),
		},
		{
			Name:  "containers_running",
			Help:  "YARN running containers.",
			Query: roleMetricQuery("containers_running", yarnServiceTypes, "NODEMANAGER"),
		},
		{
			Name:  "containers_launched_rate",
			Help:  "YARN launched containers per second.",
			Query: roleMetricQuery("containers_launched_rate", yarnServiceTypes, "NODEMANAGER"),
		},
		{
			Name:  "containers_completed_rate",
			Help:  "YARN completed containers per second.",
			Query: roleMetricQuery("containers_completed_rate", yarnServiceTypes, "NODEMANAGER"),
		},
		{
			Name:  "containers_failed_rate",
			Help:  "YARN failed containers per second.",
			Query: roleMetricQuery("containers_failed_rate", yarnServiceTypes, "NODEMANAGER"),
		},
		{
			Name:  "containers_killed_rate",
			Help:  "YARN killed containers per second.",
			Query: roleMetricQuery("containers_killed_rate", yarnServiceTypes, "NODEMANAGER"),
		},
		{
			Name:  "aggregate_containers_allocated_rate",
			Help:  "YARN aggregate allocated containers per second.",
			Query: serviceScopedMetricQuery("aggregate_containers_allocated_rate", yarnServiceTypes),
		},
		{
			Name:  "aggregate_containers_released_rate",
			Help:  "YARN aggregate released containers per second.",
			Query: serviceScopedMetricQuery("aggregate_containers_released_rate", yarnServiceTypes),
		},
		{
			Name:  "used_capacity",
			Help:  "YARN queue used capacity percentage.",
			Query: serviceScopedMetricQuery("used_capacity", yarnServiceTypes),
		},
		{
			Name:  "absolute_used_capacity",
			Help:  "YARN queue absolute used capacity percentage.",
			Query: serviceScopedMetricQuery("absolute_used_capacity", yarnServiceTypes),
		},
		{
			Name:  "fair_share_mb",
			Help:  "YARN queue fair share memory in MB.",
			Query: serviceScopedMetricQuery("fair_share_mb", yarnServiceTypes),
		},
		{
			Name:  "fair_share_vcores",
			Help:  "YARN queue fair share vcores.",
			Query: serviceScopedMetricQuery("fair_share_vcores", yarnServiceTypes),
		},
		{
			Name:  "min_share_mb",
			Help:  "YARN queue minimum share memory in MB.",
			Query: serviceScopedMetricQuery("min_share_mb", yarnServiceTypes),
		},
		{
			Name:  "min_share_vcores",
			Help:  "YARN queue minimum share vcores.",
			Query: serviceScopedMetricQuery("min_share_vcores", yarnServiceTypes),
		},
		{
			Name:  "max_share_mb",
			Help:  "YARN queue maximum share memory in MB.",
			Query: serviceScopedMetricQuery("max_share_mb", yarnServiceTypes),
		},
		{
			Name:  "max_share_vcores",
			Help:  "YARN queue maximum share vcores.",
			Query: serviceScopedMetricQuery("max_share_vcores", yarnServiceTypes),
		},
		{
			Name:  "container_wait_ratio",
			Help:  "YARN queue container wait ratio.",
			Query: serviceScopedMetricQuery("container_wait_ratio", yarnServiceTypes),
		},
		{
			Name:  "get_cluster_metrics_avg_time",
			Help:  "YARN ResourceManager get cluster metrics average time.",
			Query: roleMetricQuery("get_cluster_metrics_avg_time", yarnServiceTypes, "RESOURCEMANAGER"),
		},
		{
			Name:  "get_cluster_nodes_avg_time",
			Help:  "YARN ResourceManager get cluster nodes average time.",
			Query: roleMetricQuery("get_cluster_nodes_avg_time", yarnServiceTypes, "RESOURCEMANAGER"),
		},
		{
			Name:  "rpc_call_queue_length",
			Help:  "YARN RPC call queue length.",
			Query: roleMetricQuery("rpc_call_queue_length", yarnServiceTypes, yarnRoleTypes...),
		},
		{
			Name:  "rpc_num_open_connections",
			Help:  "YARN RPC open connections.",
			Query: roleMetricQuery("rpc_num_open_connections", yarnServiceTypes, yarnRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "YARN role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", yarnServiceTypes, yarnRoleTypes...),
		},
		{
			Name:  "jvm_total_threads",
			Help:  "YARN role JVM thread count.",
			Query: roleMetricQuery("jvm_total_threads", yarnServiceTypes, yarnRoleTypes...),
		},
		{
			Name:  "jvm_gc_rate",
			Help:  "YARN role JVM garbage collections per second.",
			Query: roleMetricQuery("jvm_gc_rate", yarnServiceTypes, yarnRoleTypes...),
		},
		{
			Name:  "fd_open",
			Help:  "YARN role open file descriptors.",
			Query: roleMetricQuery("fd_open", yarnServiceTypes, yarnRoleTypes...),
		},
	}, yarnServiceTypes, yarnRoleTypes...),
}

type ScrapeYARNMetrics struct{}

func (ScrapeYARNMetrics) Name() string {
	return YARN_SCRAPER_NAME
}

func (ScrapeYARNMetrics) Help() string {
	return yarnSpec.Help
}

func (ScrapeYARNMetrics) Version() float64 {
	return 1.0
}

func (ScrapeYARNMetrics) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, yarnSpec, ch)
}

var _ Scraper = ScrapeYARNMetrics{}
