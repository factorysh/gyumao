package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Meta struct {
	Version string
	Class   string
}

type SetupError struct {
	Error *plugin.BasicError
}

type Plugin interface {
	Meta() (Meta, error)
	Setup(map[string]interface{}) error
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

func (g *PluginRPC) Setup(config map[string]interface{}) error {
	var resp SetupError
	err := g.client.Call("Plugin.Setup", &config, &resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

type PluginRPCServer struct{ Impl Plugin }

func (s *PluginRPCServer) Meta(args interface{}, resp *Meta) error {
	var err error
	*resp, err = s.Impl.Meta()

	return err
}

func (s *PluginRPCServer) Setup(args map[string]interface{}, resp *SetupError) error {
	err := s.Impl.Setup(args)
	if err != nil {
		resp.Error.Message = err.Error()
	}
	return err
}

type PluginPlugin struct{ Impl Plugin }

func (p *PluginPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &PluginRPCServer{Impl: p.Impl}, nil
}

func (PluginPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPC{client: c}, nil
}
