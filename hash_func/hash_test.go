package hashfunc

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMD5String(t *testing.T) {
	md5map := map[string]string{}
	count := 50
	for i := 0; i < count; i++ {
		key := "1522122017554500014500000156" + strconv.FormatInt(int64(i), 10)
		md5Val, err := MD5String(key)
		if err != nil {
			t.Error(err)
			return
		}
		md5map[key] = string(md5Val)
	}

	for i := 0; i < count; i++ {
		key := "1522122017554500014500000156" + strconv.FormatInt(int64(i), 10)
		md5Val, err := MD5String(key)
		if err != nil {
			t.Error(err)
			return
		}
		oldVal, ok := md5map[key]
		if ok {
			if oldVal != string(md5Val) {
				t.Error("not equal", key, oldVal, md5Val)
				return
			}
			t.Logf("key:%s \n oldVal:%s \n md5Val:%s \b", key, oldVal, string(md5Val))
			fmt.Printf("key:%s \n oldVal:%s \n md5Val:%s \b", key, oldVal, string(md5Val))
		} else {
			t.Error(key, md5Val, "not exitst oldval")
			return
		}
	}
}
