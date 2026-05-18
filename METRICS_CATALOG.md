# Keedio Cloudera Exporter metrics catalog #

## Metrics
All the available metrics

### Status Values

| Value | Mean       |
|:-----:|:----------:|
|   0   | NO DATA    |
|   1   | OK         |
|   2   | UNKNOWN    |
|   3   | DISABLED   |
|   4   | CONCERNING |
|   5   | BAD        |


### Status Module Metrics
| Metric Name                    | Unit             | C.M. Version  | Description       | Metadata                                                                                      |
|--------------------------------|:----------------:|:-------------:|-------------------|-----------------------------------------------------------------------------------------------|
| kbdi_global_status_cluster_up  |  [1-0] (OK\|KO)  |  > 5.8         |  Cluster Status  |  cluster_name, cluster_state, full_version, maintenance_mode                                  |
| kbdi_global_status_host_up     |  [1-0] (OK\|KO)  |  > 5.8         |  Host Status     |  commission_state, health_summary, host_id, hostname, ip, maintenance_mode                    |
| kbdi_global_status_service_up  |  [1-0] (OK\|KO)  |  > 5.8         |  Service Status  |  health_summary, service_name, service_state, service_type                                    |
| kbdi_global_status_role_up     |  [1-0] (OK\|KO)  |  > 5.8         |  Role Status     |  health_summary, role_name, role_state, role_type, hostname, host_id, service                 |


### Host Module Metrics
| Metric Name                         | Unit              | C.M. Version   | Description                                             | Metadata                                                                   |
|-------------------------------------|:-----------------:|:--------------:|---------------------------------------------------------|----------------------------------------------------------------------------|
| kbdi_host_agent_cpu_system_percent  |  %                |  > 5.8         |  CPU % usage in Cloudera agent system operations        |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_agent_cpu_user_percent    |  %                |  > 5.8         |  CPU % usage in Cloudera agent user operations          |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_agent_phys_mem_use        |  bytes            |  > 5.8         |  Physical Memory usage in Cloudera agent                |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_agent_virt_mem_use        |  bytes            |  > 5.8         |  Virtual Memory usage in Cloudera agent                 |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_alerts                    |  alerts           |  > 5.8         |  Num of alerts for each host                            |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_clock_offset              |  ms               |  > 5.8         |  Milliseconds of clock offset for each host             |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_cores                 |  cores            |  > 5.8         |  Num of Cores for each host                             |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_iddle_percent         |  %                |  > 5.8         |  % of time for CPU Iddle                                |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_iowait_percent        |  %                |  > 5.8         |  % of time for CPU IOWait instructions                  |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_percent_by_host       |  %                |  > 5.8         |  % of time for CPU Usage                                |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_system_percent        |  %                |  > 5.8         |  % of time for CPU System instructions                  |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_cpu_user_percent          |  %                |  > 5.8         |  % of time for CPU User instructions                    |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_dns_resolution_time       |  ms               |  > 5.8         |  DNS query time resolution                              |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_load_1_by_host            |  Usage By Thread  |  > 5.8         |  CPU usage in last 1 minutes (Linux CPU usage format)   |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_load_5_by_host            |  Usage By Thread  |  > 5.8         |  CPU usage in last 5 minutes (Linux CPU usage format)   |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_load_15_by_host           |  Usage By Thread  |  > 5.8         |  CPU usage in last 15 minutes (Linux CPU usage format)  |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_mem_free_by_host          |  bytes            |  > 5.8         |  Free RAM memory for each host                          |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_mem_total_by_host         |  bytes            |  > 5.8         |  Total RAM memory for each host                         |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_mem_used_by_host          |  bytes            |  > 5.8         |  Used RAM memory for each host                          |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_mem_writeback_by_host     |  bytes            |  > 5.8         |  WriteBack RAM memory for each host                     |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_swap_free_by_host         |  bytes            |  > 5.8         |  Free SWAP memory for each host                         |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_swap_out_by_host          |  pages            |  > 5.8         |  Out SWAP memory for each host                          |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_swap_total_by_host        |  bytes            |  > 5.8         |  Total SWAP memory for each host                        |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_swap_used_by_host         |  bytes            |  > 5.8         |  Used SWAP memory for each host                         |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |
| kbdi_host_uptime                    |  seconds          |  > 5.8         |  Host Uptime                                            |  cluster, hostid, hostname, is_border_node, is_master_node, is_worker_node |


