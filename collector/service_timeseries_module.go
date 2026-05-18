package collector

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	log "keedio/cloudera_exporter/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

var metricNameInvalidChars = regexp.MustCompile(`[^a-zA-Z0-9_:]`)
var metricNameRepeatedUnderscores = regexp.MustCompile(`_+`)

type serviceTimeseriesMetric struct {
	Name  string
	Help  string
	Query string
}

type serviceScraperSpec struct {
	Name         string
	Help         string
	ServiceTypes []string
	Metrics      []serviceTimeseriesMetric
}

var serviceTimeseriesLabels = []string{
	"cluster",
	"entityName",
	"serviceName",
	"serviceType",
	"roleName",
	"roleType",
	"hostname",
}

var serviceHealthLabels = []string{
	"cluster",
	"service_name",
	"service_type",
	"service_state",
	"health_summary",
}

var roleHealthLabels = []string{
	"cluster",
	"service_name",
	"service_type",
	"role_name",
	"role_type",
	"hostname",
	"role_state",
	"health_summary",
}

func sanitizeMetricName(name string) string {
	sanitized := strings.ToLower(strings.TrimSpace(name))
	sanitized = metricNameInvalidChars.ReplaceAllString(sanitized, "_")
	sanitized = metricNameRepeatedUnderscores.ReplaceAllString(sanitized, "_")
	sanitized = strings.Trim(sanitized, "_")
	if sanitized == "" {
		return "unknown"
	}
	if sanitized[0] >= '0' && sanitized[0] <= '9' {
		return "_" + sanitized
	}
	return sanitized
}

func serviceTypesPredicate(serviceTypes []string) string {
	if len(serviceTypes) == 0 {
		return ""
	}
	if len(serviceTypes) == 1 {
		return fmt.Sprintf(`serviceType = "%s"`, serviceTypes[0])
	}
	escaped := make([]string, 0, len(serviceTypes))
	for _, serviceType := range serviceTypes {
		escaped = append(escaped, regexp.QuoteMeta(serviceType))
	}
	return fmt.Sprintf(`serviceType rlike "^(%s)$"`, strings.Join(escaped, "|"))
}

func serviceMetricQuery(metricName string, serviceTypes []string) string {
	predicate := serviceTypesPredicate(serviceTypes)
	if predicate == "" {
		return fmt.Sprintf("SELECT LAST(%s) WHERE category=SERVICE", metricName)
	}
	return fmt.Sprintf("SELECT LAST(%s) WHERE category=SERVICE AND %s", metricName, predicate)
}

func roleMetricQuery(metricName string, serviceTypes []string, roleTypes ...string) string {
	predicates := []string{"category=ROLE"}
	if servicePredicate := serviceTypesPredicate(serviceTypes); servicePredicate != "" {
		predicates = append(predicates, servicePredicate)
	}
	if len(roleTypes) > 0 {
		escaped := make([]string, 0, len(roleTypes))
		for _, roleType := range roleTypes {
			escaped = append(escaped, regexp.QuoteMeta(roleType))
		}
		predicates = append(predicates, fmt.Sprintf(`roleType rlike "^(%s)$"`, strings.Join(escaped, "|")))
	}
	return fmt.Sprintf("SELECT LAST(%s) WHERE %s", metricName, strings.Join(predicates, " AND "))
}

func rateRoleMetricQuery(metricName string, serviceTypes []string, roleTypes ...string) string {
	return roleMetricQuery(fmt.Sprintf("INTEGRAL(%s)", metricName), serviceTypes, roleTypes...)
}

func rateServiceMetricQuery(metricName string, serviceTypes []string) string {
	return serviceMetricQuery(fmt.Sprintf("INTEGRAL(%s)", metricName), serviceTypes)
}

