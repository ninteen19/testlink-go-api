package testcase

import (
	"errors"
	"github.com/ninteen19/testlink-go-api"
)

type CreateRequest struct {
	TestCaseName           string
	TestSuiteId            int
	TestProjectId          int
	AuthorLogin            string
	Summary                string
	Steps                  []testlink.TestCaseStep
	Preconditions          string
	Status                 testlink.TestCaseStatus
	Importance             testlink.TestImportance
	Execution              testlink.ExecutionType
	Order                  int
	InternalId             int
	CheckDuplicatedName    bool
	ActionOnDuplicatedName testlink.ActionOnDuplicate
}

func (r *CreateRequest) Validate() error {
	if isEmpty(r.AuthorLogin) {
		return errors.New("field AuthorLogin is empty")
	}

	if isEmpty(r.TestCaseName) {
		return errors.New("field TestCaseName is empty")
	}

	if r.TestSuiteId == 0 || r.TestProjectId == 0 {
		return errors.New("field TestSuiteId or TestProjectId is 0")
	}

	if len(r.Steps) == 0 {
		return errors.New("field Steps is empty")
	}

	return nil
}

func isEmpty(s string) bool {
	return len(s) < 0
}