### HDFS Module Metrics
| Metric Name                            | Unit              | C.M. Version   | Description                                                     |  Metadata |
|----------------------------------------|:-----------------:|:--------------:|-----------------------------------------------------------------|-----------|
| kbdi_hdfs_dfs_capacity                 |  %                |  > 5.8         |  Distributed File System Capacity                               |  cluster  |
| kbdi_hdfs_dfs_capacity_used            |  %                |  > 5.8         |  Distributed File System Capacity Used                          |  cluster  |
| kbdi_hdfs_dfs_capacity_used_percent    |  bytes            |  > 5.8         |  Distributed File System Capacity Used in X Percent             |  cluster  |
| kbdi_hdfs_dfs_capacity_non_hdfs_used   |  bytes            |  > 5.8         |  Distributed File System Capacity Used by Non HDFS File System  |  cluster  |
| kbdi_hdfs_block_capacity               |  alerts           |  > 5.8         |  Distributed File System Num Blocks Capacity                    |  cluster  |
| kbdi_hdfs_block_total                  |  ms               |  > 5.8         |  Distributed File System Num Blocks Total                       |  cluster  |
| kbdi_hdfs_block_corrupt_replicas       |  cores            |  > 5.8         |  Distributed File System Num Block with corrupted replicas      |  cluster  |
| kbdi_hdfs_block_excess                 |  %                |  > 5.8         |  Distributed File System Num Excess blocks                      |  cluster  |
| kbdi_hdfs_block_missing                |  %                |  > 5.8         |  Distributed File System Num Missing blocks                     |  cluster  |
| kbdi_hdfs_block_under_replicated       |  %                |  > 5.8         |  Distributed File System Num Under-Replicated blocks            |  cluster  | 
| kbdi_hdfs_block_write                  |  %                |  > 5.8         |  Distributed File System Rate Writed blocks                     |  cluster  |
| kbdi_hdfs_block_read                   |  %                |  > 5.8         |  Distributed File System Rate Readed blocks                     |  cluster  |
| kbdi_hdfs_files_total                  |  ms               |  > 5.8         |  Distributed File System Num Total Files In HDFS                |  cluster  |
| kbdi_hdfs_files_size_avg               |  Usage By Thread  |  > 5.8         |  Distributed File System Avg Size of Files In HDFS              |  cluster  |
| kbdi_hdfs_heartbeats_expired           |  Usage By Thread  |  > 5.8         |  Distributed File System Num Total Heartbeats Expired           |  cluster  |
| kbdi_hdfs_namenode_fd_max_descriptors  |  Usage By Thread  |  > 5.8         |  Distributed File System Namenode Max File Descriptors          |  cluster  |
| kbdi_hdfs_snapshot_num                 |  bytes            |  > 5.8         |  Distributed File System Num Total Snapshots                    |  cluster  |
| kbdi_hdfs_snapshot_dirs                |  bytes            |  > 5.8         |  Distributed File System Num Total Snapshottable Dirs           |  cluster  |

### Impala Module Metrics

