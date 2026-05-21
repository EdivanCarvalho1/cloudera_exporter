package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const KUDU_SCRAPER_NAME = "kudu"

var kuduServiceTypes = []string{"KUDU"}
var kuduRoleTypes = []string{"KUDU_MASTER", "KUDU_TSERVER"}

var kuduSpec = serviceScraperSpec{
	Name:         KUDU_SCRAPER_NAME,
	Help:         "Kudu service metrics",
	ServiceTypes: kuduServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "kudu_live_row_count",
			Help:  "Kudu live row count.",
			Query: roleMetricQuery("kudu_live_row_count", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_on_disk_data_size",
			Help:  "Kudu on-disk data size in bytes.",
			Query: roleMetricQuery("kudu_on_disk_data_size", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_rows_inserted_rate",
			Help:  "Kudu inserted rows per second.",
			Query: roleMetricQuery("kudu_rows_inserted_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_rows_updated_rate",
			Help:  "Kudu updated rows per second.",
			Query: roleMetricQuery("kudu_rows_updated_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_rows_deleted_rate",
			Help:  "Kudu deleted rows per second.",
			Query: roleMetricQuery("kudu_rows_deleted_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_scans_started_rate",
			Help:  "Kudu scans started per second.",
			Query: roleMetricQuery("kudu_scans_started_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_handler_latency_kudu_tserver_tabletserverservice_write_rate",
			Help:  "Kudu tablet server write handler latency rate.",
			Query: roleMetricQuery("kudu_handler_latency_kudu_tserver_tabletserverservice_write_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_handler_latency_kudu_tserver_tabletserverservice_scan_rate",
			Help:  "Kudu tablet server scan handler latency rate.",
			Query: roleMetricQuery("kudu_handler_latency_kudu_tserver_tabletserverservice_scan_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "kudu_rpcs_queue_overflow_rate",
			Help:  "Kudu RPC queue overflow rate.",
			Query: roleMetricQuery("kudu_rpcs_queue_overflow_rate", kuduServiceTypes, kuduRoleTypes...),
		},
		{
			Name:  "kudu_rpc_incoming_queue_time_rate",
			Help:  "Kudu incoming RPC queue time rate.",
			Query: roleMetricQuery("kudu_rpc_incoming_queue_time_rate", kuduServiceTypes, kuduRoleTypes...),
		},
		{
			Name:  "kudu_memory_usage",
			Help:  "Kudu memory usage in bytes.",
			Query: roleMetricQuery("kudu_memory_usage", kuduServiceTypes, kuduRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Kudu role JVM heap used in MB when exposed by Cloudera Manager.",
			Query: roleMetricQuery("jvm_heap_used_mb", kuduServiceTypes, kuduRoleTypes...),
		},
	}, kuduServiceTypes, kuduRoleTypes...),
}

type ScrapeKudu struct{}

func (ScrapeKudu) Name() string {
	return KUDU_SCRAPER_NAME
}

func (ScrapeKudu) Help() string {
	return kuduSpec.Help
}

func (ScrapeKudu) Version() float64 {
	return 1.0
}

func (ScrapeKudu) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, kuduSpec, ch)
}

var _ Scraper = ScrapeKudu{}
