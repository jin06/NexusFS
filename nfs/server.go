package nfs

import (
	"fmt"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/jin06/NexusFS/manager/server"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type ServerFS struct {
	Local        fs.InodeEmbedder
	Manager      *server.Manager
	Options      Options
	LocalOptions fs.Options
}

func New(opt ...Option) (sfs *ServerFS, err error) {
	opts := NewOptions(opt...)
	sfs = &ServerFS{}
	sfs.Options = opts
	locOpts := fs.Options{}
	locOpts.AllowOther = true
	locOpts.Debug = true
	locOpts.Name = sfs.Options.FuseName

	switch opts.FSType {
	case FSTYPE_DB:
		fmt.Println("Todo")
	case FSTYYPE_LOOKBACK:
		sfs.Local, err = fs.NewLoopbackRoot(opts.LookBackPath)
		if err != nil {
			return
		}
	case FSTYPE_DEFAULT:
		fallthrough
	default:

	}
	return
}

func (sfs *ServerFS) embed() *fs.Inode {
	return nil
}

func (sfs *ServerFS) EmbeddedInode() *fs.Inode {
	return nil
}

func Mount(dir string, root fs.InodeEmbedder, options *fs.Options) (*fuse.Server, error) {
	return fs.Mount(dir, root, options)
}

func (sfs *ServerFS) Mount() (err error) {
	logrus.Debug("fstype: ", sfs.Options.FSType)

	var fsServer *fuse.Server
	fsServer, err = fs.Mount(sfs.Options.MountPoint, sfs.Local, &sfs.LocalOptions)
	if err != nil {
		logrus.Errorf("Mount lookback type filesystem error: %v", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err = fsServer.Unmount()
	}()
	fsServer.Wait()



	return
}
