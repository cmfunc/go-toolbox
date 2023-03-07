package tables

import (
	"fmt"
	"hash/crc32"
)

// DivideTableName: sum table name that have divide by key
// param origin is origin table name
// param num is num of being divide numbers
// param key is  table's column for dividing table
func DivideTableName(origin string, num uint32, key string) (tableName string, err error) {
	md5Key, err := MD5String(key)
	if err != nil {
		return "", err
	}
	No := crc32.ChecksumIEEE(md5Key) % num
	return fmt.Sprintf(origin+"_%d", No), nil
}
