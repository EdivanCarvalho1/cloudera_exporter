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
			Name:  "kafka_bytes_received_rate",
			Help:  "Kafka broker bytes in per second.",
			Query: roleMetricQuery("kafka_bytes_received_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_bytes_fetched_rate",
			Help:  "Kafka broker bytes out per second.",
			Query: roleMetricQuery("kafka_bytes_fetched_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_messages_received_rate",
			Help:  "Kafka broker messages in per second.",
			Query: roleMetricQuery("kafka_messages_received_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_total_produce_requests_per_sec_rate",
			Help:  "Kafka produce requests per second.",
			Query: roleMetricQuery("kafka_total_produce_requests_per_sec_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_total_fetch_requests_per_sec_rate",
			Help:  "Kafka fetch requests per second.",
			Query: roleMetricQuery("kafka_total_fetch_requests_per_sec_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_bytes_rejected_rate",
			Help:  "Kafka rejected bytes per second.",
			Query: roleMetricQuery("kafka_bytes_rejected_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_under_replicated_partitions",
			Help:  "Kafka under replicated partitions.",
			Query: roleMetricQuery("kafka_under_replicated_partitions", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_offline_partitions",
			Help:  "Kafka offline partitions.",
			Query: roleMetricQuery("kafka_offline_partitions", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_isr_shrinks_rate",
			Help:  "Kafka ISR shrinks per second.",
			Query: roleMetricQuery("kafka_isr_shrinks_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_isr_expands_rate",
			Help:  "Kafka ISR expands per second.",
			Query: roleMetricQuery("kafka_isr_expands_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_leader_election_rate",
			Help:  "Kafka leader elections per second.",
			Query: roleMetricQuery("kafka_leader_election_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_unclean_leader_elections_rate",
			Help:  "Kafka unclean leader elections per second.",
			Query: roleMetricQuery("kafka_unclean_leader_elections_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_request_handler_avg_idle_rate",
			Help:  "Kafka request handler idle percentage.",
			Query: roleMetricQuery("kafka_request_handler_avg_idle_rate", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_network_processor_avg_idle",
			Help:  "Kafka network processor idle percentage.",
			Query: roleMetricQuery("kafka_network_processor_avg_idle", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_active_controller",
			Help:  "Kafka active controller state.",
			Query: roleMetricQuery("kafka_active_controller", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_broker_state",
			Help:  "Kafka broker state.",
			Query: roleMetricQuery("kafka_broker_state", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_active_broker_count",
			Help:  "Kafka active broker count.",
			Query: serviceMetricQuery("kafka_active_broker_count", kafkaServiceTypes),
		},
		{
			Name:  "kafka_response_queue_size",
			Help:  "Kafka response queue size.",
			Query: roleMetricQuery("kafka_response_queue_size", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "Kafka broker JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", kafkaServiceTypes, kafkaRoleTypes...),
		},
		{
			Name:  "kafka_thread_count",
			Help:  "Kafka broker JVM thread count.",
			Query: roleMetricQuery("kafka_thread_count", kafkaServiceTypes, kafkaRoleTypes...),
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
