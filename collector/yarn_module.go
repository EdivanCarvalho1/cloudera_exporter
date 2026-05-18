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
			Name:  "apps_failed",
			Help:  "YARN failed applications.",
			Query: serviceMetricQuery("apps_failed", yarnServiceTypes),
		},
		{
			Name:  "apps_killed",
			Help:  "YARN killed applications.",
			Query: serviceMetricQuery("apps_killed", yarnServiceTypes),
		},
		{
			Name:  "containers_allocated",
			Help:  "YARN allocated containers.",
			Query: serviceMetricQuery("containers_allocated", yarnServiceTypes),
		},
		{
			Name:  "containers_pending",
			Help:  "YARN pending containers.",
			Query: serviceMetricQuery("containers_pending", yarnServiceTypes),
		},
		{
			Name:  "total_memory_mb",
			Help:  "YARN total memory in MB.",
			Query: serviceMetricQuery("total_memory_mb", yarnServiceTypes),
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
			Name:  "total_vcores",
			Help:  "YARN total vcores.",
			Query: serviceMetricQuery("total_vcores", yarnServiceTypes),
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
			Name:  "active_nodemanagers",
			Help:  "YARN active NodeManagers.",
			Query: serviceMetricQuery("active_nodemanagers", yarnServiceTypes),
		},
		{
			Name:  "lost_nodemanagers",
			Help:  "YARN lost NodeManagers.",
			Query: serviceMetricQuery("lost_nodemanagers", yarnServiceTypes),
		},
		{
			Name:  "unhealthy_nodemanagers",
			Help:  "YARN unhealthy NodeManagers.",
			Query: serviceMetricQuery("unhealthy_nodemanagers", yarnServiceTypes),
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
