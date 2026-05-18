package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const HBASE_SCRAPER_NAME = "hbase"

var hbaseServiceTypes = []string{"HBASE"}
var hbaseRoleTypes = []string{"MASTER", "REGIONSERVER"}

var hbaseSpec = serviceScraperSpec{
	Name:         HBASE_SCRAPER_NAME,
	Help:         "HBase service metrics",
	ServiceTypes: hbaseServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "read_requests_rate",
			Help:  "HBase read requests per second.",
			Query: rateServiceMetricQuery("read_requests_rate", hbaseServiceTypes),
		},
		{
			Name:  "write_requests_rate",
			Help:  "HBase write requests per second.",
			Query: rateServiceMetricQuery("write_requests_rate", hbaseServiceTypes),
		},
		{
			Name:  "region_count",
			Help:  "HBase region count across RegionServers.",
			Query: serviceMetricQuery("regions", hbaseServiceTypes),
		},
		{
			Name:  "store_file_count",
			Help:  "HBase store file count.",
			Query: roleMetricQuery("store_file_count", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "memstore_size",
			Help:  "HBase memstore size in bytes.",
			Query: roleMetricQuery("memstore_size", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "compaction_queue_size",
			Help:  "HBase compaction queue size.",
			Query: roleMetricQuery("compaction_queue_size", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "flush_queue_size",
			Help:  "HBase flush queue size.",
			Query: roleMetricQuery("flush_queue_size", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "read_latency",
			Help:  "HBase read latency.",
			Query: roleMetricQuery("read_latency", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "write_latency",
			Help:  "HBase write latency.",
			Query: roleMetricQuery("write_latency", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "HBase role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", hbaseServiceTypes, hbaseRoleTypes...),
		},
		{
			Name:  "jvm_threads",
			Help:  "HBase role JVM thread count.",
			Query: roleMetricQuery("jvm_threads", hbaseServiceTypes, hbaseRoleTypes...),
		},
	}, hbaseServiceTypes, hbaseRoleTypes...),
}

type ScrapeHBase struct{}

func (ScrapeHBase) Name() string {
	return HBASE_SCRAPER_NAME
}

func (ScrapeHBase) Help() string {
	return hbaseSpec.Help
}

func (ScrapeHBase) Version() float64 {
	return 1.0
}

func (ScrapeHBase) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, hbaseSpec, ch)
}

var _ Scraper = ScrapeHBase{}
