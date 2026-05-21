package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const KNOX_SCRAPER_NAME = "knox"

var knoxServiceTypes = []string{"KNOX"}

var knoxSpec = serviceScraperSpec{
	Name:         KNOX_SCRAPER_NAME,
	Help:         "Knox service metrics",
	ServiceTypes: knoxServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Knox JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", knoxServiceTypes),
		},
		{
			Name:  "jvm_total_threads",
			Help:  "Knox JVM thread count.",
			Query: roleMetricQuery("jvm_total_threads", knoxServiceTypes),
		},
		{
			Name:  "fd_open",
			Help:  "Knox open file descriptors.",
			Query: roleMetricQuery("fd_open", knoxServiceTypes),
		},
		{
			Name:  "fd_max",
			Help:  "Knox maximum file descriptors.",
			Query: roleMetricQuery("fd_max", knoxServiceTypes),
		},
		{
			Name:  "uptime",
			Help:  "Knox uptime in seconds.",
			Query: roleMetricQuery("uptime", knoxServiceTypes),
		},
	}, knoxServiceTypes),
}

type ScrapeKnox struct{}

func (ScrapeKnox) Name() string {
	return KNOX_SCRAPER_NAME
}

func (ScrapeKnox) Help() string {
	return knoxSpec.Help
}

func (ScrapeKnox) Version() float64 {
	return 1.0
}

func (ScrapeKnox) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, knoxSpec, ch)
}

var _ Scraper = ScrapeKnox{}
