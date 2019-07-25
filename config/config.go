package config

type Config struct {
	Rules []*Rule `yaml:"rules"`
}

type Rule struct {
	Measurement string            `yaml:"measurement"`
	TagsPass     map[string][]string `yaml:"tags_pass"`
	TagsExclude  map[string][]string `yaml:"tags_exclude"`
	GroupBy     []string          `yaml:"group_by"`
	MaxAge      uint              `yaml:"max-age"`
	Expr        string            `yaml:"expr"`
}
