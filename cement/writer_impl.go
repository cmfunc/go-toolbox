package cement

import (
	"context"
	"database/sql"
	"sort"
	"strings"
)

type UpdateFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
type DeleteFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)
type InsertFunc func(ctx context.Context, runner Runner, query string, args ...any) (sql.Result, error)

// InsertValues 插入指定表中多条记录
func InsertValue(ctx context.Context, runner Runner, table string, columns []InsertColumn) (sql.Result, error) {
	query := []byte("insert into ")
	query = AppendString(query, table)
	columnsCombine := CombineInsertColumn(columns)
	query = AppendString(query, string(columnsCombine))
	return runner.ExecContext(ctx, string(query))
}

func CombineInsertColumn(columns []InsertColumn) []byte {
	insertColumns := []string{}
	columnValues := []ColumnValue{}
	sort.Sort(InsertColumns(columns))
	for _, column := range columns {
		insertColumns = append(insertColumns, string(column.Column))
		columnValues = append(columnValues, column.Value)
	}
	base := []byte{}
	base = AppendString(base, "(")
	columnsPart := strings.Join(insertColumns, ",")
	base = AppendString(base, columnsPart)
	base = AppendString(base, ")")
	base = AppendString(base, " Values (")
	for i, v := range columnValues {
		v.Append(base)
		if i < len(columnValues)-1 {
			base = AppendString(base, ",")
		}
	}
	base = AppendString(base, ")")
	return base
}
