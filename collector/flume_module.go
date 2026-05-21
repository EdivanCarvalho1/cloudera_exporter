package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const FLUME_SCRAPER_NAME = "flume"

var flumeServiceTypes = []string{"FLUME"}

var flumeSpec = serviceScraperSpec{
	Name:         FLUME_SCRAPER_NAME,
	Help:         "Flume service metrics",
	ServiceTypes: flumeServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "channel_size",
			Help:  "Flume channel size.",
			Query: roleMetricQuery("channel_size", flumeServiceTypes),
		},
		{
			Name:  "channel_capacity",
			Help:  "Flume channel capacity.",
			Query: roleMetricQuery("channel_capacity", flumeServiceTypes),
		},
		{
			Name:  "event_received_rate",
			Help:  "Flume source event received rate.",
			Query: roleMetricQuery("event_received_rate", flumeServiceTypes),
		},
		{
			Name:  "event_accepted_rate",
			Help:  "Flume source event accepted rate.",
			Query: roleMetricQuery("event_accepted_rate", flumeServiceTypes),
		},
		{
			Name:  "event_put_attempt_rate",
			Help:  "Flume channel event put attempt rate.",
			Query: roleMetricQuery("event_put_attempt_rate", flumeServiceTypes),
		},
		{
			Name:  "event_put_success_rate",
			Help:  "Flume channel event put success rate.",
			Query: roleMetricQuery("event_put_success_rate", flumeServiceTypes),
		},
		{
			Name:  "event_take_attempt_rate",
			Help:  "Flume channel event take attempt rate.",
			Query: roleMetricQuery("event_take_attempt_rate", flumeServiceTypes),
		},
		{
			Name:  "event_take_success_rate",
			Help:  "Flume channel event take success rate.",
			Query: roleMetricQuery("event_take_success_rate", flumeServiceTypes),
		},
		{
			Name:  "event_drain_success_rate",
			Help:  "Flume sink event drain success rate.",
			Query: roleMetricQuery("event_drain_success_rate", flumeServiceTypes),
		},
		{
			Name:  "event_drain_attempt_rate",
			Help:  "Flume sink event drain attempt rate.",
			Query: roleMetricQuery("event_drain_attempt_rate", flumeServiceTypes),
		},
		{
			Name:  "batch_complete_rate",
			Help:  "Flume sink batch complete rate.",
			Query: roleMetricQuery("batch_complete_rate", flumeServiceTypes),
		},
		{
			Name:  "batch_empty_rate",
			Help:  "Flume sink batch empty rate.",
			Query: roleMetricQuery("batch_empty_rate", flumeServiceTypes),
		},
		{
			Name:  "batch_underflow_rate",
			Help:  "Flume sink batch underflow rate.",
			Query: roleMetricQuery("batch_underflow_rate", flumeServiceTypes),
		},
		{
			Name:  "connection_failed_rate",
			Help:  "Flume sink connection failure rate.",
			Query: roleMetricQuery("connection_failed_rate", flumeServiceTypes),
		},
	}, flumeServiceTypes),
}

type ScrapeFlume struct{}

func (ScrapeFlume) Name() string {
	return FLUME_SCRAPER_NAME
}

func (ScrapeFlume) Help() string {
	return flumeSpec.Help
}

func (ScrapeFlume) Version() float64 {
	return 1.0
}

func (ScrapeFlume) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, flumeSpec, ch)
}

var _ Scraper = ScrapeFlume{}
