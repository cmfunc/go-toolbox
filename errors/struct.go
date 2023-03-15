package errors

import "fmt"

type Err struct {
	Code   uint32        `json:"code,omitempty"`
	Foramt string        `json:"foramt,omitempty"`
	Args   []interface{} `json:"args,omitempty"`
}

func NewErr(format string, code uint32, args ...interface{}) *Err {
	return &Err{
		Code:   code,
		Foramt: format,
		Args:   args,
	}
}

func (e *Err) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf(e.Foramt, e.Args...)
}

func (e *Err) SetArgs(args ...interface{}) *Err {
	if e == nil {
		return &Err{Args: args}
	}
	e.Args = args
	return e
}

func (e *Err) SetFormat(format string) *Err {
	if e == nil {
		return &Err{Foramt: format}
	}
	e.Foramt = format
	return e
}

func (e *Err) SetCode(code uint32) *Err {
	if e == nil {
		return &Err{Code: code}
	}
	e.Code = code
	return e
}
