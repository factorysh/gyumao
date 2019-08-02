package plugin

import (
	"net/rpc"
	"time"

	"github.com/hashicorp/go-plugin"
)

type Tags map[string]string

type Hours interface {
	Time(time.Time) Tags
}

type HoursRPC struct{ client *rpc.Client }

func (g *HoursRPC) Time(t time.Time) (Tags, error) {
	var resp Tags
	err := g.client.Call("Plugin.Hours", t, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *HoursRPC) Now() (Tags, error) {
	return g.Time(time.Now())
}

type HoursRPCServer struct{ Impl Hours }

func (s *HoursRPCServer) Time(args time.Time, resp *Tags) error {
	*resp = s.Impl.Time(args)
	return nil
}

type HoursPlugin struct{ Impl Hours }

func (p *HoursPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &HoursRPCServer{Impl: p.Impl}, nil
}

func (HoursPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &HoursRPC{client: c}, nil
}
