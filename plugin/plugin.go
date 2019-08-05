package plugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Response struct {
	Value interface{}
	Error error
}

type Meta struct {
	Version string
	Class   string
}

type Plugin interface {
	Meta() *Meta
}

type Expr interface {
	Do(interface{}) (interface{}, error)
}

type ExprRPC struct{ client *rpc.Client }

func (g *ExprRPC) Do(arg interface{}) (interface{}, error) {
	var resp Response
	err := g.client.Call("Plugin.Expr", arg, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Value, resp.Error
}

type ExprRPCServer struct{ Impl Expr }

func (s *ExprRPCServer) Do(arg interface{}, resp *Response) error {
	r, err := s.Impl.Do(arg)
	resp.Value = r
	resp.Error = err
	return nil
}

type ExprPlugin struct{ Impl Expr }

func (p *ExprPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ExprRPCServer{Impl: p.Impl}, nil
}
