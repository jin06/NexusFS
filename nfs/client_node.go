package nfs

import (
	"context"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/jin06/NexusFS/manager/client"
	"github.com/sirupsen/logrus"
	"syscall"
)

type ClientNode struct {
	fs.LoopbackNode
	M *client.Manager
}

func (n *ClientNode) OnAdd(ctx context.Context) {
	logrus.Debug("OnAdd testing")
}

//var _ = (fs.NodeStatfser)((*ServerNode)(nil))
//
//func (n *ServerNode) Statfs(ctx context.Context, out *fuse.StatfsOut) syscall.Errno {
//	return fs.OK
//}
//
var _ = (fs.NodeLookuper)((*ClientNode)(nil))

func (n *ClientNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	logrus.Debug("Lookup testing: ", name)
	return n.LoopbackNode.Lookup(ctx, name, out)
}

var _ = (fs.NodeMknoder)((*ClientNode)(nil))

func (n *ClientNode) Mknod(ctx context.Context, name string, mode, rdev uint32, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	logrus.Debug("Mknod testing: ", name)
	return n.LoopbackNode.Mknod(ctx, name, mode, rdev, out)
}

var _ = (fs.NodeMkdirer)((*ClientNode)(nil))

func (n *ClientNode) Mkdir(ctx context.Context, name string, mode uint32, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	logrus.Debug("Mkdir testing: ", name)

	i, g, err := n.M.Mkdir(ctx, name, mode)
	logrus.Debug("mkdir response test resp :", i, g, err)

	return n.LoopbackNode.Mkdir(ctx, name, mode, out)
}

//var _ = (fs.NodeRmdirer)((*ServerNode)(nil))
//
//func (n *ServerNode) Rmdir(ctx context.Context, name string) syscall.Errno {
//	return 0
//}
//
//var _ = (fs.NodeUnlinker)((*ServerNode)(nil))
//
//func (n *ServerNode) Unlink(ctx context.Context, name string) syscall.Errno {
//	return 0
//}

var _ = (fs.NodeRenamer)((*ClientNode)(nil))

func (n *ClientNode) Rename(ctx context.Context, name string, newParent fs.InodeEmbedder, newName string, flags uint32) syscall.Errno {
	logrus.Debugf("Rename testing: %s ---> %s", name, newName)
	return n.LoopbackNode.Rename(ctx, name, newParent, newName, flags)
}

//
var _ = (fs.NodeCreater)((*ClientNode)(nil))

func (n *ClientNode) Create(ctx context.Context, name string, flags uint32, mode uint32, out *fuse.EntryOut) (inode *fs.Inode, fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	return n.LoopbackNode.Create(ctx, name, flags, mode, out)
}

//
//var _ = (fs.NodeSymlinker)((*ServerNode)(nil))
//
//func (n *ServerNode) Symlink(ctx context.Context, target, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
//	return nil, 0
//}
//
//var _ = (fs.NodeLinker)((*ServerNode)(nil))
//
//func (n *ServerNode) Link(ctx context.Context, target fs.InodeEmbedder, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
//	return nil, 0
//}
//
//var _ = (fs.NodeReadlinker)((*ServerNode)(nil))
//
//func (n *ServerNode) Readlink(ctx context.Context) ([]byte, syscall.Errno) {
//	return nil, 0
//}
//
//var _ = (fs.NodeOpener)((*ServerNode)(nil))
//
//func (n *ServerNode) Open(ctx context.Context, flags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
//	return nil, 0, 0
//}
//
//var _ = (fs.NodeOpendirer)((*ServerNode)(nil))
//
//func (n *ServerNode) Opendir(ctx context.Context) syscall.Errno {
//	return 0
//}
//
//var _ = (fs.NodeReaddirer)((*ServerNode)(nil))
//
//func (n *ServerNode) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
//	logrus.Debug("Readdir testing")
//	return nil, 0
//}
//
//var _ = (fs.NodeGetattrer)((*ServerNode)(nil))
//
//func (n *ServerNode) Getattr(ctx context.Context, f fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
//	logrus.Debug("Getattr testing")
//	fmt.Fprintln(os.Stdout,"kjdslkfjl")
//	return 0
//}
//
//var _ = (fs.NodeSetattrer)((*ServerNode)(nil))
//
//func (n *ServerNode) Setattr(ctx context.Context, f fs.FileHandle, in *fuse.SetAttrIn, out *fuse.AttrOut) syscall.Errno {
//	return 0
//}
