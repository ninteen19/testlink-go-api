package testlink

type TestCase struct {
	Id                     int               `json:"id"`
	Name                   string            `json:"name"`
	TestSuiteId            int               `json:"testSuiteId"`
	TestProjectId          int               `json:"testProjectId"`
	AuthorLogin            string            `json:"authorLogin"`
	Summary                string            `json:"summary"`
	Steps                  []TestCaseStep    `json:"steps"`
	Preconditions          string            `json:"preconditions"`
	TestCaseStatus         TestCaseStatus    `json:"testCaseStatus"`
	TestImportance         TestImportance    `json:"testImportance"`
	ExecutionType          ExecutionType     `json:"executionType"`
	ExecutionOrder         int               `json:"executionOrder"`
	Order                  int               `json:"order"`
	InternalId             int               `json:"internalId"`
	ExternalId             int               `json:"externalId"`
	FullExternalId         string            `json:"fullExternalId"`
	CheckDuplicatedName    bool              `json:"checkDuplicatedName"`
	ActionOnDuplicatedName ActionOnDuplicate `json:"actionOnDuplicatedName"`
	VersionId              int               `json:"versionId"`
	Version                int               `json:"version"`
	ParentId               int               `json:"parentId"`
	CustomFields           []CustomField     `json:"customFields"`
	ExecutionStatus        ExecutionStatus   `json:"executionStatus"`
	Platform               Platform          `json:"platform"`
	FeatureId              int               `json:"featureId"`
	Keywords               []string          `json:"keywords"`
}

type TestCaseStep struct {
	Id                int           `json:"id"`
	TestCaseVersionId int           `json:"testCaseVersionId"`
	Number            int           `json:"number"`
	Actions           string        `json:"actions"`
	ExpectedResults   string        `json:"expectedResults"`
	Active            bool          `json:"active"`
	ExecutionType     ExecutionType `json:"executionType"`
}

func (t *TestCase) ToMap() map[string]interface{} {
	testCaseMap := map[string]interface{}{}

	putIfNotNull(testCaseMap, TestLinkParamTestCaseName, t.Name)
	putIfNotNull(testCaseMap, TestLinkParamTestCaseId, t.Id)
	putIfNotNull(testCaseMap, TestLinkParamTestSuiteId, t.TestSuiteId)
	putIfNotNull(testCaseMap, TestLinkParamTestProjectId, t.TestProjectId)
	putIfNotNull(testCaseMap, TestLinkParamAuthorLogin, t.AuthorLogin)
	putIfNotNull(testCaseMap, TestLinkParamSummary, t.Summary)

	if len(t.Steps) > 0 {
		putIfNotNull(testCaseMap, TestLinkParamSteps, ToTestCaseStepMaps(&t.Steps))
	}

	putIfNotNull(testCaseMap, TestLinkParamPreconditions, t.Preconditions)
	putIfNotNull(testCaseMap, TestLinkParamStatus, t.TestCaseStatus)
	putIfNotNull(testCaseMap, TestLinkParamImportance, t.TestImportance)
	putIfNotNull(testCaseMap, TestLinkParamExecutionType, t.ExecutionType)
	putIfNotNull(testCaseMap, TestLinkParamOrder, t.Order)
	putIfNotNull(testCaseMap, TestLinkParamInternalId, t.InternalId)
	putIfNotNull(testCaseMap, TestLinkParamCheckDuplicatedName, t.CheckDuplicatedName)
	putIfNotNull(testCaseMap, TestLinkParamActionOnDuplicatedName, t.ActionOnDuplicatedName)

	return testCaseMap
}

func ToTestCaseStepMaps(testCaseSteps *[]TestCaseStep) []map[string]interface{} {
	var testCaseStepMaps []map[string]interface{}

	for _, testCaseStep := range *testCaseSteps {
		testCaseStepMap := map[string]interface{}{}
		putIfNotNull(testCaseStepMap, TestLinkParamStepNumber, testCaseStep.Number)
		putIfNotNull(testCaseStepMap, TestLinkParamActions, testCaseStep.Actions)
		putIfNotNull(testCaseStepMap, TestLinkParamExpectedResults, testCaseStep.ExpectedResults)
		putIfNotNull(testCaseStepMap, TestLinkParamStepExecutionType, testCaseStep.ExecutionType.Value())

		testCaseStepMaps = append(testCaseStepMaps, testCaseStepMap)
	}

	return testCaseStepMaps
}

func putIfNotNull(testCaseMap map[string]interface{}, key string, val interface{}) {
	if val != nil {
		testCaseMap[key] = val
	}
}
