package plugin

import (
	"net/rpc"

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
	err := g.client.Call("Plugin.Meta", new(interface{}), &meta)
	if err != nil {
		return meta, err
	}
	return meta, nil
}

type PluginRPCServer struct{ Impl Plugin }

func (s *PluginRPCServer) Meta(args interface{}, resp *Meta) error {
	var err error
	*resp, err = s.Impl.Meta()

	return err
}

type PluginPlugin struct{ Impl Plugin }

func (p *PluginPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &PluginRPCServer{Impl: p.Impl}, nil
}

func (PluginPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPC{client: c}, nil
}
