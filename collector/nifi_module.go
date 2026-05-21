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
			Name:  "nifi_amount_items_queued",
			Help:  "NiFi queued flow files.",
			Query: roleMetricQuery("nifi_amount_items_queued", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_size_content_queued_total",
			Help:  "NiFi queued bytes.",
			Query: roleMetricQuery("nifi_size_content_queued_total", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_total_bytes_read",
			Help:  "NiFi total bytes read.",
			Query: roleMetricQuery("nifi_total_bytes_read", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_total_bytes_written",
			Help:  "NiFi total bytes written.",
			Query: roleMetricQuery("nifi_total_bytes_written", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_total_bytes_sent",
			Help:  "NiFi total bytes sent.",
			Query: roleMetricQuery("nifi_total_bytes_sent", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_total_bytes_received",
			Help:  "NiFi total bytes received.",
			Query: roleMetricQuery("nifi_total_bytes_received", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_amount_threads_active",
			Help:  "NiFi active threads.",
			Query: roleMetricQuery("nifi_amount_threads_active", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_amount_flowfiles_receive",
			Help:  "NiFi flow files received.",
			Query: roleMetricQuery("nifi_amount_flowfiles_receive", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_amount_flowfiles_sent",
			Help:  "NiFi flow files sent.",
			Query: roleMetricQuery("nifi_amount_flowfiles_sent", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_percent_used_bytes",
			Help:  "NiFi repository bytes used percentage.",
			Query: roleMetricQuery("nifi_percent_used_bytes", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_percent_used_count",
			Help:  "NiFi repository count used percentage.",
			Query: roleMetricQuery("nifi_percent_used_count", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_time_to_bytes_backpressure_prediction",
			Help:  "NiFi time to bytes backpressure prediction.",
			Query: roleMetricQuery("nifi_time_to_bytes_backpressure_prediction", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_time_to_count_backpressure_prediction",
			Help:  "NiFi time to count backpressure prediction.",
			Query: roleMetricQuery("nifi_time_to_count_backpressure_prediction", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_node_connectivity",
			Help:  "NiFi node connectivity status.",
			Query: roleMetricQuery("nifi_node_connectivity", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_open_files_count_check",
			Help:  "NiFi open files count check.",
			Query: roleMetricQuery("nifi_open_files_count_check", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_jvm_thread_count",
			Help:  "NiFi JVM thread count.",
			Query: roleMetricQuery("nifi_jvm_thread_count", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_jvm_heap_used",
			Help:  "NiFi JVM heap used in MB.",
			Query: roleMetricQuery("nifi_jvm_heap_used", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_flowfile_repo_usage",
			Help:  "NiFi FlowFile repository usage.",
			Query: roleMetricQuery("nifi_flowfile_repo_usage", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_content_repo_usage",
			Help:  "NiFi content repository usage.",
			Query: roleMetricQuery("nifi_content_repo_usage", nifiServiceTypes, nifiRoleTypes...),
		},
		{
			Name:  "nifi_provenance_repo_usage",
			Help:  "NiFi provenance repository usage.",
			Query: roleMetricQuery("nifi_provenance_repo_usage", nifiServiceTypes, nifiRoleTypes...),
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
