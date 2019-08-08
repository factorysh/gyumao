package gyumao

import (
	"net/http"

	"github.com/factorysh/gyumao/config"
	"github.com/factorysh/gyumao/plugin"
	"github.com/factorysh/gyumao/point"
	"github.com/factorysh/gyumao/rule"
	log "github.com/sirupsen/logrus"
)

// Gyumao main object
type Gyumao struct {
	plugins *plugin.Plugins
	cfg     *config.Config
	rules   rule.Rules
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
	return g, nil
}

// Serve and block
func (g *Gyumao) Serve() error {
	crusher := point.New(g.rules, g.plugins.EvaluatorPlugins())
	http.Handle("/write", crusher)
	go crusher.Start()
	log.Info("HTTP Listen", g.cfg.InfluxdbListen)
	return http.ListenAndServe(g.cfg.InfluxdbListen, nil)
}
