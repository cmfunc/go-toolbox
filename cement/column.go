package cement

type Column string

type OperateSymble string

type ColumnValue interface {
	Append(bs []byte) []byte
}

type SetColumn struct {
	Column Column
	Value  ColumnValue
	Symble OperateSymble
	Sort   int
}

type InsertColumn struct {
	Column Column
	Value  ColumnValue
	Sort   int
}

type InsertColumns []InsertColumn

func (a InsertColumns) Len() int           { return len(a) }
func (a InsertColumns) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InsertColumns) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
