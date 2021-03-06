// Code generated by go generate; DO NOT EDIT.
package basic

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = []Metric{
	{
		Name: "BurstBalance",
		Desc: prometheus.NewDesc(
			"aws_rds_burst_balance_average",
			"The percent of General Purpose SSD (gp2) burst-bucket I/O credits available. Units: Percent",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUCreditBalance",
		Desc: prometheus.NewDesc(
			"aws_rds_cpu_credit_balance_average",
			"[T2 instances] The number of CPU credits available for the instance to burst beyond its base CPU utilization. Credits are stored in the credit balance after they are earned and removed from the credit balance after they expire. Credits expire 24 hours after they are earned. CPU credit metrics are available only at a 5 minute frequency. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUCreditUsage",
		Desc: prometheus.NewDesc(
			"aws_rds_cpu_credit_usage_average",
			"[T2 instances] The number of CPU credits consumed by the instance. One CPU credit equals one vCPU running at 100% utilization for one minute or an equivalent combination of vCPUs, utilization, and time (for example, one vCPU running at 50% utilization for two minutes or two vCPUs running at 25% utilization for two minutes). CPU credit metrics are available only at a 5 minute frequency. If you specify a period greater than five minutes, use the Sum statistic instead of the Average statistic. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "CPUUtilization",
		Desc: prometheus.NewDesc(
			"node_cpu_average",
			"The percentage of CPU utilization. Units: Percent",
			[]string{"instance", "region"},
			map[string]string{"cpu": "All", "mode": "total"},
		),
	},
	{
		Name: "DatabaseConnections",
		Desc: prometheus.NewDesc(
			"aws_rds_database_connections_average",
			"The number of database connections in use. Units: Count",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "FreeStorageSpace",
		Desc: prometheus.NewDesc(
			"node_filesystem_free",
			"The amount of available storage space. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "FreeableMemory",
		Desc: prometheus.NewDesc(
			"node_memory_Cached",
			"The amount of available random access memory. Units: Bytes",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReadLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_read_latency_average",
			"The average amount of time taken per disk I/O operation. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "ReplicaLag",
		Desc: prometheus.NewDesc(
			"aws_rds_replica_lag_average",
			"The amount of time a Read Replica DB instance lags behind the source DB instance. Applies to MySQL, MariaDB, and PostgreSQL Read Replicas. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	{
		Name: "WriteLatency",
		Desc: prometheus.NewDesc(
			"aws_rds_write_latency_average",
			"The average amount of time taken per disk I/O operation. Units: Seconds",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
}
