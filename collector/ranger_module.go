package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const RANGER_SCRAPER_NAME = "ranger"

var rangerServiceTypes = []string{"RANGER", "RANGER_KMS", "RANGER_RMS", "RANGER_RAZ"}

var rangerSpec = serviceScraperSpec{
	Name:         RANGER_SCRAPER_NAME,
	Help:         "Ranger service metrics",
	ServiceTypes: rangerServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "ranger_admin_memory_heap_used",
			Help:  "Ranger Admin heap memory used.",
			Query: roleMetricQuery("ranger_admin_memory_heap_used", rangerServiceTypes),
		},
		{
			Name:  "ranger_admin_memory_non_heap_used",
			Help:  "Ranger Admin non-heap memory used.",
			Query: roleMetricQuery("ranger_admin_memory_non_heap_used", rangerServiceTypes),
		},
		{
			Name:  "ranger_kms_jvm_used_memory",
			Help:  "Ranger KMS JVM used memory.",
			Query: roleMetricQuery("ranger_kms_jvm_used_memory", rangerServiceTypes),
		},
		{
			Name:  "ranger_kms_total_call_count",
			Help:  "Ranger KMS total call count.",
			Query: roleMetricQuery("ranger_kms_total_call_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_kms_unauthorized_call_count",
			Help:  "Ranger KMS unauthorized call count.",
			Query: roleMetricQuery("ranger_kms_unauthorized_call_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_kms_unauthenticated_call_count",
			Help:  "Ranger KMS unauthenticated call count.",
			Query: roleMetricQuery("ranger_kms_unauthenticated_call_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_kms_key_count",
			Help:  "Ranger KMS key count.",
			Query: roleMetricQuery("ranger_kms_key_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_rms_jvm_used_memory",
			Help:  "Ranger RMS JVM used memory.",
			Query: roleMetricQuery("ranger_rms_jvm_used_memory", rangerServiceTypes),
		},
		{
			Name:  "ranger_rms_total_processed_notification_count",
			Help:  "Ranger RMS processed notification count.",
			Query: roleMetricQuery("ranger_rms_total_processed_notification_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_raz_jvm_used_memory",
			Help:  "Ranger RAZ JVM used memory.",
			Query: roleMetricQuery("ranger_raz_jvm_used_memory", rangerServiceTypes),
		},
		{
			Name:  "ranger_raz_requests_processed_count",
			Help:  "Ranger RAZ processed request count.",
			Query: roleMetricQuery("ranger_raz_requests_processed_count", rangerServiceTypes),
		},
		{
			Name:  "ranger_raz_requests_failed_count",
			Help:  "Ranger RAZ failed request count.",
			Query: roleMetricQuery("ranger_raz_requests_failed_count", rangerServiceTypes),
		},
	}, rangerServiceTypes),
}

type ScrapeRanger struct{}

func (ScrapeRanger) Name() string {
	return RANGER_SCRAPER_NAME
}

func (ScrapeRanger) Help() string {
	return rangerSpec.Help
}

func (ScrapeRanger) Version() float64 {
	return 1.0
}

func (ScrapeRanger) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, rangerSpec, ch)
}

var _ Scraper = ScrapeRanger{}
