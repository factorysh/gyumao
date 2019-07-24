package config

type Config struct {
	Rules []*Rule `yaml:"rules"`
}

type Rule struct {
	Measurement string            `yaml:"measurement"`
	TagPass     map[string]string `yaml:"tag_pass"`
	TagExclude  map[string]string `yaml:"tag_exclude"`
	GroupBy     []string          `yaml:"group_by"`
	MaxAge      uint              `yaml:"max-age"`
	Expr        string            `yaml:"expr"`
}
