package testlink

import "strconv"

// ActionOnDuplicate
type ActionOnDuplicate int

const (
	ActionOnDuplicateBlock ActionOnDuplicate = iota + 1
	ActionOnDuplicateGenerateNew
	ActionOnDuplicateCreateNewVersion
)

func (e *ActionOnDuplicate) Value() int {
	return int(*e)
}

func (e *ActionOnDuplicate) StringValue() string {
	return strconv.Itoa(e.Value())
}

// ExecutionStatus
type ExecutionStatus rune

const (
	ExecutionStatusNotRun  ExecutionStatus = 'n'
	ExecutionStatusPassed  ExecutionStatus = 'p'
	ExecutionStatusFailed  ExecutionStatus = 'f'
	ExecutionStatusBlocked ExecutionStatus = 'b'
)

func (e *ExecutionStatus) Value() rune {
	return rune(*e)
}

func (e *ExecutionStatus) StringValue() string {
	return string(e.Value())
}

// ExecutionType
type ExecutionType int

const (
	ExecutionTypeManual ExecutionType = iota + 1
	ExecutionTypeAutomated
)

func (e *ExecutionType) Value() int {
	return int(*e)
}

func (e *ExecutionType) StringValue() string {
	return strconv.Itoa(e.Value())
}

type TestCaseStatus int

const (
	TestCaseStatusDraft TestCaseStatus = iota + 1
	TestCaseStatusReadyForReview
	TestCaseStatusReviewInProgress
	TestCaseStatusRework
	TestCaseStatusObsolete
	TestCaseStatusFuture
	TestCaseStatusFinal
)

func (e *TestCaseStatus) Value() int {
	return int(*e)
}

func (e *TestCaseStatus) StringValue() string {
	return strconv.Itoa(e.Value())
}

// TestImportance
type TestImportance int

const (
	TestImportanceLow TestImportance = iota + 1
	TestImportanceMedium
	TestImportanceHigh
)

func (e *TestImportance) Value() int {
	return int(*e)
}

func (e *TestImportance) StringValue() string {
	return strconv.Itoa(e.Value())
}