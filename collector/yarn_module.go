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
			Name:  "jvm_heap_used_mb",
			Help:  "YARN role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", yarnServiceTypes, yarnRoleTypes...),
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
