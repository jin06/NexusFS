package nfs
//
//import (
//	"context"
//	"fmt"
//	"github.com/hanwen/go-fuse/v2/fs"
//	"github.com/hanwen/go-fuse/v2/fuse"
//	"github.com/sirupsen/logrus"
//	"os"
//	"syscall"
//)
//
//func NewServerNode(i fs.InodeEmbedder) (sn *ServerNode, err error) {
//	sn = &ServerNode{i}
//	return
//}
//
//type ServerNode struct {
//	fs.InodeEmbedder
//}
//
//func (n *ServerNode) OnAdd(ctx context.Context) {
//	logrus.Debug("OnAdd testing")
//}
//
//var _ = (fs.NodeMkdirer)((*ServerNode)(nil))
//
//func (n *ServerNode) Mkdir(ctx context.Context, name string, mode uint32, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
//	logrus.SetOutput(os.Stdout)
//	logrus.Debug("Mkdir testing")
//	fmt.Fprintln(os.Stdout, "Mkdir testing")
//	if dn, ok := n.InodeEmbedder.(fs.NodeMkdirer); ok {
//		return dn.Mkdir(ctx, name, mode, out)
//	}
//	return nil, 0
//}
//
//var _ = (fs.NodeLookuper)((*ServerNode)(nil))
//
//func (n *ServerNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
//	logrus.Debug("Lookup testing")
//	if dn, ok := n.InodeEmbedder.(fs.NodeLookuper); ok {
//		return dn.Lookup(ctx, name, out)
//	}
//	return nil, 0
//}
