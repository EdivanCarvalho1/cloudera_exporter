package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const HUE_SCRAPER_NAME = "hue"

var hueServiceTypes = []string{"HUE"}

var hueSpec = serviceScraperSpec{
	Name:         HUE_SCRAPER_NAME,
	Help:         "Hue service metrics",
	ServiceTypes: hueServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "hue_users",
			Help:  "Hue users.",
			Query: roleMetricQuery("hue_users", hueServiceTypes),
		},
		{
			Name:  "hue_users_active",
			Help:  "Hue active users.",
			Query: roleMetricQuery("hue_users_active", hueServiceTypes),
		},
		{
			Name:  "hue_requests_active",
			Help:  "Hue active requests.",
			Query: roleMetricQuery("hue_requests_active", hueServiceTypes),
		},
		{
			Name:  "hue_requests_exceptions_rate",
			Help:  "Hue request exception rate.",
			Query: roleMetricQuery("hue_requests_exceptions_rate", hueServiceTypes),
		},
		{
			Name:  "hue_requests_response_time_avg",
			Help:  "Hue average request response time.",
			Query: roleMetricQuery("hue_requests_response_time_avg", hueServiceTypes),
		},
		{
			Name:  "hue_requests_response_time_max",
			Help:  "Hue maximum request response time.",
			Query: roleMetricQuery("hue_requests_response_time_max", hueServiceTypes),
		},
		{
			Name:  "hue_threads_total",
			Help:  "Hue total threads.",
			Query: roleMetricQuery("hue_threads_total", hueServiceTypes),
		},
		{
			Name:  "hue_threads_daemon",
			Help:  "Hue daemon threads.",
			Query: roleMetricQuery("hue_threads_daemon", hueServiceTypes),
		},
		{
			Name:  "hue_python_gc_objects",
			Help:  "Hue Python garbage collector objects.",
			Query: roleMetricQuery("hue_python_gc_objects", hueServiceTypes),
		},
		{
			Name:  "hue_multiprocessing_processes_total",
			Help:  "Hue multiprocessing total processes.",
			Query: roleMetricQuery("hue_multiprocessing_processes_total", hueServiceTypes),
		},
	}, hueServiceTypes),
}

type ScrapeHue struct{}

func (ScrapeHue) Name() string {
	return HUE_SCRAPER_NAME
}

func (ScrapeHue) Help() string {
	return hueSpec.Help
}

func (ScrapeHue) Version() float64 {
	return 1.0
}

func (ScrapeHue) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, hueSpec, ch)
}

var _ Scraper = ScrapeHue{}
