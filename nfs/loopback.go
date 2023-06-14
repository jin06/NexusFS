package nfs

import (
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/jin06/NexusFS/manager/client"
	"syscall"
)

//func loopbackNewNode(rootData *fs.LoopbackRoot, _ *fs.Inode, _ string, _ *syscall.Stat_t) fs.InodeEmbedder {
//	n := &ServerNode{
//		LoopbackNode: fs.LoopbackNode{
//			RootData: rootData,
//		},
//	}
//	return n
//}

func loopbackNewNodeWrapper(manager *client.Manager) func(*fs.LoopbackRoot, *fs.Inode, string, *syscall.Stat_t) fs.InodeEmbedder {
	return func(rootData *fs.LoopbackRoot, _ *fs.Inode, _ string, _ *syscall.Stat_t) fs.InodeEmbedder {
		n := &ServerNode{
			LoopbackNode: fs.LoopbackNode{
				RootData: rootData,
			},
			M: manager,
		}
		return n
	}
}
