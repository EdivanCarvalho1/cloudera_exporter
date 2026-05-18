package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const SPARK_SCRAPER_NAME = "spark"

var sparkServiceTypes = []string{"SPARK_ON_YARN", "SPARK3_ON_YARN", "SPARK"}
var sparkRoleTypes = []string{"SPARK_YARN_HISTORY_SERVER", "SPARK3_YARN_HISTORY_SERVER", "GATEWAY"}

var sparkSpec = serviceScraperSpec{
	Name:         SPARK_SCRAPER_NAME,
	Help:         "Spark service metrics",
	ServiceTypes: sparkServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "apps_running",
			Help:  "Spark applications currently running.",
			Query: serviceMetricQuery("apps_running", sparkServiceTypes),
		},
		{
			Name:  "apps_completed",
			Help:  "Spark applications completed.",
			Query: serviceMetricQuery("apps_completed", sparkServiceTypes),
		},
		{
			Name:  "apps_failed",
			Help:  "Spark applications failed.",
			Query: serviceMetricQuery("apps_failed", sparkServiceTypes),
		},
		{
			Name:  "executors_active",
			Help:  "Spark active executors.",
			Query: serviceMetricQuery("executors_active", sparkServiceTypes),
		},
		{
			Name:  "drivers_active",
			Help:  "Spark active drivers.",
			Query: serviceMetricQuery("drivers_active", sparkServiceTypes),
		},
		{
			Name:  "memory_used",
			Help:  "Spark memory used.",
			Query: serviceMetricQuery("memory_used", sparkServiceTypes),
		},
		{
			Name:  "cores_used",
			Help:  "Spark cores used.",
			Query: serviceMetricQuery("cores_used", sparkServiceTypes),
		},
		{
			Name:  "history_server_jvm_heap_used_mb",
			Help:  "Spark History Server JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", sparkServiceTypes, "SPARK_YARN_HISTORY_SERVER", "SPARK3_YARN_HISTORY_SERVER"),
		},
	}, sparkServiceTypes, sparkRoleTypes...),
}

type ScrapeSpark struct{}

func (ScrapeSpark) Name() string {
	return SPARK_SCRAPER_NAME
}

func (ScrapeSpark) Help() string {
	return sparkSpec.Help
}

func (ScrapeSpark) Version() float64 {
	return 1.0
}

func (ScrapeSpark) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, sparkSpec, ch)
}

var _ Scraper = ScrapeSpark{}
