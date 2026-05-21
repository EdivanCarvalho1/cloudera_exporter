package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const FLINK_SCRAPER_NAME = "flink"

var flinkServiceTypes = []string{"FLINK"}
var flinkRoleTypes = []string{"JOBMANAGER", "TASKMANAGER", "FLINK_HISTORY_SERVER", "GATEWAY"}

var flinkSpec = serviceScraperSpec{
	Name:         FLINK_SCRAPER_NAME,
	Help:         "Flink service metrics",
	ServiceTypes: flinkServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Flink role JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "jvm_total_threads",
			Help:  "Flink role JVM thread count.",
			Query: roleMetricQuery("jvm_total_threads", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "jvm_gc_rate",
			Help:  "Flink role JVM garbage collections per second.",
			Query: roleMetricQuery("jvm_gc_rate", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "fd_open",
			Help:  "Flink role open file descriptors.",
			Query: roleMetricQuery("fd_open", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "fd_max",
			Help:  "Flink role maximum file descriptors.",
			Query: roleMetricQuery("fd_max", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "uptime",
			Help:  "Flink role uptime in seconds.",
			Query: roleMetricQuery("uptime", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "unexpected_exits_rate",
			Help:  "Flink role unexpected exits per second.",
			Query: roleMetricQuery("unexpected_exits_rate", flinkServiceTypes, flinkRoleTypes...),
		},
		{
			Name:  "oom_exits_rate",
			Help:  "Flink role out-of-memory exits per second.",
			Query: roleMetricQuery("oom_exits_rate", flinkServiceTypes, flinkRoleTypes...),
		},
	}, flinkServiceTypes, flinkRoleTypes...),
}

type ScrapeFlink struct{}

func (ScrapeFlink) Name() string {
	return FLINK_SCRAPER_NAME
}

func (ScrapeFlink) Help() string {
	return flinkSpec.Help
}

func (ScrapeFlink) Version() float64 {
	return 1.0
}

func (ScrapeFlink) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, flinkSpec, ch)
}

var _ Scraper = ScrapeFlink{}