func commonRoleMetrics(serviceTypes []string, roleTypes ...string) []serviceTimeseriesMetric {
	return []serviceTimeseriesMetric{
		{
			Name:  "cpu_user_rate",
			Help:  "User CPU seconds per second for service roles.",
			Query: rateRoleMetricQuery("cpu_user_rate", serviceTypes, roleTypes...),
		},
		{
			Name:  "cpu_system_rate",
			Help:  "System CPU seconds per second for service roles.",
			Query: rateRoleMetricQuery("cpu_system_rate", serviceTypes, roleTypes...),
		},
		{
			Name:  "mem_rss",
			Help:  "Resident memory used by service roles in bytes.",
			Query: roleMetricQuery("mem_rss", serviceTypes, roleTypes...),
		},
		{
			Name:  "mem_virtual",
			Help:  "Virtual memory used by service roles in bytes.",
			Query: roleMetricQuery("mem_virtual", serviceTypes, roleTypes...),
		},
	}
}

func appendCommonRoleMetrics(metrics []serviceTimeseriesMetric, serviceTypes []string, roleTypes ...string) []serviceTimeseriesMetric {
	return append(metrics, commonRoleMetrics(serviceTypes, roleTypes...)...)
}

func scrapeServiceModule(ctx context.Context, config Collector_connection_data, spec serviceScraperSpec, ch chan<- prometheus.Metric) error {
	log.Debug_msg("Executing %s Metrics Scraper", spec.Name)

	successQueries := 0
	errorQueries := 0

	if scrapeServiceAndRoleHealth(ctx, config, spec, ch) {
		successQueries++
	} else {
		errorQueries++
	}

	for _, metric := range spec.Metrics {
		if scrapeServiceTimeseriesMetric(ctx, config, spec, metric, ch) {
			successQueries++
		} else {
			errorQueries++
		}
	}

	log.Debug_msg("In the %s Module has been executed %d queries. %d success and %d with errors", spec.Name, successQueries+errorQueries, successQueries, errorQueries)
	return nil
}

func scrapeServiceTimeseriesMetric(ctx context.Context, config Collector_connection_data, spec serviceScraperSpec, metric serviceTimeseriesMetric, ch chan<- prometheus.Metric) bool {
	if strings.TrimSpace(metric.Query) == "" {
		return true
	}

	jsonParsed, err := make_and_parse_timeseries_query(ctx, config, metric.Query)
	if err != nil {
		log.Warn_msg("Skipping %s metric %s after query error: %s", spec.Name, metric.Name, err.Error())
		return false
	}

	series := getTimeseriesSeries(jsonParsed)
	if len(series) == 0 {
		return true
	}

	desc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, sanitizeMetricName(spec.Name), sanitizeMetricName(metric.Name)),
		metric.Help,
		serviceTimeseriesLabels,
		nil,
	)

	for _, serie := range series {
		value, ok := latestTimeseriesValue(serie)
		if !ok {
			continue
		}

		labels := timeseriesLabels(serie)
		ch <- prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			value,
			labels["cluster"],
			labels["entityName"],
			labels["serviceName"],
			labels["serviceType"],
			labels["roleName"],
			labels["roleType"],
			labels["hostname"],
		)
	}
	return true
}

