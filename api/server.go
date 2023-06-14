package api

import (
	"context"
	"fmt"
)

type Server struct {
}

func NewServer() *Server {
	s := &Server{}
	return s
}

func (s *Server) SayHello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	rs := fmt.Sprintf("Hello, %s", req.Name)
	resp := &HelloResp{
		Reply: rs,
	}
	return resp, nil
}

