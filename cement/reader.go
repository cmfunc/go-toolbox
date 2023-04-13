package cement

import (
	"context"
	"database/sql"
)

// Query 查询器
type Reader interface {
	Select(ctx context.Context, runner Runner, query string, args ...any) (*sql.Rows, error)
	First(ctx context.Context, runner Runner, query string, args ...any) *sql.Row
}