func scrapeServiceAndRoleHealth(ctx context.Context, config Collector_connection_data, spec serviceScraperSpec, ch chan<- prometheus.Metric) bool {
	clusters, err := make_and_parse_api_query(ctx, config, "clusters?view=full")
	if err != nil {
		log.Warn_msg("Skipping %s service status after cluster discovery error: %s", spec.Name, err.Error())
		return false
	}

	serviceDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, sanitizeMetricName(spec.Name), "service_health"),
		fmt.Sprintf("%s service health summary as a numeric value.", spec.Name),
		serviceHealthLabels,
		nil,
	)
	roleDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, sanitizeMetricName(spec.Name), "role_health"),
		fmt.Sprintf("%s role health summary as a numeric value.", spec.Name),
		roleHealthLabels,
		nil,
	)

	for _, cluster := range clusters.Get("items").Array() {
		clusterName := firstNonEmpty(cluster.Get("name").String(), cluster.Get("displayName").String())
		if clusterName == "" {
			continue
		}

		services, err := make_and_parse_api_query(ctx, config, fmt.Sprintf("clusters/%s/services?view=full", url.PathEscape(clusterName)))
		if err != nil {
			log.Warn_msg("Skipping %s services for cluster %s after query error: %s", spec.Name, clusterName, err.Error())
			continue
		}

		for _, service := range services.Get("items").Array() {
			serviceType := service.Get("type").String()
			if !containsString(spec.ServiceTypes, serviceType) {
				continue
			}
			serviceName := service.Get("name").String()
			serviceState := service.Get("serviceState").String()
			healthSummary := service.Get("healthSummary").String()
			ch <- prometheus.MustNewConstMetric(
				serviceDesc,
				prometheus.GaugeValue,
				get_value_from_state(healthSummary),
				clusterName,
				serviceName,
				serviceType,
				serviceState,
				healthSummary,
			)

			scrapeRoleHealth(ctx, config, clusterName, serviceName, serviceType, roleDesc, ch)
		}
	}

	return true
}

func scrapeRoleHealth(ctx context.Context, config Collector_connection_data, clusterName string, serviceName string, serviceType string, roleDesc *prometheus.Desc, ch chan<- prometheus.Metric) {
	roles, err := make_and_parse_api_query(ctx, config, fmt.Sprintf("clusters/%s/services/%s/roles?view=full", url.PathEscape(clusterName), url.PathEscape(serviceName)))
	if err != nil {
		log.Warn_msg("Skipping roles for service %s in cluster %s after query error: %s", serviceName, clusterName, err.Error())
		return
	}

	for _, role := range roles.Get("items").Array() {
		roleName := role.Get("name").String()
		roleType := role.Get("type").String()
		hostname := firstNonEmpty(role.Get("hostRef.hostname").String(), role.Get("hostRef.hostName").String())
		roleState := role.Get("roleState").String()
		healthSummary := role.Get("healthSummary").String()

		ch <- prometheus.MustNewConstMetric(
			roleDesc,
			prometheus.GaugeValue,
			get_value_from_state(healthSummary),
			clusterName,
			serviceName,
			serviceType,
			roleName,
			roleType,
			hostname,
			roleState,
			healthSummary,
		)
	}
}

func getTimeseriesSeries(jsonParsed gjson.Result) []gjson.Result {
	if items := jsonParsed.Get("items").Array(); len(items) > 0 {
		series := make([]gjson.Result, 0)
		for _, item := range items {
			series = append(series, item.Get("timeSeries").Array()...)
		}
		return series
	}
	return jsonParsed.Get("timeSeries").Array()
}

func latestTimeseriesValue(serie gjson.Result) (float64, bool) {
	data := serie.Get("data").Array()
	for i := len(data) - 1; i >= 0; i-- {
		valueResult := data[i].Get("value")
		if !valueResult.Exists() || valueResult.Type == gjson.Null {
			continue
		}
		if valueResult.Type == gjson.Number {
			return valueResult.Float(), true
		}
		value, err := strconv.ParseFloat(valueResult.String(), 64)
		if err == nil {
			return value, true
		}
	}
	return 0, false
}

func timeseriesLabels(serie gjson.Result) map[string]string {
	attributes := serie.Get("metadata.attributes")
	return map[string]string{
		"cluster":     firstNonEmpty(attributes.Get("clusterName").String(), attributes.Get("clusterDisplayName").String(), attributes.Get("cluster").String()),
		"entityName":  attributes.Get("entityName").String(),
		"serviceName": attributes.Get("serviceName").String(),
		"serviceType": attributes.Get("serviceType").String(),
		"roleName":    firstNonEmpty(attributes.Get("roleName").String(), attributes.Get("role").String()),
		"roleType":    attributes.Get("roleType").String(),
		"hostname":    attributes.Get("hostname").String(),
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func containsString(values []string, needle string) bool {
	for _, value := range values {
		if value == needle {
			return true
		}
	}
	return false
}
