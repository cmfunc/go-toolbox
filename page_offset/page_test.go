package pageoffset

import (
	"context"
	"testing"
)

func TestTransOffset(t *testing.T) {
	ctx := context.Background()
	cases := []struct{ index, size uint64 }{{0, 1}, {0, 10}, {1, 10}, {10, 0}, {0, 0}, {10, 20}}
	results := []struct{ offset, limit uint64 }{{0, 1}, {0, 10}, {0, 10}, {90, 10}, {0, 10}, {180, 20}}
	for i, c := range cases {
		res := results[i]
		offset, limit := TransOffset(ctx, c.index, c.size)
		if offset == res.offset && limit == res.limit {
			continue
		}
		t.Errorf("cases{index:%d,size:%d} TransOffset return offset:%d,limit:%d ,in fact results{offset:%d,limit:%d}",
			c.index, c.size, offset, limit, res.offset, res.limit,
		)
	}
	t.Log("Success")
}
