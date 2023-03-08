package errors

// 常见错误输出
const (
	_ err = iota
	TypeErr
)

var builtinErr = []string{}

func init() {
	initHttp()
	
	builtinErr = append(builtinErr, httpErr...)
	builtinErr[TypeErr] = "type err"
}
