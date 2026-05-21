package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const OOZIE_SCRAPER_NAME = "oozie"

var oozieServiceTypes = []string{"OOZIE"}

var oozieSpec = serviceScraperSpec{
	Name:         OOZIE_SCRAPER_NAME,
	Help:         "Oozie service metrics",
	ServiceTypes: oozieServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "oozie_memory_total_used",
			Help:  "Oozie total memory used.",
			Query: roleMetricQuery("oozie_memory_total_used", oozieServiceTypes),
		},
		{
			Name:  "oozie_memory_heap_used",
			Help:  "Oozie heap memory used.",
			Query: roleMetricQuery("oozie_memory_heap_used", oozieServiceTypes),
		},
		{
			Name:  "oozie_memory_non_heap_used",
			Help:  "Oozie non-heap memory used.",
			Query: roleMetricQuery("oozie_memory_non_heap_used", oozieServiceTypes),
		},
		{
			Name:  "oozie_jdbc_active_connections_histogram_avg",
			Help:  "Oozie average active JDBC connections.",
			Query: roleMetricQuery("oozie_jdbc_active_connections_histogram_avg", oozieServiceTypes),
		},
		{
			Name:  "oozie_jdbc_active_connections_histogram_max",
			Help:  "Oozie maximum active JDBC connections.",
			Query: roleMetricQuery("oozie_jdbc_active_connections_histogram_max", oozieServiceTypes),
		},
		{
			Name:  "oozie_jdbc_idle_connections_histogram_avg",
			Help:  "Oozie average idle JDBC connections.",
			Query: roleMetricQuery("oozie_jdbc_idle_connections_histogram_avg", oozieServiceTypes),
		},
		{
			Name:  "oozie_jdbc_idle_connections_histogram_max",
			Help:  "Oozie maximum idle JDBC connections.",
			Query: roleMetricQuery("oozie_jdbc_idle_connections_histogram_max", oozieServiceTypes),
		},
		{
			Name:  "oozie_workflow_action_query_executor_get_action_duration_timer_avg",
			Help:  "Oozie average workflow action query duration.",
			Query: roleMetricQuery("oozie_workflow_action_query_executor_get_action_duration_timer_avg", oozieServiceTypes),
		},
		{
			Name:  "oozie_workflow_action_query_executor_get_action_completed_duration_timer_avg",
			Help:  "Oozie average completed workflow action query duration.",
			Query: roleMetricQuery("oozie_workflow_action_query_executor_get_action_completed_duration_timer_avg", oozieServiceTypes),
		},
		{
			Name:  "oozie_jvm_pause_time_rate",
			Help:  "Oozie JVM pause time rate.",
			Query: roleMetricQuery("oozie_jvm_pause_time_rate", oozieServiceTypes),
		},
		{
			Name:  "oozie_jvm_pauses_warn_threshold_rate",
			Help:  "Oozie JVM pauses above warning threshold per second.",
			Query: roleMetricQuery("oozie_jvm_pauses_warn_threshold_rate", oozieServiceTypes),
		},
	}, oozieServiceTypes),
}

type ScrapeOozie struct{}

func (ScrapeOozie) Name() string {
	return OOZIE_SCRAPER_NAME
}

func (ScrapeOozie) Help() string {
	return oozieSpec.Help
}

func (ScrapeOozie) Version() float64 {
	return 1.0
}

func (ScrapeOozie) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, oozieSpec, ch)
}

var _ Scraper = ScrapeOozie{}
