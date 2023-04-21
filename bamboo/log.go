package bamboo

// 日志接口
// 用于打印记录消息传递过程中的重要信息
type Logger interface {
	Error(err error, format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
}


