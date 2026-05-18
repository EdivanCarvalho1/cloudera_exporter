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
			Name:  "tablet_count",
			Help:  "Kudu tablet count.",
			Query: roleMetricQuery("tablet_count", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "leader_tablet_count",
			Help:  "Kudu leader tablet count.",
			Query: roleMetricQuery("leader_tablet_count", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "replica_count",
			Help:  "Kudu tablet replica count.",
			Query: roleMetricQuery("replica_count", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "write_ops_rate",
			Help:  "Kudu write operations per second.",
			Query: rateRoleMetricQuery("write_ops_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "read_ops_rate",
			Help:  "Kudu read operations per second.",
			Query: rateRoleMetricQuery("read_ops_rate", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "write_latency",
			Help:  "Kudu write latency.",
			Query: roleMetricQuery("write_latency", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "read_latency",
			Help:  "Kudu read latency.",
			Query: roleMetricQuery("read_latency", kuduServiceTypes, "KUDU_TSERVER"),
		},
		{
			Name:  "rpc_queue_length",
			Help:  "Kudu RPC queue length.",
			Query: roleMetricQuery("rpc_queue_length", kuduServiceTypes, kuduRoleTypes...),
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
