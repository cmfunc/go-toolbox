package errors

// 常见错误输出
const (
	_ err = iota
	TypeErr
)

var builtinErr = []string{}

func init() {
	builtinErr[TypeErr] = "type err"
}
