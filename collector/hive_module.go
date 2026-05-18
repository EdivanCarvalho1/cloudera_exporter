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
			Name:  "open_connections",
			Help:  "Hive open connections.",
			Query: roleMetricQuery("open_connections", hiveServiceTypes, "HIVESERVER2", "HIVEMETASTORE"),
		},
		{
			Name:  "open_sessions",
			Help:  "HiveServer2 open sessions.",
			Query: roleMetricQuery("open_sessions", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "active_operations",
			Help:  "HiveServer2 active operations.",
			Query: roleMetricQuery("active_operations", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "completed_operations",
			Help:  "HiveServer2 completed operations.",
			Query: roleMetricQuery("completed_operations", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Hive role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", hiveServiceTypes, hiveRoleTypes...),
		},
		{
			Name:  "jvm_threads",
			Help:  "Hive role JVM thread count.",
			Query: roleMetricQuery("jvm_threads", hiveServiceTypes, hiveRoleTypes...),
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
