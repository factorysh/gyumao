package plugin

import (
	"net/rpc"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Meta struct {
	Version string
	Class   string
}

type Plugin interface {
	Meta() (Meta, error)
}

type PluginRPC struct{ client *rpc.Client }

func (g *PluginRPC) Meta() (Meta, error) {
	var meta Meta
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	err := g.client.Call("Plugin.Meta", new(interface{}), &meta)
	logger.Info("PluginRPC.Meta", "meta", meta, "err", err)
	if err != nil {
		return meta, err
	}
	return meta, nil
}

type PluginRPCServer struct{ Impl Plugin }

func (s *PluginRPCServer) Meta(args interface{}, resp *Meta) error {
	var err error
	*resp, err = s.Impl.Meta()

	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	logger.Info("PluginRPCServer.Meta", "resp", resp, "error", err)
	return err
}

type PluginPlugin struct{ Impl Plugin }

func (p *PluginPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	logger.Info("PluginPlugin.Server", "broker", broker)
	return &PluginRPCServer{Impl: p.Impl}, nil
}

func (PluginPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	logger.Info("PluginPlugin.Client", "muxBroker", b, "rpc.Client", c)
	return &PluginRPC{client: c}, nil
}
