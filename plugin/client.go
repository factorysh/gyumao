package plugin

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var pluginMap = map[string]plugin.Plugin{}

func init() {
	pluginMap["plugin"] = &PluginPlugin{}
	pluginMap["hours"] = &HoursPlugin{}
}

type Plugins struct {
	HoursPlugins map[string]Hours
	Logger       hclog.Logger
}

func NewPlugins() *Plugins {
	return &Plugins{
		HoursPlugins: make(map[string]Hours),
		Logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Trace,
			Output:     os.Stderr,
			JSONFormat: true,
			//IncludeLocation: true,
		}),
	}
}
func (p *Plugins) register(path string, config map[string]interface{}) error {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(path),
		Logger:          p.Logger,
	})
	//defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}
	err = rpcClient.Ping()
	if err != nil {
		return err
	}
	raw, err := rpcClient.Dispense("plugin")
	if err != nil {
		return err
	}
	plug := raw.(Plugin)
	err = plug.Setup(config)
	if err != nil {
		return err
	}
	m, err := plug.Meta()
	if err != nil {
		return err
	}
	switch m.Class {
	case "hours":
		raw, err = rpcClient.Dispense("hours")
		if err != nil {
			return err
		}
		p.HoursPlugins[filepath.Base(path)] = raw.(Hours)
	default:
	}
	return nil
}