| Metric Name                                                               | Unit             | C.M. Version   | Description                                                                                                                                                                                                                                                                                                                                                                                            | Metadata             |
|---------------------------------------------------------------------------|:-----------------|:--------------:|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| kbdi_impala_cgroup_cpu_system_rate                                        |  s/s             |  > 5.8         |  CPU usage of the role's cgroup	                                                                                                                                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_cgroup_cpu_user_rate                                          |  s/s             |  > 5.8         |  The ratio of cpu space usage of the system for cgroup                                                                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_cgroup_mem_page_cache                                         |  bytes           |  > 5.8         |  User Space CPU usage of the role's cgroup	                                                                                                                                                                                                                                                                                                                                                           |  cluster, entityName | 
| kbdi_impala_cgroup_mem_rss                                                |  bytes           |  > 5.8         |  Page cache usage of the role's cgroup	                                                                                                                                                                                                                                                                                                                                                               |  cluster, entityName | 
| kbdi_impala_cgroup_mem_swap                                               |  bytes           |  > 5.8         |  Resident memory of the role's cgroup	                                                                                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_cgroup_read_bytes_rate                                        |  bytes/s         |  > 5.8         |  Swap usage of the role's cgroup	                                                                                                                                                                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_cgroup_read_ios_rate                                          |  IO opts/s       |  > 5.8         |  Bytes read from all disks by the role's cgroup	                                                                                                                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_cgroup_write_bytes_rate                                       |  bytes/s         |  > 5.8         |  Number of read I/O operations from all disks by the role's cgroup	                                                                                                                                                                                                                                                                                                                                   |  cluster, entityName | 
| kbdi_impala_cgroup_write_ios_rate                                         |  IO optsalida/s  |  > 5.8         |  Bytes written to all disks by the role's cgroup	                                                                                                                                                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_impala_catalogserver_jvm_heap_committed_usage_bytes           |  bytes           |  > 5.8         |  Number of write I/O operations to all disks by the role's cgroup	                                                                                                                                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_impala_catalogserver_jvm_heap_current_usage_bytes             |  bytes           |  > 5.8         |  Current byte usage by Jvm heap                                                                                                                                                                                                                                                                                                                                                                        |  cluster, entityName | 
| kbdi_impala_impala_catalogserver_jvm_heap_init_usage_bytes                |  bytes           |  > 5.8         |  Use of bytes in the initialization of Jvm heap                                                                                                                                                                                                                                                                                                                                                        |  cluster, entityName | 
| kbdi_impala_impala_catalogserver_jvm_heap_max_usage_bytes                 |  bytes           |  > 5.8         |  Maximum use of bytes by Jvm heap                                                                                                                                                                                                                                                                                                                                                                      |  cluster, entityName | 
| kbdi_impala_impala_query_admission_wait_rate                              |  sin unidad      |  5.16.1        |  The waiting time ratio for each query admitted                                                                                                                                                                                                                                                                                                                                                        |  cluster, entityName | 
| kbdi_impala_impala_query_bytes_streamed_rate                              |  bytes/s         |  5.16.1        |  The ratio of bytes sent between Impala Daemons when processing a query                                                                                                                                                                                                                                                                                                                                |  cluster, entityName | 
| kbdi_impala_impala_query_cm_cpu_milliseconds_rate                         |  sin nidad       |  5.16.1        |  Metric that measures the ratio                                                                                                                                                                                                                                                                                                                                                                        |  cluster, entityName | 
| kbdi_impala_impala_query_hdfs_bytes_read_rate                             |  bytes/s         |  5.16.1        |  Total of bytes read from hdfs for that query.                                                                                                                                                                                                                                                                                                                                                         |  cluster, entityName | 
| kbdi_impala_impala_query_hdfs_bytes_written_rate                          |  bytes/s         |  5.16.1        |  Total of bytes written from hdfs for that query                                                                                                                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_impala_query_memory_accrual_rate                              |  bytes x s / s   |  5.16.1        |  Total of the accumulated memory per query. It is calculated: average use of aggregate memory of the query x its duration                                                                                                                                                                                                                                                                              |  cluster, entityName | 
| kbdi_impala_impala_query_memory_spilled_rate                              |  sin unidad      |  5.16.1        |  The amount of memory spilled on the disk                                                                                                                                                                                                                                                                                                                                                              |  cluster, entityName | 
| kbdi_impala_impala_query_query_duration_rate                              |  ms/s            |  5.16.1        |  The ratio of duration per query                                                                                                                                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_impala_query_thread_cpu_time_rate                             |  ms/s            |  5.16.1        |  The sum of the CPU time used by all the subprocesses of the query kbdi_impala_mem_rss                                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_mem_rss                                                       |  bytes           |  > 5.8         |  Memory rss (Resident set size) of cgroup                                                                                                                                                                                                                                                                                                                                                              |  cluster, entityName | 
| kbdi_impala_mem_swap                                                      |  bytes           |  > 5.8         |  Memory swap of cgroup                                                                                                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_mem_virtual                                                   |  bytes           |  > 5.8         |  Virtual memory used                                                                                                                                                                                                                                                                                                                                                                                   |  cluster, entityName | 
| kbdi_impala_oom_exits_rate                                                |  exits/s         |  > 5.8         |  The number of times the role's backing process was killed due to an OutOfMemory error. This counter is only incremented if the Cloudera Manager "Kill When Out of Memory" option is enabled.	                                                                                                                                                                                                         |  cluster, entityName | 
| kbdi_impala_queries_ingested_rate                                         |  queries/s       |  > 5.8         |  The ratio of queries ingested by the service monitor (Service Monitor)                                                                                                                                                                                                                                                                                                                                |  cluster, entityName | 
| kbdi_impala_queries_oom_rate                                              |  queries/s       |  > 5.8         |  Metric measuring impala queries whose memory consumption exceeds what is permitted                                                                                                                                                                                                                                                                                                                    |  cluster, entityName | 
| kbdi_impala_queries_rejected_rate                                         |  queries/s       |  > 5.8         |  The total number of queries made to Llama that have been rejected over the life of this Impala Daemon.	                                                                                                                                                                                                                                                                                               |  cluster, entityName | 
| kbdi_impala_queries_spilleed_rate                                         |  sin unidad      |  > 5.8         |  The ratio of spilled kerys                                                                                                                                                                                                                                                                                                                                                                            |  cluster, entityName | 
| kbdi_impala_queries_successful_rate                                       |  queries/s       |  > 5.8         |  Metric measuring impala consultations that were successfully executed                                                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_queries_time_out_rate                                         |  queries/s       |  > 5.8         |  The total number of requests made to Llama that have timed out over the life of this Impala Daemon.	                                                                                                                                                                                                                                                                                                 |  cluster, entityName | 
| kbdi_impala_read_bytes_rate                                               |  bytes           |  > 5.8         |  the ratio of bytes read by impala                                                                                                                                                                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_last            |  seconds         |  > 5.8         |  The most recent interval between heartbeats from this Impala Daemon to the StateStore.	                                                                                                                                                                                                                                                                                                               |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_max             |  seconds         |  > 5.8         |  The maximum interval between heartbeats from this Impala Daemon to the StateStore. This is calculated over the lifetime of the Impala Daemon.	                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_mean            |  seconds         |  > 5.8         |  The average interval between heartbeats from this Impala Daemon to the StateStore. This is calculated over the lifetime of the Impala Daemon.	                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_min             |  seconds         |  > 5.8         |  The minimum interval between heartbeats from this Impala Daemon to the StateStore. This is calculated over the lifetime of the Impala Daemon.	                                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_rate            |  muestras/s      |  > 5.8         |  The total number of samples taken of the Impala Daemon's StateStore heartbeat interval.	                                                                                                                                                                                                                                                                                                             |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_heartbeat_interval_time_stddev          |  seconds         |  > 5.8         |  The standard deviation in the interval between heartbeats from this Impala Daemon to the StateStore. This is calculated over the lifetime of the Impala Daemon.	                                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_last_recovery_duration                  |  seconds         |  > 5.8         |  The amount of time the StateStore subscriber took to recover the connection the last time it was lost.	                                                                                                                                                                                                                                                                                               |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_statestore_client_cache_clients_in_use  |  connections     |  > 5.8         |  The number of active StateStore subscriber clients in this Impala Daemon's client cache. These clients are for communication from this role to the StateStore.	                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_statestore_subscriber_statestore_client_cache_total_clients   |  connections     |  > 5.8         |  The total number of StateStore subscriber clients in this Impala Daemon's client cache. These clients are for communication from this role to the StateStore.	                                                                                                                                                                                                                                       |  cluster, entityName | 
| kbdi_impala_tcmalloc_bytes_in_use                                         |  bytes           |  > 5.8         |  Number of bytes used by the application. This will not typically match the memory use reported by the OS, because it does not include TCMalloc overhead or memory fragmentation.	                                                                                                                                                                                                                     |  cluster, entityName | 
| kbdi_impala_tcmalloc_pageheap_free_bytes                                  |  bytes           |  > 5.8         |  Number of bytes in free, mapped pages in page heap. These bytes can be used to fulfill allocation requests. They always count towards virtual memory usage, and unless the underlying memory is swapped out by the OS, they also count towards physical memory usage.	                                                                                                                               |  cluster, entityName | 
| kbdi_impala_tcmalloc_pageheap_unmapped_bytes                              |  bytes           |  > 5.8         |  Number of bytes in free, unmapped pages in page heap. These are bytes that have been released back to the OS, possibly by one of the MallocExtension "Release" calls. They can be used to fulfill allocation requests, but typically incur a page fault. They always count towards virtual memory usage, and depending on the OS, typically do not count towards physical memory usage.	             |  cluster, entityName | 
| kbdi_impala_tcmalloc_physical_bytes_reserved                              |  bytes           |  > 5.8         |  Derived metric computing the amount of physical memory (in bytes) used by the process, including that actually in use and free bytes reserved by tcmalloc. Does not include the tcmalloc metadata.	                                                                                                                                                                                                   |  cluster, entityName | 
| kbdi_impala_tcmalloc_total_bytes_reserved                                 |  bytes           |  > 5.8         |  Bytes of system memory reserved by TCMalloc.	                                                                                                                                                                                                                                                                                                                                                         |  cluster, entityName | 
| kbdi_impala_thrift_server_catalog_service_connections_in_use              |  connections     |  > 5.8         |  The active catalog server connections to this catalog server                                                                                                                                                                                                                                                                                                                                          |  cluster, entityName | 
| kbdi_impala_thrift_server_catalog_service_connections_rate                |  connections     |  > 5.8         |  The total number of connections made to this Catalog Server service catalog during its useful life.                                                                                                                                                                                                                                                                                                   |  cluster, entityName | 
| kbdi_impala_write_bytes_rate                                              |  bytes/s         |  > 5.8         |  The number of bytes written to the device	                                                                                                                                                                                                                                                                                                                                                           |  cluster, entityName | 




