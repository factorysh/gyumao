package plugin

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var pluginMap = map[string]plugin.Plugin{}

func meta(path string) (Meta, error) {
	pluginMap["meta"] = &PluginPlugin{}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(path),
	})
	//defer client.Kill()
	rpcClient, err := client.Client()
	if err != nil {
		return Meta{}, err
	}
	err = rpcClient.Ping()
	if err != nil {
		return Meta{}, err
	}
	raw, err := rpcClient.Dispense("meta")
	if err != nil {
		return Meta{}, err
	}
	p := raw.(Plugin)
	return p.Meta()
}
