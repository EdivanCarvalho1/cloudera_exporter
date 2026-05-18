package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const NIFI_SCRAPER_NAME = "nifi"

var nifiServiceTypes = []string{"NIFI"}
var nifiRoleTypes = []string{"NIFI_NODE", "NIFI", "NODE"}

var nifiSpec = serviceScraperSpec{
	Name:         NIFI_SCRAPER_NAME,
	Help:         "NiFi service metrics",
	ServiceTypes: nifiServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "flow_files_queued",
			Help:  "NiFi queued flow files.",
			Query: roleMetricQuery("flow_files_queued", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "bytes_queued",
			Help:  "NiFi queued bytes.",
			Query: roleMetricQuery("bytes_queued", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "bytes_read_rate",
			Help:  "NiFi bytes read per second.",
			Query: rateRoleMetricQuery("bytes_read_rate", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "bytes_written_rate",
			Help:  "NiFi bytes written per second.",
			Query: rateRoleMetricQuery("bytes_written_rate", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "bytes_sent_rate",
			Help:  "NiFi bytes sent per second.",
			Query: rateRoleMetricQuery("bytes_sent_rate", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "bytes_received_rate",
			Help:  "NiFi bytes received per second.",
			Query: rateRoleMetricQuery("bytes_received_rate", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "active_threads",
			Help:  "NiFi active threads.",
			Query: roleMetricQuery("active_threads", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "garbage_collection_rate",
			Help:  "NiFi garbage collections per second.",
			Query: rateRoleMetricQuery("garbage_collection_rate", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "NiFi JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", nifiServiceTypes, nifiRoleTypes...),
		},
	}, nifiServiceTypes, nifiRoleTypes...),
}

type ScrapeNiFi struct{}

func (ScrapeNiFi) Name() string {
	return NIFI_SCRAPER_NAME
}

func (ScrapeNiFi) Help() string {
	return nifiSpec.Help
}

func (ScrapeNiFi) Version() float64 {
	return 1.0
}

func (ScrapeNiFi) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, nifiSpec, ch)
}

var _ Scraper = ScrapeNiFi{}
