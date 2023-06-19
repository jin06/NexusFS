package handlers

import (
	"context"
	"fmt"
	"github.com/jin06/NexusFS/api"
	"github.com/spf13/afero"
	"os"
	"syscall"
)

type FS struct {
	LocalFS afero.Fs
}

func NewFS(path string) *FS {
	fs := &FS{
	}
	fs.LocalFS = afero.NewBasePathFs(afero.NewOsFs(), path)

	return fs
}

func (s *FS) Mkdir(ctx context.Context, req *api.MkdirReq) (*api.MkdirResp, error) {
	fmt.Println(req.Name, req.Mode)
	resp := &api.MkdirResp{}
	err := s.LocalFS.Mkdir(req.Name, os.FileMode(req.Mode))
	if err != nil {
		return nil, err
	}
	fi, err := s.LocalFS.Stat(req.Name)
	if err != nil {
		return nil, err
	}
	if st, ok := fi.Sys().(syscall.Stat_t); ok {
		resp.Ino = st.Ino
		resp.Mode = uint32(fi.Mode())
		resp.Gen = uint64(st.Gen)
		_ = st.Atimespec
		_ = st.Mtimespec
		_ = st.Ctimespec
	}
	return resp, nil
}
