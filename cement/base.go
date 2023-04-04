package cement

// sql拼接基础结构
type Base struct {
	Select    []string
	TableName string
}



func Update(table string) *Base {
	base := &Base{TableName: table}
	return base
}

func Delete(table string) *Base {
	base := &Base{TableName: table}
	return base
}

func Insert(table string) *Base {
	base := &Base{TableName: table}
	return base
}