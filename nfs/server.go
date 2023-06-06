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
	"time"
)

type ServerFS struct {
	Local        fs.InodeEmbedder
	Manager      *server.Manager
	Options      Options
	LocalOptions *fs.Options
}

func New(opt ...Option) (sfs *ServerFS, err error) {
	opts := NewOptions(opt...)
	sfs = &ServerFS{}
	sfs.Options = opts
	sec := time.Second
	locOpts := fs.Options{
		// The timeout options are to be compatible with libfuse defaults,
		// making benchmarking easier.
		AttrTimeout:  &sec,
		EntryTimeout: &sec,

		NullPermissions: true,

		MountOptions: fuse.MountOptions{
			AllowOther: true,
			//Debug:             true,
			FsName: sfs.Options.LookBackPath, // First column in "df -T": original dir
			Name:   sfs.Options.FuseName,     // Second column in "df -T" will be shown as "fuse." + Name
		},
	}
	if locOpts.AllowOther {
		// Make the kernel check file permissions for us
		locOpts.MountOptions.Options = append(locOpts.MountOptions.Options, "default_permissions")
	}

	sfs.LocalOptions = &locOpts

	switch opts.FSType {
	case FSTYPE_DB:
		fmt.Println("Todo")
	case FSTYYPE_LOOPBACK:
		//sfs.Local, err = fs.NewLoopbackRoot(opts.LookBackPath)
		//var im fs.InodeEmbedder
		//im , err = fs.NewLoopbackRoot(opts.LookBackPath)
		//if err != nil {
		//	logrus.Error("New loopback root error: ", err)
		//	return
		//}
		//sfs.Local, err  = NewServerNode(im)

		//if err != nil {
		//	return
		//}

		rootData := &fs.LoopbackRoot{
			Path:    opts.LookBackPath,
			NewNode: loopbackNewNode,
		}
		sfs.Local = loopbackNewNode(rootData, nil, "", nil)
	case FSTYPE_DEBUG:
		//sfs.Local = ServerNode{fs.Inode{}}
	case FSTYPE_DEFAULT:
		fallthrough
	default:

	}
	return
}

func (sfs *ServerFS) Mount() (err error) {
	logrus.Debug("fstype: ", sfs.Options.FSType)

	var fsServer *fuse.Server
	fsServer, err = fs.Mount(sfs.Options.MountPoint, sfs.Local, sfs.LocalOptions)
	if err != nil {
		logrus.Errorf("Mount lookback type filesystem error: %v", err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logrus.Debug("Unmounting...")
		err = fsServer.Unmount()
	}()
	fsServer.Wait()
	return
}
