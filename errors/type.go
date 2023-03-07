package errors

type err int

func (e err) String() string {
	return builtinErr[e]
}
