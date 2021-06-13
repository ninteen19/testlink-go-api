package enums

import "strconv"

type ActionOnDuplicate int

const (
	BLOCK ActionOnDuplicate = iota + 1
	GENERATE_NEW
	CREATE_NEW_VERSION
)

func (e ActionOnDuplicate) Value() int {
	return int(e)
}

func (e ActionOnDuplicate) StringValue() string {
	return strconv.Itoa(e.Value())
}