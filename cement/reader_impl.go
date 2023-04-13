package cement

import (
	"context"
	"database/sql"
)

// 查询器的实例
type SelectFunc func(ctx context.Context, runner Runner, query string, args ...any) (*sql.Rows, error)
type FirstFunc func(ctx context.Context, runner Runner, query string, args ...any) *sql.Row
