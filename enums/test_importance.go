package enums

import "strconv"

type TestImportance int

const (
	LOW TestImportance = iota + 1
	MEDIUM
	HIGH
)

func (e TestImportance) Value() int {
	return int(e)
}

func (e TestImportance) StringValue() string {
	return strconv.Itoa(e.Value())
}
