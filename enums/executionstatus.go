package enums

type ExecutionStatus rune

const (
	NOT_RUN ExecutionStatus = 'n'
	PASSED  ExecutionStatus = 'p'
	FAILED  ExecutionStatus = 'f'
	BLOCKED ExecutionStatus = 'b'
)

func (e ExecutionStatus) Value() rune {
	return rune(e)
}

func (e ExecutionStatus) StringValue() string {
	return string(e.Value())
}
