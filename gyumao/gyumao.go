package gyumao

import (
	"fmt"
	"net/http"

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
	l := log.WithField("cfg", cfg)
	g := &Gyumao{
		plugins: plugin.NewPlugins(),
		cfg:     cfg,
	}
	err := g.plugins.RegisterAll(cfg.PluginFolder, cfg.Plugins)
	if err != nil {
		l.WithError(err).Error("Plugin Register")
		return nil, err
	}
	r, err := rule.FromRules(cfg.Rules...)
	if err != nil {
		l.WithError(err).Error("Rules Register")
		return nil, err
	}
	g.rules = r
	if len(cfg.Probes) != 1 {
		err = fmt.Errorf(`Just one "probes", not %d`, len(cfg.Probes))
		l.WithError(err).Error("One probe")
		return nil, err
	}
	for k, v := range cfg.Probes {
		factory, ok := probes.ProbesPlugin[k]
		if !ok {
			err = fmt.Errorf(`Unknown probe : %s`, k)
			l.WithError(err).Error()
			return nil, err
		}
		var err error
		g.probes, err = factory(v)
		if err != nil {
			l.WithError(err).Error("Probes factory")
			return nil, err
		}
	}
	l.WithField("probes", g.probes).Debug()
	return g, nil
}

// Serve and block
func (g *Gyumao) Serve() error {
	global := g.plugins.EvaluatorPlugins()
	eval := evaluator.NewConsumer(global)
	dead := deadman.NewConsumer(deadman.New(g.cfg.Deadman.Generations,
		g.probes.Keys(), g.cfg.Deadman.Duration))
	consumers := point.NewMultiConsumer(dead, eval)
	crusher := crusher.New(g.rules, consumers, g.probes)
	http.Handle("/write", crusher)
	go crusher.Start()
	log.Info("HTTP Listen", g.cfg.InfluxdbListen)
	return http.ListenAndServe(g.cfg.InfluxdbListen, nil)
}
