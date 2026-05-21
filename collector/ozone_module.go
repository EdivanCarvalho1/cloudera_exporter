package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const OZONE_SCRAPER_NAME = "ozone"

var ozoneServiceTypes = []string{"OZONE"}

var ozoneSpec = serviceScraperSpec{
	Name:         OZONE_SCRAPER_NAME,
	Help:         "Ozone service metrics",
	ServiceTypes: ozoneServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "ozone_total_capacity",
			Help:  "Ozone total capacity.",
			Query: roleMetricQuery("ozone_total_capacity", ozoneServiceTypes),
		},
		{
			Name:  "ozone_total_used",
			Help:  "Ozone total used capacity.",
			Query: roleMetricQuery("ozone_total_used", ozoneServiceTypes),
		},
		{
			Name:  "ozone_datanodes",
			Help:  "Ozone datanode count.",
			Query: roleMetricQuery("ozone_datanodes", ozoneServiceTypes),
		},
		{
			Name:  "ozone_om_num_volumes",
			Help:  "Ozone Manager volume count.",
			Query: roleMetricQuery("ozone_om_num_volumes", ozoneServiceTypes),
		},
		{
			Name:  "ozone_om_num_buckets",
			Help:  "Ozone Manager bucket count.",
			Query: roleMetricQuery("ozone_om_num_buckets", ozoneServiceTypes),
		},
		{
			Name:  "ozone_om_num_keys",
			Help:  "Ozone Manager key count.",
			Query: roleMetricQuery("ozone_om_num_keys", ozoneServiceTypes),
		},
		{
			Name:  "ozone_om_num_open_connections",
			Help:  "Ozone Manager open connections.",
			Query: roleMetricQuery("ozone_om_num_open_connections", ozoneServiceTypes),
		},
		{
			Name:  "ozone_om_call_queue_length",
			Help:  "Ozone Manager call queue length.",
			Query: roleMetricQuery("ozone_om_call_queue_length", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_healthy_nodes",
			Help:  "Ozone SCM healthy nodes.",
			Query: roleMetricQuery("ozone_scm_healthy_nodes", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_stale_nodes",
			Help:  "Ozone SCM stale nodes.",
			Query: roleMetricQuery("ozone_scm_stale_nodes", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_dead_nodes",
			Help:  "Ozone SCM dead nodes.",
			Query: roleMetricQuery("ozone_scm_dead_nodes", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_num_current_healthy_pipeline",
			Help:  "Ozone SCM current healthy pipelines.",
			Query: roleMetricQuery("ozone_scm_num_current_healthy_pipeline", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_num_open_connections",
			Help:  "Ozone SCM open connections.",
			Query: roleMetricQuery("ozone_scm_num_open_connections", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_java_heap_memory_used",
			Help:  "Ozone SCM Java heap memory used.",
			Query: roleMetricQuery("ozone_scm_java_heap_memory_used", ozoneServiceTypes),
		},
		{
			Name:  "ozone_scm_call_queue_length",
			Help:  "Ozone SCM call queue length.",
			Query: roleMetricQuery("ozone_scm_call_queue_length", ozoneServiceTypes),
		},
	}, ozoneServiceTypes),
}

type ScrapeOzone struct{}

func (ScrapeOzone) Name() string {
	return OZONE_SCRAPER_NAME
}

func (ScrapeOzone) Help() string {
	return ozoneSpec.Help
}

func (ScrapeOzone) Version() float64 {
	return 1.0
}

func (ScrapeOzone) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, ozoneSpec, ch)
}

var _ Scraper = ScrapeOzone{}
