package cement

type IntValue int64

func (t *IntValue) Append(bs []byte) []byte {
	return AppendInt(bs, int64(*t))
}

type UintValue uint64

func (t *UintValue) Append(bs []byte) []byte {
	return AppendUint(bs, uint64(*t))
}
