package tables

import (
	"strconv"
	"testing"
)

func TestDivideTableName(t *testing.T) {
	resMap := map[int]string{}
	origin := "my_table"
	num := 10
	for i := 0; i < 700; i++ {
		tbname, err := DivideTableName(origin, uint32(num), strconv.FormatInt(int64(i), 10))
		if err != nil {
			t.Error(err)
			return
		}
		resMap[i] = tbname
	}

	for i := 0; i < 700; i++ {
		tbname, err := DivideTableName(origin, uint32(num), strconv.FormatInt(int64(i), 10))
		if err != nil {
			t.Error(err)
			return
		}
		tbnameOld := resMap[i]
		if tbname != tbnameOld {
			t.Errorf("i:%d\n tbname:%s\n tbnameOld:%s\n", i, tbname, tbnameOld)
			return
		}
		t.Logf("i:%d\n tbname:%s\n tbnameOld:%s\n", i, tbname, tbnameOld)
	}
}