### Additional HDFS Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_hdfs_dfs_capacity_remaining | bytes | >= 5.8 | HDFS capacity remaining. | cluster, entityName |
| kbdi_hdfs_missing_blocks | blocks | >= 5.8 | HDFS missing blocks compatibility alias. | cluster, entityName |
| kbdi_hdfs_under_replicated_blocks | blocks | >= 5.8 | HDFS under replicated blocks compatibility alias. | cluster, entityName |
| kbdi_hdfs_corrupt_blocks | blocks | >= 5.8 | HDFS corrupt blocks. | cluster, entityName |
| kbdi_hdfs_namenode_jvm_heap_used_mb | MB | >= 5.8 | NameNode JVM heap used. | cluster, entityName |
| kbdi_hdfs_datanode_jvm_heap_used_mb | MB | >= 5.8 | DataNode JVM heap used. | cluster, entityName |

### Additional Impala Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_impala_num_queries | queries | >= 5.8 | Current number of Impala queries when exposed by CM. | cluster, entityName |
| kbdi_impala_num_sessions | sessions | >= 5.8 | Current number of Impala sessions when exposed by CM. | cluster, entityName |
| kbdi_impala_queries_spilled_rate | queries/s | >= 5.8 | Impala spilled query rate compatibility alias. | cluster, entityName |
| kbdi_impala_queries_timeout_rate | queries/s | >= 5.8 | Impala timed out query rate compatibility alias. | cluster, entityName |
| kbdi_impala_jvm_heap_used_mb | MB | >= 5.8 | JVM heap used by Impala roles. | cluster, entityName |

