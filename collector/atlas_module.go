package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const ATLAS_SCRAPER_NAME = "atlas"

var atlasServiceTypes = []string{"ATLAS"}

var atlasSpec = serviceScraperSpec{
	Name:         ATLAS_SCRAPER_NAME,
	Help:         "Atlas service metrics",
	ServiceTypes: atlasServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "atlas_server_memory_heap_used",
			Help:  "Atlas server heap memory used.",
			Query: roleMetricQuery("atlas_server_memory_heap_used", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_memory_heap_committed",
			Help:  "Atlas server heap memory committed.",
			Query: roleMetricQuery("atlas_server_memory_heap_committed", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_memory_heap_max",
			Help:  "Atlas server heap memory max.",
			Query: roleMetricQuery("atlas_server_memory_heap_max", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_memory_non_heap_used",
			Help:  "Atlas server non-heap memory used.",
			Query: roleMetricQuery("atlas_server_memory_non_heap_used", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_hour_entity_creates_rate",
			Help:  "Atlas current-hour entity create rate.",
			Query: roleMetricQuery("atlas_server_current_hour_entity_creates_rate", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_hour_entity_updates_rate",
			Help:  "Atlas current-hour entity update rate.",
			Query: roleMetricQuery("atlas_server_current_hour_entity_updates_rate", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_hour_entity_deletes_rate",
			Help:  "Atlas current-hour entity delete rate.",
			Query: roleMetricQuery("atlas_server_current_hour_entity_deletes_rate", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_day_entity_creates_rate",
			Help:  "Atlas current-day entity create rate.",
			Query: roleMetricQuery("atlas_server_current_day_entity_creates_rate", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_day_entity_updates_rate",
			Help:  "Atlas current-day entity update rate.",
			Query: roleMetricQuery("atlas_server_current_day_entity_updates_rate", atlasServiceTypes),
		},
		{
			Name:  "atlas_server_current_day_entity_deletes_rate",
			Help:  "Atlas current-day entity delete rate.",
			Query: roleMetricQuery("atlas_server_current_day_entity_deletes_rate", atlasServiceTypes),
		},
	}, atlasServiceTypes),
}

type ScrapeAtlas struct{}

func (ScrapeAtlas) Name() string {
	return ATLAS_SCRAPER_NAME
}

func (ScrapeAtlas) Help() string {
	return atlasSpec.Help
}

func (ScrapeAtlas) Version() float64 {
	return 1.0
}

func (ScrapeAtlas) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, atlasSpec, ch)
}

var _ Scraper = ScrapeAtlas{}
