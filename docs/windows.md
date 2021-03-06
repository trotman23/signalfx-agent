# Windows Setup

The agent supports standalone installations on Windows Server 2008 and above.

## Installation

To get started deploying the Smart Agent directly on a Windows host, see the
[Quick Install](./quick-install.md) guide.

## Monitor Support

The agent supports all of the python based collectd monitors as well as all
Go-based monitors.  C-based collectd monitors are not available on Windows.
The following monitors are available on Windows.

- [aspdotnet](./monitors/aspdotnet.md)
- [collectd/consul](./monitors/collectd-consul.md)
- [collectd/couchbase](./monitors/collectd-couchbase.md)
- [collectd/etcd](./monitors/collectd-etcd.md)
- [collectd/hadoop](./monitors/collectd-hadoop.md)
- [collectd/health-checker](./monitors/collectd-health-checker.md)
- [collectd/jenkins](./monitors/collectd-jenkins.md)
- [collectd/kong](./monitors/collectd-kong.md)
- [collectd/marathon](./monitors/collectd-marathon.md)
- [collectd/mongodb](./monitors/collectd-mongodb.md)
- [collectd/openstack](./monitors/collectd-openstack.md)
- [collectd/python](./monitors/collectd-python.md)
- [collectd/rabbitmq](./monitors/collectd-rabbitmq.md)
- [collectd/redis](./monitors/collectd-redis.md)
- [collectd/spark](./monitors/collectd-spark.md)
- [collectd/zookeeper](./monitors/collectd-zookeeper.md)
- [conviva](./monitors/conviva.md)
- [cpu](./monitors/cpu.md)
- [disk-io](./monitors/disk-io.md)
- [docker-container-stats](./monitors/docker-container-stats.md)
- [dotnet](./monitors/dotnet.md)
- [elasticsearch](./monitors/elasticsearch.md)
- [filesystems](./monitors/filesystems.md)
- [haproxy](./monitors/haproxy.md)
- [host-metadata](./monitors/host-metadata.md)
- [internal-metrics](./monitors/internal-metrics.md)
- [memory](./monitors/memory.md)
- [net-io](./monitors/net-io.md)
- [processlist](./monitors/processlist.md)
- [prometheus-exporter](./monitors/prometheus-exporter.md)
- [telegraf/logparser](./monitors/telegraf-logparser.md)
- [telegraf/procstat](./monitors/telegraf-procstat.md)
- [telegraf/snmp](./monitors/telegraf-snmp.md)
- [telegraf/sqlserver](./monitors/telegraf-sqlserver.md)
- [telegraf/statsd](./monitors/telegraf-statsd.md)
- [telegraf/tail](./monitors/telegraf-tail.md)
- [telegraf/win_perf_counters](./monitors/telegraf-win_perf_counters.md)
- [telegraf/win_services](./monitors/telegraf-win_services.md)
- [trace-forwarder](./monitors/trace-forwarder.md)
- [vmem](./monitors/vmem.md)
- [windows-iis](./monitors/windows-iis.md)
- [windows-legacy](./monitors/windows-legacy.md)

## Logging

The agent logs to the console when executed directly in a powershell.  If it is configured
as a service and the `-logEvents` flag is supplied, the agent will log messages as
Windows Events Log.  Use the Event Viewer application to read the logs.  The Event
Viewer is located under `Start > Administrative Tools > Event Viewer`.  You can
see logged events from the agent service under `Windows Logs > Application`.

## Service Configuration

The agent can be registered as a Windows service by invoking the agent executable
with a few command line flags in powershell.

- Install Service

		PS> SignalFx\SignalFxAgent\bin\signalfx-agent.exe -service "install" -logEvents -config <path to config file>

- Start Service

		PS> SignalFx\SignalFxAgent\bin\signalfx-agent.exe -service "start"

- Stop Service

		PS> SignalFx\SignalFxAgent\bin\signalfx-agent.exe -service "stop"

- Uninstall Service

		PS> SignalFx\SignalFxAgent\bin\signalfx-agent.exe -service "uninstall"
