package nfs

import (
	"github.com/hanwen/go-fuse/v2/fs"
	"syscall"
)

func loopbackNewNode(rootData *fs.LoopbackRoot, _ *fs.Inode, _ string, _ *syscall.Stat_t) fs.InodeEmbedder {
	n := &ServerNode{
		LoopbackNode: fs.LoopbackNode{
			RootData: rootData,
		},
	}
	return n
}
