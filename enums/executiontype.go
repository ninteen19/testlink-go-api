package enums

import "strconv"

type ExecutionType int

const (
	MANUAL ExecutionType = iota + 1
	AUTOMATED
)

func (e ExecutionType) Value() int {
	return int(e)
}

func (e ExecutionType) StringValue() string {
	return strconv.Itoa(e.Value())
}