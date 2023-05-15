package nfs

import (
	"github.com/hanwen/go-fuse/v2/fs"
)

type Client struct {
	Local  fs.InodeEmbedder
}
