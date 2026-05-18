package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const FLINK_SCRAPER_NAME = "flink"

var flinkServiceTypes = []string{"FLINK"}
var flinkRoleTypes = []string{"JOBMANAGER", "TASKMANAGER", "FLINK_HISTORY_SERVER", "GATEWAY"}

var flinkSpec = serviceScraperSpec{
	Name:         FLINK_SCRAPER_NAME,
	Help:         "Flink service metrics",
	ServiceTypes: flinkServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "jobs_running",
			Help:  "Flink running jobs.",
			Query: serviceMetricQuery("jobs_running", flinkServiceTypes),
		},
		{
			Name:  "jobs_failed",
			Help:  "Flink failed jobs.",
			Query: serviceMetricQuery("jobs_failed", flinkServiceTypes),
		},
		{
			Name:  "jobs_finished",
			Help:  "Flink finished jobs.",
			Query: serviceMetricQuery("jobs_finished", flinkServiceTypes),
		},
		{
			Name:  "task_slots_total",
			Help:  "Flink total task slots.",
			Query: serviceMetricQuery("task_slots_total", flinkServiceTypes),
		},
		{
			Name:  "task_slots_available",
			Help:  "Flink available task slots.",
			Query: serviceMetricQuery("task_slots_available", flinkServiceTypes),
		},
		{
			Name:  "checkpoints_completed",
			Help:  "Flink completed checkpoints.",
			Query: serviceMetricQuery("checkpoints_completed", flinkServiceTypes),
		},
		{
			Name:  "checkpoints_failed",
			Help:  "Flink failed checkpoints.",
			Query: serviceMetricQuery("checkpoints_failed", flinkServiceTypes),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Flink role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", flinkServiceTypes, flinkRoleTypes...),
		},
	}, flinkServiceTypes, flinkRoleTypes...),
}

type ScrapeFlink struct{}

func (ScrapeFlink) Name() string {
	return FLINK_SCRAPER_NAME
}

func (ScrapeFlink) Help() string {
	return flinkSpec.Help
}

func (ScrapeFlink) Version() float64 {
	return 1.0
}

func (ScrapeFlink) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, flinkSpec, ch)
}

var _ Scraper = ScrapeFlink{}
