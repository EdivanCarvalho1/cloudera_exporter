package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const KAFKA_SCRAPER_NAME = "kafka"

var kafkaServiceTypes = []string{"KAFKA"}
var kafkaRoleTypes = []string{"KAFKA_BROKER"}

var kafkaSpec = serviceScraperSpec{
	Name:         KAFKA_SCRAPER_NAME,
	Help:         "Kafka service metrics",
	ServiceTypes: kafkaServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "bytes_in_rate",
			Help:  "Kafka broker bytes in per second.",
			Query: rateRoleMetricQuery("bytes_in_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "bytes_out_rate",
			Help:  "Kafka broker bytes out per second.",
			Query: rateRoleMetricQuery("bytes_out_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "messages_in_rate",
			Help:  "Kafka broker messages in per second.",
			Query: rateRoleMetricQuery("messages_in_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "under_replicated_partitions",
			Help:  "Kafka under replicated partitions.",
			Query: roleMetricQuery("under_replicated_partitions", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "offline_partitions",
			Help:  "Kafka offline partitions.",
			Query: roleMetricQuery("offline_partitions", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "isr_shrinks_rate",
			Help:  "Kafka ISR shrinks per second.",
			Query: rateRoleMetricQuery("isr_shrinks_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "isr_expands_rate",
			Help:  "Kafka ISR expands per second.",
			Query: rateRoleMetricQuery("isr_expands_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "leader_election_rate",
			Help:  "Kafka leader elections per second.",
			Query: rateRoleMetricQuery("leader_election_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "request_handler_idle",
			Help:  "Kafka request handler idle percentage.",
			Query: roleMetricQuery("request_handler_idle_percent", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "network_processor_idle",
			Help:  "Kafka network processor idle percentage.",
			Query: roleMetricQuery("network_processor_idle_percent", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Kafka broker JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "jvm_threads",
			Help:  "Kafka broker JVM thread count.",
			Query: roleMetricQuery("jvm_threads", kafkaServiceTypes, kafkaRoleTypes...),
		},
	}, kafkaServiceTypes, kafkaRoleTypes...),
}

type ScrapeKafka struct{}

func (ScrapeKafka) Name() string {
	return KAFKA_SCRAPER_NAME
}

func (ScrapeKafka) Help() string {
	return kafkaSpec.Help
}

func (ScrapeKafka) Version() float64 {
	return 1.0
}

func (ScrapeKafka) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, kafkaSpec, ch)
}

var _ Scraper = ScrapeKafka{}
