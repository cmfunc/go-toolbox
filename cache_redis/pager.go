package cacheredis

import "context"

type Pager interface {
	Del(ctx context.Context, key string) (n int64, err error)
	Set(ctx context.Context, key string, rows []interface{}, offset uint64) (err error)
	Get(ctx context.Context, key string, offset, limit uint64) (bytesArray [][]byte, err error)
}
