package config

import (
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"gopkg.in/yaml.v2"
)

type Rule struct {
	Name string            `yaml:"name"`
	Tags map[string]string `yaml:"tags"`
	Keys []string          `yaml:"keys"`
	Expr string            `yaml:"expr"`
	prog *vm.Program
}

func (r *Rule) Do(env map[string]interface{}) (interface{}, error) {
	return expr.Run(r.prog, env)
}

type Rules struct {
	Rules []*Rule `yaml:"rules"`
}

func Load(raw []byte) (*Rules, error) {
	var rules Rules
	err := yaml.Unmarshal(raw, &rules)
	if err != nil {
		return nil, err
	}
	for _, rule := range rules.Rules {
		prog, err := expr.Compile(rule.Expr)
		if err != nil {
			return nil, err
		}
		rule.prog = prog
	}
	return &rules, nil
}
