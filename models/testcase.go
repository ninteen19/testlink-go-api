package models

import (
	"github.com/ninteen19/testlink-go-api/constants"
	"github.com/ninteen19/testlink-go-api/enums"
)

type TestCase struct {
	Id                     int                     `json:"id"`
	Name                   string                  `json:"name"`
	TestSuiteId            int                     `json:"testSuiteId"`
	TestProjectId          int                     `json:"testProjectId"`
	AuthorLogin            string                  `json:"authorLogin"`
	Summary                string                  `json:"summary"`
	Steps                  []TestCaseStep          `json:"steps"`
	Preconditions          string                  `json:"preconditions"`
	TestCaseStatus         enums.TestCaseStatus    `json:"testCaseStatus"`
	TestImportance         enums.TestImportance    `json:"testImportance"`
	ExecutionType          enums.ExecutionType     `json:"executionType"`
	ExecutionOrder         int                     `json:"executionOrder"`
	Order                  int                     `json:"order"`
	InternalId             int                     `json:"internalId"`
	ExternalId             int                     `json:"externalId"`
	FullExternalId         string                  `json:"fullExternalId"`
	CheckDuplicatedName    bool                    `json:"checkDuplicatedName"`
	ActionOnDuplicatedName enums.ActionOnDuplicate `json:"actionOnDuplicatedName"`
	VersionId              int                     `json:"versionId"`
	Version                int                     `json:"version"`
	ParentId               int                     `json:"parentId"`
	CustomFields           []CustomField           `json:"customFields"`
	ExecutionStatus        enums.ExecutionStatus   `json:"executionStatus"`
	Platform               Platform                `json:"platform"`
	FeatureId              int                     `json:"featureId"`
	Keywords               []string                `json:"keywords"`
}

type TestCaseStep struct {
	Id                int                 `json:"id"`
	TestCaseVersionId int                 `json:"testCaseVersionId"`
	Number            int                 `json:"number"`
	Actions           string              `json:"actions"`
	ExpectedResults   string              `json:"expectedResults"`
	Active            bool                `json:"active"`
	ExecutionType     enums.ExecutionType `json:"executionType"`
}

func (t *TestCase) ToMap() map[string]interface{} {
	testCaseMap := map[string]interface{}{}

	putIfNotNull(testCaseMap, constants.TestCaseName, t.Name)
	putIfNotNull(testCaseMap, constants.TestCaseId, t.Id)
	putIfNotNull(testCaseMap, constants.TestSuiteId, t.TestSuiteId)
	putIfNotNull(testCaseMap, constants.TestProjectId, t.TestProjectId)
	putIfNotNull(testCaseMap, constants.AuthorLogin, t.AuthorLogin)
	putIfNotNull(testCaseMap, constants.Summary, t.Summary)

	if len(t.Steps) > 0 {
		putIfNotNull(testCaseMap, constants.Steps, ToTestCaseStepMaps(&t.Steps))
	}

	putIfNotNull(testCaseMap, constants.Preconditions, t.Preconditions)
	putIfNotNull(testCaseMap, constants.Status, t.TestCaseStatus)
	putIfNotNull(testCaseMap, constants.Importance, t.TestImportance)
	putIfNotNull(testCaseMap, constants.ExecutionType, t.ExecutionType)
	putIfNotNull(testCaseMap, constants.Order, t.Order)
	putIfNotNull(testCaseMap, constants.InternalId, t.InternalId)
	putIfNotNull(testCaseMap, constants.CheckDuplicatedName, t.CheckDuplicatedName)
	putIfNotNull(testCaseMap, constants.ActionOnDuplicatedName, t.ActionOnDuplicatedName)

	return testCaseMap
}

func ToTestCaseStepMaps(testCaseSteps *[]TestCaseStep) []map[string]interface{} {
	var testCaseStepMaps []map[string]interface{}

	for _, testCaseStep := range *testCaseSteps {
		testCaseStepMap := map[string]interface{}{}
		putIfNotNull(testCaseStepMap, constants.StepNumber, testCaseStep.Number)
		putIfNotNull(testCaseStepMap, constants.Actions, testCaseStep.Actions)
		putIfNotNull(testCaseStepMap, constants.ExpectedResults, testCaseStep.ExpectedResults)
		putIfNotNull(testCaseStepMap, constants.StepExecutionType, testCaseStep.ExecutionType.Value())

		testCaseStepMaps = append(testCaseStepMaps, testCaseStepMap)
	}

	return testCaseStepMaps
}

func putIfNotNull(testCaseMap map[string]interface{}, key string, val interface{}) {
	if val != nil {
		testCaseMap[key] = val
	}
}
