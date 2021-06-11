package enums

import "strconv"

type TestCaseStatus int

const (
	DRAFT TestCaseStatus = iota + 1
	READY_FOR_REVIEW
	REVIEW_IN_PROGRESS
	REWORK
	OBSOLETE
	FUTURE
	FINAL
)

func (e TestCaseStatus) Value() int {
	return int(e)
}

func (e TestCaseStatus) StringValue() string {
	return strconv.Itoa(e.Value())
}