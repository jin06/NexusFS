package nfs

import (
	"context"
	"syscall"
)

type Storage interface{
	Statfs(ctx context.Context) (syscall.Statfs_t, error)
}
