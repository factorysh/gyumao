package config

import "time"

// Config is the main config object
type Config struct {
	InfluxdbListen string                            `yaml:"influxdb_listen"`
	Rules          []*Rule                           `yaml:"rules"`
	PluginFolder   string                            `yaml:"plugin_folder"`
	Plugins        map[string]map[string]interface{} `yaml:"plugins"`
	Probes         map[string]map[string]interface{} `yaml:"probes"`
	Deadman        *Deadman                          `yaml:"deadman"`
}

func (c *Config) Default() {
	if c.InfluxdbListen == "" {
		c.InfluxdbListen = ":8066"
	}
	if c.Rules == nil {
		c.Rules = make([]*Rule, 0)
	}
	if c.PluginFolder == "" {
		c.PluginFolder = "/var/lib/gyumao/plugins"
	}
	if c.Probes == nil {
		c.Probes = make(map[string]map[string]interface{})
		c.Probes["file"] = map[string]interface{}{
			"path": "/var/lib/gyumao/probes.yml",
		}
	}
	if c.Deadman == nil {
		c.Deadman = &Deadman{
			Duration:    30 * time.Second,
			Generations: 3,
		}
	}
}

type ProbeFileConfig struct {
	Path string `yaml:"path"`
}

// Rule describes what to do with Influxdb events
type Rule struct {
	Measurement string              `yaml:"measurement"`
	TagsPass    map[string][]string `yaml:"tags_pass"`
	TagsExclude map[string][]string `yaml:"tags_exclude"`
	GroupBy     []string            `yaml:"group_by"`
	MaxAge      uint                `yaml:"max_age"`
	Expr        string              `yaml:"expr"`
}

// Probes is a collection of probe
type Probes []Probe

// Probe is a unique probe : a measurement and tags
type Probe string

// Deadman settings
type Deadman struct {
	Duration    time.Duration `yaml:"duration"`
	Generations int           `yaml:"generations"`
}
