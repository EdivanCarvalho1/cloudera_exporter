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
			Name:  "hive_memory_total_used",
			Help:  "Hive total memory used.",
			Query: roleMetricQuery("hive_memory_total_used", hiveServiceTypes, "HIVESERVER2", "HIVEMETASTORE"),
		},
		{
			Name:  "hive_memory_heap_used",
			Help:  "Hive heap memory used.",
			Query: roleMetricQuery("hive_memory_heap_used", hiveServiceTypes, "HIVESERVER2", "HIVEMETASTORE"),
		},
		{
			Name:  "hive_memory_non_heap_used",
			Help:  "Hive non-heap memory used.",
			Query: roleMetricQuery("hive_memory_non_heap_used", hiveServiceTypes, "HIVESERVER2", "HIVEMETASTORE"),
		},
		{
			Name:  "hive_num_open_transactions",
			Help:  "Hive open transactions.",
			Query: roleMetricQuery("hive_num_open_transactions", hiveServiceTypes, "HIVEMETASTORE"),
		},
		{
			Name:  "hive_num_aborted_transactions",
			Help:  "Hive aborted transactions.",
			Query: roleMetricQuery("hive_num_aborted_transactions", hiveServiceTypes, "HIVEMETASTORE"),
		},
		{
			Name:  "hive_total_num_committed_transactions_rate",
			Help:  "Hive committed transactions per second.",
			Query: roleMetricQuery("hive_total_num_committed_transactions_rate", hiveServiceTypes, "HIVEMETASTORE"),
		},
		{
			Name:  "hive_total_num_aborted_transactions_rate",
			Help:  "Hive aborted transactions per second.",
			Query: roleMetricQuery("hive_total_num_aborted_transactions_rate", hiveServiceTypes, "HIVEMETASTORE"),
		},
		{
			Name:  "hive_api_compile_avg",
			Help:  "Hive average compile time.",
			Query: roleMetricQuery("hive_api_compile_avg", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_api_compile_rate",
			Help:  "Hive compile rate.",
			Query: roleMetricQuery("hive_api_compile_rate", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_api_driver_run_avg",
			Help:  "Hive average driver run time.",
			Query: roleMetricQuery("hive_api_driver_run_avg", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_api_driver_run_rate",
			Help:  "Hive driver run rate.",
			Query: roleMetricQuery("hive_api_driver_run_rate", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_api_operation_running_avg",
			Help:  "Hive average running operations.",
			Query: roleMetricQuery("hive_api_operation_running_avg", hiveServiceTypes, "HIVESERVER2"),
		},
		{
			Name:  "hive_api_operation_pending_avg",
			Help:  "Hive average pending operations.",
			Query: roleMetricQuery("hive_api_operation_pending_avg", hiveServiceTypes, "HIVESERVER2"),
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
