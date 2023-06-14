package handlers

import (
	"context"
	"fmt"
	"github.com/jin06/NexusFS/api"
)

type FS struct {
}

func NewFS() *FS {
	return &FS{}
}

func (s *FS) Mkdir(ctx context.Context, req *api.MkdirReq) (*api.MkdirResp, error) {
	fmt.Println(req.Name, req.Mode)
	resp := &api.MkdirResp{
		Mode: 1,
		Ino:  1,
		Gen:  1,
	}
	return resp, nil
}
