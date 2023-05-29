package nfs

import (
	"github.com/hanwen/go-fuse/v2/fs"
)

type ClientFS struct {
	Local  fs.InodeEmbedder
}
