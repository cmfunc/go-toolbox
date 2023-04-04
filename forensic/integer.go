package forensic

import "strconv"

// 明确的整形变量转换
// TODO: 参考zapcore的jsonencoder

type ForensicType interface {
	String() string
}

type ForeInt int

func (f ForeInt) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func Int(i int) ForensicType {
	return ForeInt(i)
}
