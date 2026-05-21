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
			Name:  "requests_rate",
			Help:  "HBase requests per second.",
			Query: roleMetricQuery("requests_rate", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "region_count",
			Help:  "HBase region count across RegionServers.",
			Query: serviceMetricQuery("regions", hbaseServiceTypes),
		},
		{
			Name:  "stores",
			Help:  "HBase store count.",
			Query: roleMetricQuery("stores", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "storefiles",
			Help:  "HBase store file count.",
			Query: roleMetricQuery("storefiles", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "block_cache_hit_ratio",
			Help:  "HBase block cache hit ratio.",
			Query: roleMetricQuery("block_cache_hit_ratio", hbaseServiceTypes, "REGIONSERVER"),
		},
		{
			Name:  "block_cache_evicted_rate",
			Help:  "HBase block cache eviction rate.",
			Query: roleMetricQuery("block_cache_evicted_rate", hbaseServiceTypes, "REGIONSERVER"),
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
			Name:  "jvm_heap_used_mb",
			Help:  "HBase role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", hbaseServiceTypes, hbaseRoleTypes...),
		},
		{
			Name:  "jvm_total_threads",
			Help:  "HBase role JVM thread count.",
			Query: roleMetricQuery("jvm_total_threads", hbaseServiceTypes, hbaseRoleTypes...),
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
