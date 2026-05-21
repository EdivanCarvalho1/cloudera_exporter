package config_parser

import (
	"io/ioutil"
	"os"
	"testing"

	log "keedio/cloudera_exporter/logger"
)

func TestParseConfigRecognizesServiceModules(t *testing.T) {
	log.Init(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, 0)

	configFile, err := ioutil.TempFile("", "cloudera-exporter-*.ini")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configFile.Name())

	_, err = configFile.WriteString(`[target]
host = cm.example.local
port = 7180
version = v54

[user]
username = exporter
password = secret

[modules]
global_status_module = true
host_module = true
hdfs_module = true
impala_module = true
hbase_module = true
hive_module = true
kafka_module = true
kudu_module = true
spark_module = true
yarn_module = true
zookeeper_module = true
nifi_module = true
flink_module = true
oozie_module = true
ozone_module = true
solr_module = true
ranger_module = true
atlas_module = true
hue_module = true
knox_module = true
flume_module = true
observability_module = true

[system]
num_procs = 2
deploy_ip =
deploy_port = 9200
log_level = 0
`)
	if err != nil {
		t.Fatal(err)
	}
	if err := configFile.Close(); err != nil {
		t.Fatal(err)
	}

	config, err := Parse_config(configFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	enabled := map[string]bool{}
	for scraper, isEnabled := range config.Scrapers.Scrapers {
		if isEnabled {
			enabled[scraper.Name()] = true
		}
	}

	expected := []string{
		"status_collector",
		"host",
		"hdfs",
		"impala",
		"hbase",
		"hive",
		"kafka",
		"kudu",
		"spark",
		"yarn",
		"zookeeper",
		"nifi",
		"flink",
		"oozie",
		"ozone",
		"solr",
		"ranger",
		"atlas",
		"hue",
		"knox",
		"flume",
		"observability",
	}
	for _, name := range expected {
		if !enabled[name] {
			t.Fatalf("expected scraper %s to be enabled; enabled=%#v", name, enabled)
		}
	}
}
