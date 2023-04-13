package cement

import (
	"context"
	"database/sql"
)

// Writer 修改器
type Writer interface {
	Update(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
	Delete(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
	Insert(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
}
