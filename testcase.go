package testlink

type TestCaseRequest struct {
	Id             int            `json:"id" xml:"testcaseid"`
	Name           string         `json:"name" xml:"testcasename"`
	TestSuiteId    int            `json:"testSuiteId" xml:"testsuiteid"`
	TestProjectId  int            `json:"testProjectId" xml:"testprojectid"`
	AuthorLogin    string         `json:"authorLogin" xml:"authorlogin"`
	Summary        string         `json:"summary" xml:"summary"`
	Steps          []TestCaseStep `json:"steps" xml:"steps"`
	Preconditions  string         `json:"preconditions" xml:"preconditions"`
	TestCaseStatus int            `json:"testCaseStatus" xml:"status"`
	TestImportance int            `json:"testImportance" xml:"importance"`
	ExecutionType  int            `json:"executionType" xml:"executiontype"`
	//ExecutionOrder         int               `json:"executionOrder"`
	Order      int `json:"order" xml:"order"`
	InternalId int `json:"internalId" xml:"internalid"`
	ExternalId int `json:"externalId" xml:"externalid"`
	//FullExternalId         string            `json:"fullExternalId"`
	CheckDuplicatedName    bool `json:"checkDuplicatedName" xml:"checkduplicatedname"`
	ActionOnDuplicatedName int  `json:"actionOnDuplicatedName" xml:"actiononduplicatedname"`
	//VersionId              int               `json:"versionId" `
	Version int `json:"version"`
	//ParentId               int               `json:"parentId"`
	//CustomFields           []CustomField     `json:"customFields"`
	//ExecutionStatus        ExecutionStatus   `json:"executionStatus"`
	//Platform               Platform          `json:"platform"`
	//FeatureId              int               `json:"featureId"`
	//Keywords               []string          `json:"keywords"`
	DevKey string `json:"devKey" xml:"devKey"`
}

type TestCaseResponse struct {
	Id                     int            `json:"id" xml:"testcaseid"`
	Name                   string         `json:"name" xml:"testcasename"`
	TestSuiteId            int            `json:"testSuiteId" xml:"testsuiteid"`
	TestProjectId          int            `json:"testProjectId" xml:"testprojectid"`
	AuthorLogin            string         `json:"authorLogin" xml:"authorlogin"`
	Summary                string         `json:"summary" xml:"summary"`
	Steps                  []TestCaseStep `json:"steps" xml:"steps"`
	Preconditions          string         `json:"preconditions" xml:"preconditions"`
	TestCaseStatus         int            `json:"testCaseStatus" xml:"status"`
	TestImportance         int            `json:"testImportance" xml:"importance"`
	ExecutionType          int            `json:"executionType" xml:"executiontype"`
	ExecutionOrder         int            `json:"executionOrder"`
	Order                  int            `json:"order" xml:"order"`
	InternalId             int            `json:"internalId" xml:"internalid"`
	ExternalId             int            `json:"externalId" xml:"externalid"`
	FullExternalId         string         `json:"fullExternalId"`
	CheckDuplicatedName    bool           `json:"checkDuplicatedName" xml:"checkduplicatedname"`
	ActionOnDuplicatedName int            `json:"actionOnDuplicatedName" xml:"actiononduplicatedname"`
	VersionId              int            `json:"versionId" `
	Version                int            `json:"version"`
	ParentId               int            `json:"parentId"`
	CustomFields           []CustomField  `json:"customFields"`
	ExecutionStatus        string         `json:"executionStatus"`
	Platform               Platform       `json:"platform"`
	FeatureId              int            `json:"featureId"`
	Keywords               []string       `json:"keywords"`
}

type TestCaseStep struct {
	Id                int    `json:"id"`
	TestCaseVersionId int    `json:"testCaseVersionId"`
	Number            int    `json:"number" xml:"step_number"`
	Actions           string `json:"actions" xml:"actions"`
	ExpectedResults   string `json:"expectedResults" xml:"expected_results"`
	Active            bool   `json:"active"`
	ExecutionType     int    `json:"executionType" xml:"execution_type"`
}

//
//func (t *TestCase) ToMap() map[string]interface{} {
//	testCaseMap := map[string]interface{}{}
//
//	putIfNotNull(testCaseMap, TestLinkParamTestCaseName, t.Name)
//	putIfNotNull(testCaseMap, TestLinkParamTestCaseId, t.Id)
//	putIfNotNull(testCaseMap, TestLinkParamTestSuiteId, t.TestSuiteId)
//	putIfNotNull(testCaseMap, TestLinkParamTestProjectId, t.TestProjectId)
//	putIfNotNull(testCaseMap, TestLinkParamAuthorLogin, t.AuthorLogin)
//	putIfNotNull(testCaseMap, TestLinkParamSummary, t.Summary)
//
//	if len(t.Steps) > 0 {
//		putIfNotNull(testCaseMap, TestLinkParamSteps, ToTestCaseStepMaps(&t.Steps))
//	}
//
//	putIfNotNull(testCaseMap, TestLinkParamPreconditions, t.Preconditions)
//	putIfNotNull(testCaseMap, TestLinkParamStatus, t.TestCaseStatus)
//	putIfNotNull(testCaseMap, TestLinkParamImportance, t.TestImportance)
//	putIfNotNull(testCaseMap, TestLinkParamExecutionType, t.ExecutionType)
//	putIfNotNull(testCaseMap, TestLinkParamOrder, t.Order)
//	putIfNotNull(testCaseMap, TestLinkParamInternalId, t.InternalId)
//	putIfNotNull(testCaseMap, TestLinkParamCheckDuplicatedName, t.CheckDuplicatedName)
//	putIfNotNull(testCaseMap, TestLinkParamActionOnDuplicatedName, t.ActionOnDuplicatedName)
//
//	return testCaseMap
//}
//
//func ToTestCaseStepMaps(testCaseSteps *[]TestCaseStep) []map[string]interface{} {
//	var testCaseStepMaps []map[string]interface{}
//
//	for _, testCaseStep := range *testCaseSteps {
//		testCaseStepMap := map[string]interface{}{}
//		putIfNotNull(testCaseStepMap, TestLinkParamStepNumber, testCaseStep.Number)
//		putIfNotNull(testCaseStepMap, TestLinkParamActions, testCaseStep.Actions)
//		putIfNotNull(testCaseStepMap, TestLinkParamExpectedResults, testCaseStep.ExpectedResults)
//		putIfNotNull(testCaseStepMap, TestLinkParamStepExecutionType, testCaseStep.ExecutionType)
//
//		testCaseStepMaps = append(testCaseStepMaps, testCaseStepMap)
//	}
//
//	return testCaseStepMaps
//}
//
//func putIfNotNull(testCaseMap map[string]interface{}, key string, val interface{}) {
//	if val != nil {
//		testCaseMap[key] = val
//	}
//}
