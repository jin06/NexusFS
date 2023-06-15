package nfs

import (
	"github.com/hanwen/go-fuse/v2/fs"
)

type ClientFS11 struct {
	Local  fs.InodeEmbedder
}
