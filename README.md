# Keedio Cloudera Exporter

![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg?style=plastic) ![GO](https://img.shields.io/badge/go_version-1.12-blue.svg?style=plastic) ![Docker](https://img.shields.io/badge/docker_container-ready-brightgreen.svg?style=plastic)




## What is it?
This project builds a Prometheus Exporter to collect Cloudera status and usage metrics.




## Modules
This exporter scrape the metrics by independent modules (Scrapers). This modules are:
* **Status:**  Scrapes the metrics about the current status of the Clusters, services, roles and hosts
* **Hosts:**  Scrapes the metrics about the Hosts: CPU usage, RAM, SWAP, Agent stats and more useful metrics
* **HDFS:**  Scrapes the metrics about HDFS: Capacity, blocks stats, file stats, Namenode properties and Snapshots.
* **Impala:**  Scrapes the metrics about Impala: Catalog, usage stats, queries stats, state-store info …




## Supported Service Modules
The exporter keeps one Scraper per module. Modules can be enabled or disabled in the `[modules]` section of `config.ini`.

* Status
* Hosts
* HDFS
* HBase
* Hive
* Impala
* Kafka
* Kudu
* Spark
* Yarn
* Zookeeper
* NiFi
* Flink

The modules use the Cloudera Manager API. Status and inventory data are read from the cluster, service and role endpoints, while detailed service metrics are read from `/timeseries` using tsquery.

Some Cloudera Manager metric names vary by Cloudera Manager version, CDH/CDP release and installed parcels or CSDs. If a metric does not exist in a given installation, that metric is skipped and the remaining modules continue scraping. NiFi and Flink are treated as optional CSD/custom services; if they are absent, their modules do not fail the exporter.

Example module flags:

```ini
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
```

## Building and Running
To launch the exporter we recommend use a Docker container.  Whether it is inside a container or in the local system, some Golang packages are needed as dependencies of the code to be able to implement some Prometheus functions.




### Docker Deploy
#### Configuration
Edit the *config.ini* file and change the host ip (cloudera_manager) in the *target* section by your Cloudera Manager IP addr.  If you want to use "cloudera_manager" as a name-domain, edit the *Makefile.common* file and change "--add-host" property in the *Docker_RUN_FLAGS* variable.
```sh
# Compile on docker and create container
make docker_build

# Launch on docker
make docker_run

# Compile and launch
make docker
```




### Local Deploy
#### Configuration
Edit the *config.ini* file and change the host ip (cloudera_manager) in the *target* section by your Cloudera Manager IP addr.  If you want to use "cloudera_manager" as a name-domain, add the entry to your */etc/hosts* file.
```sh
# Compile on local
make all

# Launch on local on shell
./cloudera_exporter --config-file config.ini
```

Cloudera Exporter args:
```sh
  usage: cloudera_exporter [<flags>]

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --config-file="config.ini" Path to ini file.
      --web.listen-address=""    Listent Address.
      --num-procs=0              Number Processes for parallel execution
      --log-level=0              Debug Log Mode
      --timeout-offset=0.25      Time to subtract from timeout in seconds.
      --version                  Show application version.
```


### Docker Deploy
#### Build Docker Image
```sh
# First of all, edit the config.ini file with the parameters for your environment

# Build de Docker Image with the *Dockerfile* in this repository:
docker build --build-arg VERSION="$(cat VERSION)" -t "keedio/kce:$(cat VERSION)" .

# Run the container. (Maybe you want to open the exporter port)
docker run keedio/kce:$(cat VERSION) [-p 9200:9200] --name kce
```





### Clean Environment
To clean the local binary and the cloudera_exporter container
```sh
make clean
```




### Test if is running
Test if cloudera_exporter is running
```sh
make test
```




## Metrics Catalog
To get a list of all the metrics of each Scraper, read the  Scraper, read the [METRICS_CATALOG.md](METRICS_CATALOG.md) file.




### Contacts and Ownership
To contact with developers for any question or proposal read the [MAINTAINERS.md](MAINTAINERS.md) file.
