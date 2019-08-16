package gyumao

import (
	"fmt"
	"net/http"
	"time"

	"github.com/factorysh/gyumao/config"
	"github.com/factorysh/gyumao/crusher"
	"github.com/factorysh/gyumao/deadman"
	evaluator "github.com/factorysh/gyumao/evaluator/consumer"
	"github.com/factorysh/gyumao/plugin"
	"github.com/factorysh/gyumao/point"
	"github.com/factorysh/gyumao/probes"
	"github.com/factorysh/gyumao/rule"
	log "github.com/sirupsen/logrus"
)

// Gyumao main object
type Gyumao struct {
	plugins  *plugin.Plugins
	cfg      *config.Config
	rules    rule.Rules
	consumer point.Consumer
	probes   probes.Probes
}

// New Gyumao instance
func New(cfg *config.Config) (*Gyumao, error) {
	g := &Gyumao{
		plugins: plugin.NewPlugins(),
		cfg:     cfg,
	}
	err := g.plugins.RegisterAll(cfg.PluginFolder, cfg.Plugins)
	if err != nil {
		return nil, err
	}
	r, err := rule.FromRules(cfg.Rules...)
	if err != nil {
		return nil, err
	}
	g.rules = r
	if len(cfg.Probes) != 1 {
		return nil, fmt.Errorf(`Just one "probes", not %d`, len(cfg.Probes))
	}
	for k, v := range cfg.Probes {
		factory, ok := probes.ProbesPlugin[k]
		if !ok {
			return nil, fmt.Errorf(`Unknown probe : %s`, k)
		}
		var err error
		g.probes, err = factory(v)
		if err != nil {
			return nil, err
		}
	}
	return g, nil
}

// Serve and block
func (g *Gyumao) Serve() error {
	global := g.plugins.EvaluatorPlugins()
	eval := evaluator.NewConsumer(global)
	dead := deadman.NewConsumer(deadman.New(3, g.probes.Keys(), 30*time.Second))
	consumers := point.NewMultiConsumer(dead, eval)
	crusher := crusher.New(g.rules, consumers)
	http.Handle("/write", crusher)
	go crusher.Start()
	log.Info("HTTP Listen", g.cfg.InfluxdbListen)
	return http.ListenAndServe(g.cfg.InfluxdbListen, nil)
}
