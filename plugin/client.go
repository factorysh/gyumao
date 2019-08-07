package plugin

import (
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var pluginMap = map[string]plugin.Plugin{}

func getPlugin(path string) (Plugin, error) {
	pluginMap["meta"] = &PluginPlugin{}

	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
		//IncludeLocation: true,
	})
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(path),
		Logger:          logger,
	})
	//defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	err = rpcClient.Ping()
	if err != nil {
		return nil, err
	}
	raw, err := rpcClient.Dispense("meta")
	if err != nil {
		return nil, err
	}
	p := raw.(Plugin)
	return p, nil
}
