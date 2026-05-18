/*
 *
 * title           :consultor.go
 * description     :File with the common code to all the Scrapers
 * author		       :Alejandro Villegas Lopez (avillegas@keedio.com)
 * date            :2018/10/05
 * version         :0.1
 *
 */
package collector

/* ======================================================================
 * Dependencies and libraries
 * ====================================================================== */
import (
	// Go Default libraries
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	// Own libraries
	jp "keedio/cloudera_exporter/json_parser"
	log "keedio/cloudera_exporter/logger"

	// Go Prometheus libraries
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

/* ======================================================================
 * Constants
 * ====================================================================== */
const MASTER_POS = 0
const BORDER_POS = 1
const WORKER_POS = 2
const maxTimeseriesGetURLLength = 1800

/* ======================================================================
 * Data Structs
 * ====================================================================== */
// Structure to relate the sentence of TSquery with its metric of Prometheus
type relation struct {
	Query         string
	Metric_struct prometheus.Desc
}

type timeseriesAPIRequest struct {
	Query                string `json:"query"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	ContentType          string `json:"contentType"`
	DesiredRollup        string `json:"desiredRollup"`
	MustUseDesiredRollup bool   `json:"mustUseDesiredRollup"`
}

/* ======================================================================
 * Functions
 * ====================================================================== */
// Make the query specified to the Cloudera Manager API and returns the JSON response
func make_query(ctx context.Context, uri string, user string, passwd string) (body string, err error) {
	log.Debug_msg("Making API Query: %s ", uri)

	// Get HTTP Protocol Client
	httpClient := http.DefaultClient

	// Build the request Object
	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		log.Err_msg("Building Request for URL:%s, Failed. Error: %s", uri, err)
		return "", err
	}

	// Overwrite request with timeout context.
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// Request response header
	req.Header.Add("Content-Type", "application/json")

	// Set Authentication credentials
	req.SetBasicAuth(user, passwd)

	// Make the API request
	res, err := httpClient.Do(req)
	if err != nil {
		log.Err_msg("%s", err)
		return "", err
	}
	if res == nil {
		log.Err_msg("HTTP response is NULL")
		return "", errors.New("HTTP response is NULL")
	}
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		log.Err_msg("Invalid HTTP response code: %s for the request: %s", res.Status, uri)
		res.Body.Close()
		if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
			return "", fmt.Errorf("Cloudera Manager API authorization failed: %s", res.Status)
		}
		if res.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("Cloudera Manager API resource not found: %s", res.Status)
		}
		return "", fmt.Errorf("Invalid HTTP response code: %s", res.Status)
	}

	// Get Body Response
	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Err_msg("Failed to parse response with error: %s", err)
		res.Body.Close()
		return "", err
	}

	res.Body.Close()

	return string(content), err
}

func make_timeseries_query(ctx context.Context, config Collector_connection_data, query string) (body string, err error) {
	now := time.Now().UTC()
	requestData := timeseriesAPIRequest{
		Query:                query,
		From:                 now.Add(-5 * time.Minute).Format(time.RFC3339),
		To:                   now.Format(time.RFC3339),
		ContentType:          "application/json",
		DesiredRollup:        "RAW",
		MustUseDesiredRollup: false,
	}

	endpoint := fmt.Sprintf("http://%s:%s/api/%s/timeseries", config.Host, config.Port, config.Api_version)
	values := url.Values{}
	values.Set("query", requestData.Query)
	values.Set("from", requestData.From)
	values.Set("to", requestData.To)
	values.Set("contentType", requestData.ContentType)
	values.Set("desiredRollup", requestData.DesiredRollup)
	values.Set("mustUseDesiredRollup", strconv.FormatBool(requestData.MustUseDesiredRollup))

	uri := fmt.Sprintf("%s?%s", endpoint, values.Encode())
	if len(uri) <= maxTimeseriesGetURLLength || !api_version_at_least(config.Api_version, 11) {
		return make_query(ctx, uri, config.User, config.Passwd)
	}

	log.Debug_msg("Making API Query: POST %s", endpoint)
	payload, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(payload))
	if err != nil {
		log.Err_msg("Building Request for URL:%s, Failed. Error: %s", endpoint, err)
		return "", err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(config.User, config.Passwd)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Err_msg("%s", err)
		return "", err
	}
	if res == nil {
		log.Err_msg("HTTP response is NULL")
		return "", errors.New("HTTP response is NULL")
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		log.Err_msg("Invalid HTTP response code: %s for the request: POST %s", res.Status, endpoint)
		if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
			return "", fmt.Errorf("Cloudera Manager API authorization failed: %s", res.Status)
		}
		if res.StatusCode == http.StatusNotFound {
			return "", fmt.Errorf("Cloudera Manager API resource not found: %s", res.Status)
		}
		return "", fmt.Errorf("Invalid HTTP response code: %s", res.Status)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Err_msg("Failed to parse response with error: %s", err)
		return "", err
	}
	return string(content), nil
}

func api_version_at_least(version string, minimum int) bool {
	trimmed := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(version)), "v")
	value, err := strconv.Atoi(trimmed)
	if err != nil {
		return false
	}
	return value >= minimum
}

// Create a empty map to storage the host_id as Key and a list of flags for Border, Worker or Master Host Role
func init_host_types_map(ctx context.Context, config Collector_connection_data) map[string][]string {
	node_map := make(map[string][]string)

	// Get Hosts list
	json_hosts_data, _ := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			fmt.Sprintf("hosts")),
		config.User,
		config.Passwd,
	)
	json_hosts_results := jp.Parse_json_response(json_hosts_data)
	num_hosts, _ := strconv.Atoi(jp.Get_json_field(json_hosts_results, "items.#"))
	for host_index := 0; host_index < num_hosts; host_index++ {
		host_id := jp.Get_json_field(json_hosts_results, fmt.Sprintf("items.%d.hostId", host_index))
		node_map[host_id] = []string{"0", "0", "0"}
	}
	return node_map
}

// Activate the flags for the border nodes
func look_for_border_nodes(ctx context.Context, config Collector_connection_data, cluster_name string, node_map map[string][]string) map[string][]string {
	json_type_data, _ := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			fmt.Sprintf("clusters/%s/services/hdfs/roles", cluster_name)),
		config.User,
		config.Passwd,
	)

	// Parse JSON Response
	json_type_results := jp.Parse_json_response(json_type_data)

	// For each Host
	num_hosts, _ := strconv.Atoi(jp.Get_json_field(json_type_results, "items.#"))
	for host_index := 0; host_index < num_hosts; host_index++ {
		host_id := jp.Get_json_field(json_type_results, fmt.Sprintf("items.%d.hostRef.hostId", host_index))
		host_type := jp.Get_json_field(json_type_results, fmt.Sprintf("items.%d.type", host_index))
		_, ok := node_map[host_id]
		if strings.Contains(host_type, "GATEWAY") && ok {
			node_map[host_id][BORDER_POS] = "1"
		}
	}
	return node_map
}

// Activate the flags for the worker nodes
func look_for_worker_nodes(ctx context.Context, config Collector_connection_data, cluster_name string, node_map map[string][]string) map[string][]string {
	json_type_data, _ := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			fmt.Sprintf("clusters/%s/services/hdfs/roles", cluster_name)),
		config.User,
		config.Passwd,
	)

	// Parse JSON Response
	json_type_results := jp.Parse_json_response(json_type_data)

	// For each Host
	num_hosts, _ := strconv.Atoi(jp.Get_json_field(json_type_results, "items.#"))
	for host_index := 0; host_index < num_hosts; host_index++ {
		host_id := jp.Get_json_field(json_type_results, fmt.Sprintf("items.%d.hostRef.hostId", host_index))
		host_type := jp.Get_json_field(json_type_results, fmt.Sprintf("items.%d.type", host_index))
		if strings.Contains(host_type, "DATANODE") {
			node_map[host_id][WORKER_POS] = "1"
		}
	}
	return node_map
}

// Activate the flags for the master nodes
func look_for_master_nodes(ctx context.Context, config Collector_connection_data, cluster_name string, node_map map[string][]string) map[string][]string {
	json_master_data, _ := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			fmt.Sprintf("cm/service/roles")),
		config.User,
		config.Passwd,
	)

	// Parse JSON Response
	json_master_results := jp.Parse_json_response(json_master_data)

	// For each Host
	num_hosts, _ := strconv.Atoi(jp.Get_json_field(json_master_results, "items.#"))
	for host_index := 0; host_index < num_hosts; host_index++ {
		host_id := jp.Get_json_field(json_master_results, fmt.Sprintf("items.%d.hostRef.hostId", host_index))
		host_type := jp.Get_json_field(json_master_results, fmt.Sprintf("items.%d.serviceRef.serviceName", host_index))
		if strings.Contains(host_type, "mgmt") {
			node_map[host_id][MASTER_POS] = "1"
		}
	}
	return node_map
}

// fill and return the role map of hosts
func get_type_node_list(ctx context.Context, config Collector_connection_data) map[string][]string {
	node_map := init_host_types_map(ctx, config)

	// Get Cluster list
	json_clusters_data, _ := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			fmt.Sprintf("clusters")),
		config.User,
		config.Passwd,
	)

	// Parse JSON Response
	json_clusters_results := jp.Parse_json_response(json_clusters_data)
	// For each Cluster
	num_clusters, _ := strconv.Atoi(jp.Get_json_field(json_clusters_results, "items.#"))
	for cluster_index := 0; cluster_index < num_clusters; cluster_index++ {
		cluster_name := jp.Get_json_field(json_clusters_results, fmt.Sprintf("items.%d.name", cluster_index))
		node_map = look_for_border_nodes(ctx, config, cluster_name, node_map)
		node_map = look_for_master_nodes(ctx, config, cluster_name, node_map)
		node_map = look_for_worker_nodes(ctx, config, cluster_name, node_map)
	}
	return node_map
}

// Return the is_master flag
func get_if_is_master(host_id string) string {
	return string(type_node_list[host_id][MASTER_POS])
}

// Return the is_border flag
func get_if_is_border(host_id string) string {
	return type_node_list[host_id][BORDER_POS]
}

// Return the is_worker flag
func get_if_is_worker(host_id string) string {
	return type_node_list[host_id][WORKER_POS]
}

// Make the query and parse the json response.
func make_and_parse_timeseries_query(ctx context.Context, config Collector_connection_data, query string) (result gjson.Result, err error) {
	// Make query
	json_timeseries, err := make_timeseries_query(ctx, config, query)

	// parse and return the result
	if err != nil {
		log.Err_msg("Error making query: %s", err)
	}
	return jp.Parse_json_response(json_timeseries), err
}

// Make and parse a Cloudera API Query
func make_and_parse_api_query(ctx context.Context, config Collector_connection_data, query string) (result gjson.Result, err error) {
	// Make query
	json_timeseries, err := make_query(
		ctx,
		jp.Build_api_query_url(
			config.Host,
			config.Port,
			config.Api_version,
			query),
		config.User,
		config.Passwd,
	)

	// parse and return the result
	return jp.Parse_json_response(json_timeseries), err
}

// Returns a string with the Cloudera Manager version
func get_cloudera_manager_version(ctx context.Context, config Collector_connection_data) string {
	// Make query
	json_parsed, err := make_and_parse_api_query(ctx, config, "cm/version")
	if err != nil {
		return ""
	}
	return jp.Get_api_query_cm_version(json_parsed)
}

// Returns a string with the highest version of the Cloudera API
func Get_api_cloudera_version(ctx context.Context, config Collector_connection_data) (string, error) {
	// Make query
	json_parsed, err := make_query(
		ctx,
		fmt.Sprintf("http://%s:%s/api/version", config.Host, config.Port),
		config.User,
		config.Passwd,
	)
	if err != nil {
		return "", errors.New("The exporter can not determine the API version by consulting the cloudera Manager API")
	}
	return json_parsed, nil
}
