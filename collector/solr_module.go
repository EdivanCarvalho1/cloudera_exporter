package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const SOLR_SCRAPER_NAME = "solr"

var solrServiceTypes = []string{"SOLR"}

var solrSpec = serviceScraperSpec{
	Name:         SOLR_SCRAPER_NAME,
	Help:         "Solr service metrics",
	ServiceTypes: solrServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "solr_active_cores",
			Help:  "Solr active cores.",
			Query: roleMetricQuery("solr_active_cores", solrServiceTypes),
		},
		{
			Name:  "solr_all_cores",
			Help:  "Solr all cores.",
			Query: roleMetricQuery("solr_all_cores", solrServiceTypes),
		},
		{
			Name:  "solr_recovering_cores",
			Help:  "Solr recovering cores.",
			Query: roleMetricQuery("solr_recovering_cores", solrServiceTypes),
		},
		{
			Name:  "solr_info_requests_rate",
			Help:  "Solr request rate.",
			Query: roleMetricQuery("solr_info_requests_rate", solrServiceTypes),
		},
		{
			Name:  "solr_info_errors_rate",
			Help:  "Solr error rate.",
			Query: roleMetricQuery("solr_info_errors_rate", solrServiceTypes),
		},
		{
			Name:  "solr_info_timeouts_rate",
			Help:  "Solr timeout rate.",
			Query: roleMetricQuery("solr_info_timeouts_rate", solrServiceTypes),
		},
		{
			Name:  "solr_info_avg_time_per_request",
			Help:  "Solr average request time.",
			Query: roleMetricQuery("solr_info_avg_time_per_request", solrServiceTypes),
		},
		{
			Name:  "solr_info_99th_pc_request_time",
			Help:  "Solr 99th percentile request time.",
			Query: roleMetricQuery("solr_info_99th_pc_request_time", solrServiceTypes),
		},
		{
			Name:  "solr_info_avg_requests_per_second",
			Help:  "Solr average requests per second.",
			Query: roleMetricQuery("solr_info_avg_requests_per_second", solrServiceTypes),
		},
		{
			Name:  "solr_info_time_rate",
			Help:  "Solr request time rate.",
			Query: roleMetricQuery("solr_info_time_rate", solrServiceTypes),
		},
		{
			Name:  "solr_info_median_request_time",
			Help:  "Solr median request time.",
			Query: roleMetricQuery("solr_info_median_request_time", solrServiceTypes),
		},
		{
			Name:  "solr_core_status_collection_status",
			Help:  "Solr core status collection status.",
			Query: roleMetricQuery("solr_core_status_collection_status", solrServiceTypes),
		},
		{
			Name:  "solr_core_status_collection_duration",
			Help:  "Solr core status collection duration.",
			Query: roleMetricQuery("solr_core_status_collection_duration", solrServiceTypes),
		},
		{
			Name:  "solr_counters_elements",
			Help:  "Solr counter elements.",
			Query: roleMetricQuery("solr_counters_elements", solrServiceTypes),
		},
		{
			Name:  "solr_counters_relations",
			Help:  "Solr counter relations.",
			Query: roleMetricQuery("solr_counters_relations", solrServiceTypes),
		},
	}, solrServiceTypes),
}

type ScrapeSolr struct{}

func (ScrapeSolr) Name() string {
	return SOLR_SCRAPER_NAME
}

func (ScrapeSolr) Help() string {
	return solrSpec.Help
}

func (ScrapeSolr) Version() float64 {
	return 1.0
}

func (ScrapeSolr) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, solrSpec, ch)
}

var _ Scraper = ScrapeSolr{}
