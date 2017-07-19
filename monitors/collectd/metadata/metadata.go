package metadata

//go:generate collectd-template-to-go metadata.tmpl

import (
	"github.com/signalfx/neo-agent/core/config"
	"github.com/signalfx/neo-agent/monitors"
	"github.com/signalfx/neo-agent/monitors/collectd"
)

const monitorType = "collectd/signalfx-metadata"

func init() {
	monitors.Register(monitorType, func() interface{} {
		return &MetadataMonitor{
			*collectd.NewStaticMonitorCore(CollectdTemplate),
		}
	}, &Config{})
}

type Config struct {
	config.MonitorConfig
}

type MetadataMonitor struct {
	collectd.StaticMonitorCore
}

func (m *MetadataMonitor) Configure(conf *Config) bool {
	return m.SetConfigurationAndRun(conf.MonitorConfig)
}
