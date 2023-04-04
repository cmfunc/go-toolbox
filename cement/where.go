package cement

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Where interface {
	SetWhere
	ExecWhere
}

type SetWhere interface {
	In(column string, list []interface{})
	And(column string, ex string, value interface{})
	Or(column string, ex string, value interface{})
}

type ExecWhere interface {
	String() string
}

// ===========================================

type WhereUint struct {
	Column string
	Ex     string
	Value  interface{}
}

type WhereCond struct {
	Ands []WhereUint
	Ors  []WhereUint
}

type Logic string

func (l *Logic) Add(w *WhereCond, unit WhereUint) error {
	if *l == "and" {
		w.Ands = append(w.Ands, unit)
	} else if *l == "or" {
		w.Ors = append(w.Ors, unit)
	} else {
		return errors.Errorf("unknown logic:[%s]", l)
	}
	return nil
}

func (w *WhereCond) In(logic Logic, column string, list []interface{}) *WhereCond {
	if w.Ands != nil {
		w.Ands = []WhereUint{}
	}
	strs := []string{}
	for _, v := range list {
		strs = append(strs, fmt.Sprint(v))
	}
	logic.Add(w, WhereUint{
		Column: column,
		Ex:     "IN",
		Value:  fmt.Sprintf("(%s)", strings.Join(strs, ",")),
	})
	return w
}

func (w *WhereCond) And(column string, ex string, value interface{}) *WhereCond {
	if w.Ands != nil {
		w.Ands = []WhereUint{}
	}
	w.Ands = append(w.Ands, WhereUint{
		Column: column,
		Ex:     ex,
		Value:  fmt.Sprint(value),
	})
	return w
}

func (w *WhereCond) Or(column string, ex string, value interface{}) *WhereCond {
	if w.Ors != nil {
		w.Ors = []WhereUint{}
	}
	// 对写入的数据进行校验，column与ex、value类型进行单独封装
	w.Ors = append(w.Ors, WhereUint{
		Column: column,
		Ex:     ex,
		Value:  fmt.Sprint(value),
	})
	return w
}

func (w *WhereCond) String() string {
	if w == nil {
		return ""
	}
	as := []string{}
	for _, u := range w.Ands {
		// 上层对value进行明确转换
		as = append(as, fmt.Sprintf("%s %s %v", u.Column, u.Ex, u.Value))
	}
	os := []string{}
	for _, u := range w.Ors {
		// 上层对value进行明确转换
		os = append(os, fmt.Sprintf("%s %s %v", u.Column, u.Ex, u.Value))
	}
	ass := strings.Join(as, " and ")
	oss := strings.Join(os, " or ")
	if len(os) > 0 && len(as) > 0 {
		return strings.Join([]string{ass, oss}, " and ")
	}

	return ass + oss
}
