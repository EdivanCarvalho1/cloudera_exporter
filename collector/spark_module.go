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
			Name:  "spark_on_yarn_event_log_data_export_fail_counts_rate",
			Help:  "Spark event log export failure rate.",
			Query: serviceMetricQuery("spark_on_yarn_event_log_data_export_fail_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_event_log_data_ingest_fail_counts_rate",
			Help:  "Spark event log ingest failure rate.",
			Query: serviceMetricQuery("spark_on_yarn_event_log_data_ingest_fail_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_event_log_data_export_success_counts_rate",
			Help:  "Spark event log export success rate.",
			Query: serviceMetricQuery("spark_on_yarn_event_log_data_export_success_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_event_log_data_ingest_success_counts_rate",
			Help:  "Spark event log ingest success rate.",
			Query: serviceMetricQuery("spark_on_yarn_event_log_data_ingest_success_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_lineage_data_export_fail_counts_rate",
			Help:  "Spark lineage data export failure rate.",
			Query: serviceMetricQuery("spark_on_yarn_lineage_data_export_fail_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_lineage_data_ingest_fail_counts_rate",
			Help:  "Spark lineage data ingest failure rate.",
			Query: serviceMetricQuery("spark_on_yarn_lineage_data_ingest_fail_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_lineage_data_export_success_counts_rate",
			Help:  "Spark lineage data export success rate.",
			Query: serviceMetricQuery("spark_on_yarn_lineage_data_export_success_counts_rate", sparkServiceTypes),
		},
		{
			Name:  "spark_on_yarn_lineage_data_ingest_success_counts_rate",
			Help:  "Spark lineage data ingest success rate.",
			Query: serviceMetricQuery("spark_on_yarn_lineage_data_ingest_success_counts_rate", sparkServiceTypes),
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
