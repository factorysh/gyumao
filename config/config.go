package config

// Config is the main config object
type Config struct {
	InfluxdbListen string                            `yaml:"influxdb_listen"`
	Rules          []*Rule                           `yaml:"rules"`
	PluginFolder   string                            `yaml:"plugin_folder"`
	Plugins        map[string]map[string]interface{} `yaml:"plugins"`
}

// Rule describes what to do with Influxdb events
type Rule struct {
	Measurement string              `yaml:"measurement"`
	TagsPass    map[string][]string `yaml:"tags_pass"`
	TagsExclude map[string][]string `yaml:"tags_exclude"`
	GroupBy     []string            `yaml:"group_by"`
	MaxAge      uint                `yaml:"max-age"`
	Expr        string              `yaml:"expr"`
}
