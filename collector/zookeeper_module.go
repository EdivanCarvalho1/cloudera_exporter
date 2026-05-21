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
			Name:  "data_nodes",
			Help:  "ZooKeeper znode count.",
			Query: roleMetricQuery("data_nodes", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "data_watches",
			Help:  "ZooKeeper watch count.",
			Query: roleMetricQuery("data_watches", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "open_connections",
			Help:  "ZooKeeper alive client connections.",
			Query: roleMetricQuery("open_connections", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "packets_receive_rate",
			Help:  "ZooKeeper packets received per second.",
			Query: roleMetricQuery("packets_receive_rate", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "packets_transmit_rate",
			Help:  "ZooKeeper packets sent per second.",
			Query: roleMetricQuery("packets_transmit_rate", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "min_request_latency",
			Help:  "ZooKeeper minimum request latency.",
			Query: roleMetricQuery("min_request_latency", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "average_request_latency",
			Help:  "ZooKeeper average request latency.",
			Query: roleMetricQuery("average_request_latency", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "max_request_latency",
			Help:  "ZooKeeper maximum request latency.",
			Query: roleMetricQuery("max_request_latency", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "data_size",
			Help:  "ZooKeeper data tree size.",
			Query: roleMetricQuery("data_size", zookeeperServiceTypes, zookeeperRoleTypes...),
		},
		{
			Name:  "zk_server_mode",
			Help:  "ZooKeeper server mode.",
			Query: roleMetricQuery("zk_server_mode", zookeeperServiceTypes, zookeeperRoleTypes...),
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