### HBase Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_hbase_service_health | state | >= 5.8 | HBase service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_hbase_role_health | state | >= 5.8 | HBase role health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_hbase_read_requests_rate | requests/s | version dependent | HBase read requests. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_write_requests_rate | requests/s | version dependent | HBase write requests. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_region_count | regions | version dependent | HBase region count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_store_file_count | files | version dependent | HBase store file count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_memstore_size | bytes | version dependent | HBase memstore size. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_compaction_queue_size | count | version dependent | HBase compaction queue size. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_flush_queue_size | count | version dependent | HBase flush queue size. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_read_latency | time | version dependent | HBase read latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_write_latency | time | version dependent | HBase write latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_jvm_heap_used_mb | MB | version dependent | HBase role JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_jvm_threads | threads | version dependent | HBase JVM threads. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_cpu_user_rate | seconds/s | version dependent | HBase role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_cpu_system_rate | seconds/s | version dependent | HBase role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_mem_rss | bytes | version dependent | HBase role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hbase_mem_virtual | bytes | version dependent | HBase role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Hive Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_hive_service_health | state | >= 5.8 | Hive or Hive-on-Tez service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_hive_role_health | state | >= 5.8 | Hive role health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_hive_open_connections | connections | version dependent | Hive open connections. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_open_sessions | sessions | version dependent | HiveServer2 open sessions. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_active_operations | operations | version dependent | Hive active operations. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_completed_operations | operations | version dependent | Hive completed operations. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_jvm_heap_used_mb | MB | version dependent | Hive role JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_jvm_threads | threads | version dependent | Hive JVM threads. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_cpu_user_rate | seconds/s | version dependent | Hive role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_cpu_system_rate | seconds/s | version dependent | Hive role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_mem_rss | bytes | version dependent | Hive role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_hive_mem_virtual | bytes | version dependent | Hive role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Kafka Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_kafka_service_health | state | >= 5.8 | Kafka service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_kafka_role_health | state | >= 5.8 | Kafka broker health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_kafka_bytes_in_rate | bytes/s | version dependent | Kafka bytes in. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_bytes_out_rate | bytes/s | version dependent | Kafka bytes out. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_messages_in_rate | messages/s | version dependent | Kafka messages in. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_under_replicated_partitions | partitions | version dependent | Kafka under replicated partitions. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_offline_partitions | partitions | version dependent | Kafka offline partitions. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_isr_shrinks_rate | events/s | version dependent | Kafka ISR shrinks. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_isr_expands_rate | events/s | version dependent | Kafka ISR expands. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_leader_election_rate | elections/s | version dependent | Kafka leader elections. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_request_handler_idle | percent | version dependent | Kafka request handler idle. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_network_processor_idle | percent | version dependent | Kafka network processor idle. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_jvm_heap_used_mb | MB | version dependent | Kafka JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_jvm_threads | threads | version dependent | Kafka JVM threads. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_cpu_user_rate | seconds/s | version dependent | Kafka role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_cpu_system_rate | seconds/s | version dependent | Kafka role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_mem_rss | bytes | version dependent | Kafka role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kafka_mem_virtual | bytes | version dependent | Kafka role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Kudu Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_kudu_service_health | state | >= 5.8 | Kudu service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_kudu_role_health | state | >= 5.8 | Kudu role health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_kudu_tablet_count | tablets | version dependent | Kudu tablet count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_leader_tablet_count | tablets | version dependent | Kudu leader tablet count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_replica_count | replicas | version dependent | Kudu replica count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_write_ops_rate | ops/s | version dependent | Kudu write operations. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_read_ops_rate | ops/s | version dependent | Kudu read operations. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_write_latency | time | version dependent | Kudu write latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_read_latency | time | version dependent | Kudu read latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_rpc_queue_length | count | version dependent | Kudu RPC queue length. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_jvm_heap_used_mb | MB | version dependent | Kudu JVM heap used when exposed by CM. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_cpu_user_rate | seconds/s | version dependent | Kudu role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_cpu_system_rate | seconds/s | version dependent | Kudu role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_mem_rss | bytes | version dependent | Kudu role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_kudu_mem_virtual | bytes | version dependent | Kudu role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Spark Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_spark_service_health | state | >= 5.8 | Spark service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_spark_role_health | state | >= 5.8 | Spark role health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_spark_apps_running | apps | version dependent | Spark running applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_apps_completed | apps | version dependent | Spark completed applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_apps_failed | apps | version dependent | Spark failed applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_executors_active | executors | version dependent | Spark active executors. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_drivers_active | drivers | version dependent | Spark active drivers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_memory_used | bytes | version dependent | Spark memory used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_cores_used | cores | version dependent | Spark cores used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_history_server_jvm_heap_used_mb | MB | version dependent | Spark History Server JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_cpu_user_rate | seconds/s | version dependent | Spark role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_cpu_system_rate | seconds/s | version dependent | Spark role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_mem_rss | bytes | version dependent | Spark role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_spark_mem_virtual | bytes | version dependent | Spark role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Yarn Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_yarn_service_health | state | >= 5.8 | YARN service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_yarn_role_health | state | >= 5.8 | YARN role health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_yarn_apps_running | apps | version dependent | YARN running applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_apps_pending | apps | version dependent | YARN pending applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_apps_failed | apps | version dependent | YARN failed applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_apps_killed | apps | version dependent | YARN killed applications. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_containers_allocated | containers | version dependent | YARN allocated containers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_containers_pending | containers | version dependent | YARN pending containers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_total_memory_mb | MB | version dependent | YARN total memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_allocated_memory_mb | MB | version dependent | YARN allocated memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_available_memory_mb | MB | version dependent | YARN available memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_total_vcores | vcores | version dependent | YARN total vcores. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_allocated_vcores | vcores | version dependent | YARN allocated vcores. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_available_vcores | vcores | version dependent | YARN available vcores. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_active_nodemanagers | nodes | version dependent | YARN active NodeManagers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_lost_nodemanagers | nodes | version dependent | YARN lost NodeManagers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_unhealthy_nodemanagers | nodes | version dependent | YARN unhealthy NodeManagers. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_jvm_heap_used_mb | MB | version dependent | YARN role JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_cpu_user_rate | seconds/s | version dependent | YARN role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_cpu_system_rate | seconds/s | version dependent | YARN role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_mem_rss | bytes | version dependent | YARN role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_yarn_mem_virtual | bytes | version dependent | YARN role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Zookeeper Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_zookeeper_service_health | state | >= 5.8 | ZooKeeper service health. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_zookeeper_role_health | state | >= 5.8 | ZooKeeper server health. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_zookeeper_outstanding_requests | requests | version dependent | ZooKeeper outstanding requests. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_znode_count | znodes | version dependent | ZooKeeper znode count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_watch_count | watches | version dependent | ZooKeeper watch count. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_num_alive_connections | connections | version dependent | ZooKeeper alive connections. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_packets_received_rate | packets/s | version dependent | ZooKeeper packets received. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_packets_sent_rate | packets/s | version dependent | ZooKeeper packets sent. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_latency_min | time | version dependent | ZooKeeper minimum latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_latency_avg | time | version dependent | ZooKeeper average latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_latency_max | time | version dependent | ZooKeeper maximum latency. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_jvm_heap_used_mb | MB | version dependent | ZooKeeper JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_cpu_user_rate | seconds/s | version dependent | ZooKeeper role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_cpu_system_rate | seconds/s | version dependent | ZooKeeper role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_mem_rss | bytes | version dependent | ZooKeeper role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_zookeeper_mem_virtual | bytes | version dependent | ZooKeeper role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### NiFi Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_nifi_service_health | state | version dependent | NiFi service health when exposed by CM or CSD. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_nifi_role_health | state | version dependent | NiFi role health when exposed by CM or CSD. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_nifi_flow_files_queued | files | version dependent | NiFi queued flow files. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_bytes_queued | bytes | version dependent | NiFi queued bytes. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_bytes_read_rate | bytes/s | version dependent | NiFi bytes read. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_bytes_written_rate | bytes/s | version dependent | NiFi bytes written. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_bytes_sent_rate | bytes/s | version dependent | NiFi bytes sent. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_bytes_received_rate | bytes/s | version dependent | NiFi bytes received. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_active_threads | threads | version dependent | NiFi active threads. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_garbage_collection_rate | collections/s | version dependent | NiFi garbage collection rate. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_jvm_heap_used_mb | MB | version dependent | NiFi JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_cpu_user_rate | seconds/s | version dependent | NiFi role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_cpu_system_rate | seconds/s | version dependent | NiFi role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_mem_rss | bytes | version dependent | NiFi role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_nifi_mem_virtual | bytes | version dependent | NiFi role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### Flink Module Metrics
| Metric Name | Unit | C.M. Version | Description | Metadata/Labels |
|-------------|:----:|:------------:|-------------|-----------------|
| kbdi_flink_service_health | state | version dependent | Flink service health when exposed by CM or CSD. | cluster, service_name, service_type, service_state, health_summary |
| kbdi_flink_role_health | state | version dependent | Flink role health when exposed by CM or CSD. | cluster, service_name, service_type, role_name, role_type, hostname, role_state, health_summary |
| kbdi_flink_jobs_running | jobs | version dependent | Flink running jobs. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_jobs_failed | jobs | version dependent | Flink failed jobs. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_jobs_finished | jobs | version dependent | Flink finished jobs. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_task_slots_total | slots | version dependent | Flink total task slots. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_task_slots_available | slots | version dependent | Flink available task slots. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_checkpoints_completed | checkpoints | version dependent | Flink completed checkpoints. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_checkpoints_failed | checkpoints | version dependent | Flink failed checkpoints. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_jvm_heap_used_mb | MB | version dependent | Flink JVM heap used. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_cpu_user_rate | seconds/s | version dependent | Flink role user CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_cpu_system_rate | seconds/s | version dependent | Flink role system CPU. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_mem_rss | bytes | version dependent | Flink role resident memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |
| kbdi_flink_mem_virtual | bytes | version dependent | Flink role virtual memory. | cluster, entityName, serviceName, serviceType, roleName, roleType, hostname |

### KBDI Metrics
| Metric Name | Unit           | Description                     | Metadata |
|-------------|:--------------:|---------------------------------|----------|
| kbdi_up     | [1-0] (OK-KO) | Keedio Big Data Insights Status | None     | 


