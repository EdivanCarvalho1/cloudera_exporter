package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

const ZOOKEEPER_SCRAPER_NAME = "zookeeper"

var zookeeperServiceTypes = []string{"ZOOKEEPER"}
var zookeeperRoleTypes = []string{"SERVER"}

var zookeeperSpec = serviceScraperSpec{
	Name:         ZOOKEEPER_SCRAPER_NAME,
	Help:         "ZooKeeper service metrics",
	ServiceTypes: zookeeperServiceTypes,
	Metrics: appendCommonRoleMetrics([]serviceTimeseriesMetric{
		{
			Name:  "outstanding_requests",
			Help:  "ZooKeeper outstanding requests.",
			Query: roleMetricQuery("outstanding_requests", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "znode_count",
			Help:  "ZooKeeper znode count.",
			Query: roleMetricQuery("znode_count", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "watch_count",
			Help:  "ZooKeeper watch count.",
			Query: roleMetricQuery("watch_count", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "num_alive_connections",
			Help:  "ZooKeeper alive client connections.",
			Query: roleMetricQuery("num_alive_connections", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "packets_received_rate",
			Help:  "ZooKeeper packets received per second.",
			Query: rateRoleMetricQuery("packets_received_rate", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "packets_sent_rate",
			Help:  "ZooKeeper packets sent per second.",
			Query: rateRoleMetricQuery("packets_sent_rate", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "latency_min",
			Help:  "ZooKeeper minimum request latency.",
			Query: roleMetricQuery("latency_min", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "latency_avg",
			Help:  "ZooKeeper average request latency.",
			Query: roleMetricQuery("latency_avg", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "latency_max",
			Help:  "ZooKeeper maximum request latency.",
			Query: roleMetricQuery("latency_max", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "jvm_heap_used_mb",
			Help:  "ZooKeeper server JVM heap used in MB.",
			Query: roleMetricQuery("jvm_heap_used_mb", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
	}, zookeeperServiceTypes, zookeeperRoleTypes...),
}

type ScrapeZookeeper struct{}

func (ScrapeZookeeper) Name() string {
	return ZOOKEEPER_SCRAPER_NAME
}

func (ScrapeZookeeper) Help() string {
	return zookeeperSpec.Help
}

func (ScrapeZookeeper) Version() float64 {
	return 1.0
}

func (ScrapeZookeeper) Scrape(ctx context.Context, config *Collector_connection_data, ch chan<- prometheus.Metric) error {
	return scrapeServiceModule(ctx, *config, zookeeperSpec, ch)
}

var _ Scraper = ScrapeZookeeper{}
