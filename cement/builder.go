package cement

import (
	"strconv"
	"time"
)

// sql 字符串拼接
func AppendByte(bs []byte, v byte) []byte {
	return append(bs, v)
}

func AppendString(bs []byte, s string) []byte {
	return append(bs, s...)
}

func AppendInt(bs []byte, i int64) []byte {
	return strconv.AppendInt(bs, i, 10)
}

func AppendTime(bs []byte, t time.Time, layout string) []byte {
	return t.AppendFormat(bs, layout)
}

func AppendUint(bs []byte, i uint64) []byte {
	return strconv.AppendUint(bs, i, 10)
}

func AppendBool(bs []byte, v bool) []byte {
	return strconv.AppendBool(bs, v)
}

// AppendFloat appends a float to the underlying buffer. It doesn't quote NaN
// or +/- Inf.
func AppendFloat(bs []byte, f float64, bitSize int) []byte {
	return strconv.AppendFloat(bs, f, 'f', -1, bitSize)
}

func Len(bs []byte) int {
	return len(bs)
}

func Cap(bs []byte) int {
	return cap(bs)
}

func String(bs []byte) string {
	return string(bs)
}

func Reset(bs []byte) {
	bs = (bs)[:0]
}

