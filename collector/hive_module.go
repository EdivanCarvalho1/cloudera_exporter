package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const HIVE_SCRAPER_NAME = "hive"

var hiveServiceTypes = []string{"HIVE", "HIVE_ON_TEZ"}
var hiveRoleTypes = []string{"HIVESERVER2", "HIVEMETASTORE", "GATEWAY"}

var hiveSpec = serviceScraperSpec{
	Name:         HIVE_SCRAPER_NAME,
	Help:         "Hive service metrics",
	ServiceTypes: hiveServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "hive_open_connections",
			Help:  "Hive open connections.",
			Query: roleMetricQuery("hive_open_connections", hiveServiceTypes, "HIVESERVER2", "HIVEMETASTORE"),
		},
		{
			Name:  "hive_on_tez_open_connections",
			Help:  "Hive on Tez open connections.",
			Query: roleMetricQuery("hive_on_tez_open_connections", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_open_operations",
			Help:  "Hive open operations.",
			Query: roleMetricQuery("hive_open_operations", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_on_tez_open_operations",
			Help:  "Hive on Tez open operations.",
			Query: roleMetricQuery("hive_on_tez_open_operations", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_completed_operation_closed_rate",
			Help:  "Hive completed operations closed per second.",
			Query: roleMetricQuery("hive_completed_operation_closed_rate", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Hive role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", hiveServiceTypes, hiveRoleTypes...),
		},
		{
			Name:  "hive_threads_thread_count",
			Help:  "Hive role thread count.",
			Query: roleMetricQuery("hive_threads_thread_count", hiveServiceTypes, hiveRoleTypes...),
		},
		{
			Name:  "hive_on_tez_threads_thread_count",
			Help:  "Hive on Tez role thread count.",
			Query: roleMetricQuery("hive_on_tez_threads_thread_count", hiveServiceTypes, hiveRoleTypes...),
		},
	}, hiveServiceTypes, hiveRoleTypes...),
}

type ScrapeHive struct{}

func (ScrapeHive) Name() string {
	return HIVE_SCRAPER_NAME
}

func (ScrapeHive) Help() string {
	return hiveSpec.Help
}

func (ScrapeHive) Version() float64 {
	return 1.0
}

func (ScrapeHive) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, hiveSpec, ch)
}

var _ Scraper = ScrapeHive{}
