package cement

import (
	"fmt"
	"strings"
)

type SelectBase struct {
	columns   []string
	tableName string
	condition []string
	orderby   []string
	offset    uint64
	limit     uint64
	query     string
}

func Select(columns ...string) *SelectBase {
	base := &SelectBase{columns: []string{}}
	base.columns = append(base.columns, columns...)
	return base
}

func (b *SelectBase) From(table string) *SelectBase {
	b.tableName = table
	return b
}

func (b *SelectBase) Where(where string) *SelectBase {
	b.condition = append(b.condition, where)
	return b
}

func (b *SelectBase) Order(column string, asc bool) *SelectBase {
	sort := "DESC"
	if asc {
		sort = "ASC"
	}
	b.orderby = append(b.orderby, column+" "+sort)
	return b
}

func (b *SelectBase) Offset(offset uint64) *SelectBase {
	b.offset = offset
	return b
}

func (b *SelectBase) Limit(limit uint64) *SelectBase {
	b.limit = limit
	return b
}

func (b *SelectBase) Build() string {
	s := strings.Join(b.columns, ",")
	w := strings.Join(b.condition, " and ")
	o := strings.Join(b.orderby, ",")
	b.query = fmt.Sprintf("select %s from %s where %s order by %s offfset %d limit %d", s, b.tableName, w, o, b.offset, b.limit)
	return b.query
}
