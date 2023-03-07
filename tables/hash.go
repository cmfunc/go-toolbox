package tables

import (
	"crypto/md5"
)

func MD5String(key string) ([]byte, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(key))
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
