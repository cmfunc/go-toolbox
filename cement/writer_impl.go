package cement

import (
	"context"
	"database/sql"
)

type UpdateFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
type DeleteFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
type InsertFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
