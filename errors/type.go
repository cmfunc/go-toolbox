package errors

type err int

func (e err) Error() string {
	return builtinErr[e]
}
