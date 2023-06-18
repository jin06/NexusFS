package client

import (
	"context"
	"github.com/jin06/NexusFS/api"
	"github.com/jin06/NexusFS/store"
	"github.com/spf13/viper"
)

type Manager struct {
	Conf
	store.ClientStore
}

type Conf struct {
	ServerAddr string
}

func NewManager(conf Conf) (*Manager, error) {
	m := &Manager{
		Conf: conf,
	}
	boltPaht := viper.GetString("bolt_path")
	s, err := store.NewClientStore(store.BoltStoreVal(boltPaht))
	if err != nil {
		return nil, err
	}
	m.ClientStore = s

	return m, nil
}

func (m *Manager) Mkdir(ctx context.Context, name string, mode uint32) (ino uint64, gen uint64, err error) {
	c, err := api.NewFSC(m.ServerAddr)
	if err != nil {
		return
	}
	req := &api.MkdirReq{
		Name: name,
		Mode: mode,
	}
	resp, err := c.Mkdir(ctx, req)
	if err != nil {
		return
	}
	return resp.Ino, resp.Gen, nil
}
