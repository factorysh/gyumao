package main

import (
	"os"
	"time"

	_plugin "github.com/factorysh/gyumao/plugin"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var version = "0.0.0"

type WorkingHours struct {
	logger hclog.Logger
	config map[string]interface{}
}

func (w *WorkingHours) Time(t time.Time) (_plugin.Tags, error) {
	tags := make(_plugin.Tags)
	if t.Hour() < 8 || t.Hour() > 18 {
		tags["hours"] = "not"
		return tags, nil
	}
	tags["hours"] = "working"
	return tags, nil
}

func (w *WorkingHours) Meta() (_plugin.Meta, error) {
	m := _plugin.Meta{
		Class:   "hours",
		Version: version,
	}
	w.logger.Info("Meta rulez", "meta", m)
	return m, nil
}

func (w *WorkingHours) Setup(config map[string]interface{}) error {
	w.logger.Info("Setup", "config", config)
	w.config = config
	return nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:           hclog.Trace,
		Output:          os.Stderr,
		JSONFormat:      true,
		IncludeLocation: true,
	})

	w := &WorkingHours{logger: logger}

	logger.Debug("Workinghours plugin is initialized")
	var pluginMap = map[string]plugin.Plugin{
		"hours":  &_plugin.HoursPlugin{Impl: w},
		"plugin": &_plugin.PluginPlugin{Impl: w},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
	logger.Debug("Workinghours plugin is closed")
}
