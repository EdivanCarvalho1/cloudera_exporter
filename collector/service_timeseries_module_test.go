package collector

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	log "keedio/cloudera_exporter/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

func initTestLogger() {
	log.Init(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, 0)
}

func TestSanitizeMetricName(t *testing.T) {
	tests := map[string]string{
		"HBase Region Count":       "hbase_region_count",
		"9bad.metric/name":         "_9bad_metric_name",
		"kbdi__hdfs---capacity":    "kbdi_hdfs_capacity",
		"   ":                      "unknown",
		"queries_successful_rate":  "queries_successful_rate",
		"flink.task-slots/total":   "flink_task_slots_total",
		"kafka#bytes#in#rate####":  "kafka_bytes_in_rate",
		"zookeeper latency avg ms": "zookeeper_latency_avg_ms",
	}

	for input, expected := range tests {
		if actual := sanitizeMetricName(input); actual != expected {
			t.Fatalf("sanitizeMetricName(%q) = %q, want %q", input, actual, expected)
		}
	}
}

func TestGetTimeseriesSeriesFromFixture(t *testing.T) {
	content, err := ioutil.ReadFile("../testdata/timeseries_hdfs.json")
	if err != nil {
		t.Fatal(err)
	}

	series := getTimeseriesSeries(gjson.ParseBytes(content))
	if len(series) != 1 {
		t.Fatalf("len(series) = %d, want 1", len(series))
	}

	labels := timeseriesLabels(series[0])
	if labels["cluster"] != "ClusterOne" || labels["serviceType"] != "HDFS" || labels["roleType"] != "NAMENODE" {
		t.Fatalf("unexpected labels: %#v", labels)
	}
}

func TestLatestTimeseriesValueUsesLastValidPoint(t *testing.T) {
	serie := gjson.Parse(`{
		"data": [
			{"timestamp": "2026-05-18T10:00:00Z", "value": 1},
			{"timestamp": "2026-05-18T10:01:00Z", "value": null},
			{"timestamp": "2026-05-18T10:02:00Z", "value": "2.5"}
		]
	}`)

	value, ok := latestTimeseriesValue(serie)
	if !ok {
		t.Fatal("expected a value")
	}
	if value != 2.5 {
		t.Fatalf("value = %f, want 2.5", value)
	}
}

func TestLatestTimeseriesValueSkipsEmptyAndNullSeries(t *testing.T) {
	tests := []string{
		`{"data":[]}`,
		`{"data":[{"value":null}]}`,
		`{"data":[{}]}`,
	}

	for _, test := range tests {
		if value, ok := latestTimeseriesValue(gjson.Parse(test)); ok {
			t.Fatalf("latestTimeseriesValue(%s) = %f, true; want false", test, value)
		}
	}
}

func TestScrapeServiceModuleDoesNotFailOnMissingServiceOrQueryError(t *testing.T) {
	initTestLogger()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v54/clusters":
			w.Write([]byte(`{"items":[{"name":"ClusterOne","displayName":"ClusterOne"}]}`))
		case "/api/v54/clusters/ClusterOne/services":
			w.Write([]byte(`{"items":[]}`))
		case "/api/v54/timeseries":
			http.Error(w, "not found", http.StatusNotFound)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	host, port, err := net.SplitHostPort(server.Listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	config := Collector_connection_data{
		Host:        host,
		Port:        port,
		Api_version: "v54",
		User:        "user",
		Passwd:      "password",
	}
	spec := serviceScraperSpec{
		Name:         "missing",
		Help:         "missing service",
		ServiceTypes: []string{"MISSING"},
		Metrics: []serviceTimeseriesMetric{
			{Name: "bad_metric", Help: "bad metric", Query: "SELECT LAST(bad_metric) WHERE serviceType = \"MISSING\""},
		},
	}
	ch := make(chan prometheus.Metric, 1)
	err = scrapeServiceModule(context.Background(), config, spec, ch)
	if err != nil {
		t.Fatalf("scrapeServiceModule returned error: %s", err)
	}
}
