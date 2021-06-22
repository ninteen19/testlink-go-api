package testlink

type TestCaseRequest struct {
	Id             int            `json:"id" xmlrpc:"testcaseid"`
	Name           string         `json:"name" xmlrpc:"testcasename"`
	TestSuiteId    int            `json:"testSuiteId" xmlrpc:"testsuiteid"`
	TestProjectId  int            `json:"testProjectId" xmlrpc:"testprojectid"`
	AuthorLogin    string         `json:"authorLogin" xmlrpc:"authorlogin"`
	Summary        string         `json:"summary" xmlrpc:"summary"`
	Steps          []TestCaseStep `json:"steps" xmlrpc:"steps"`
	Preconditions  string         `json:"preconditions" xmlrpc:"preconditions"`
	TestCaseStatus int            `json:"testCaseStatus" xmlrpc:"status"`
	TestImportance int            `json:"testImportance" xmlrpc:"importance"`
	ExecutionType  int            `json:"executionType" xmlrpc:"executiontype"`
	//ExecutionOrder         int               `json:"executionOrder"`
	Order      int `json:"order" xmlrpc:"order"`
	InternalId int `json:"internalId" xmlrpc:"internalid"`
	ExternalId int `json:"externalId" xmlrpc:"externalid"`
	//FullExternalId         string            `json:"fullExternalId"`
	CheckDuplicatedName    bool `json:"checkDuplicatedName" xmlrpc:"checkduplicatedname"`
	ActionOnDuplicatedName int  `json:"actionOnDuplicatedName" xmlrpc:"actiononduplicatedname"`
	//VersionId              int               `json:"versionId" `
	Version int `json:"version"`
	//ParentId               int               `json:"parentId"`
	//CustomFields           []CustomField     `json:"customFields"`
	//ExecutionStatus        ExecutionStatus   `json:"executionStatus"`
	//Platform               Platform          `json:"platform"`
	//FeatureId              int               `json:"featureId"`
	//Keywords               []string          `json:"keywords"`
	DevKey string `json:"devKey" xmlrpc:"devKey"`
}

type TestCaseResponse struct {
	Id                     int            `json:"id" xmlrpc:"testcaseid"`
	Name                   string         `json:"name" xmlrpc:"testcasename"`
	TestSuiteId            int            `json:"testSuiteId" xmlrpc:"testsuiteid"`
	TestProjectId          int            `json:"testProjectId" xmlrpc:"testprojectid"`
	AuthorLogin            string         `json:"authorLogin" xmlrpc:"authorlogin"`
	Summary                string         `json:"summary" xmlrpc:"summary"`
	Steps                  []TestCaseStep `json:"steps" xmlrpc:"steps"`
	Preconditions          string         `json:"preconditions" xmlrpc:"preconditions"`
	TestCaseStatus         int            `json:"testCaseStatus" xmlrpc:"status"`
	TestImportance         int            `json:"testImportance" xmlrpc:"importance"`
	ExecutionType          int            `json:"executionType" xmlrpc:"executiontype"`
	ExecutionOrder         int            `json:"executionOrder"`
	Order                  int            `json:"order" xmlrpc:"order"`
	InternalId             int            `json:"internalId" xmlrpc:"internalid"`
	ExternalId             int            `json:"externalId" xmlrpc:"externalid"`
	FullExternalId         string         `json:"fullExternalId"`
	CheckDuplicatedName    bool           `json:"checkDuplicatedName" xmlrpc:"checkduplicatedname"`
	ActionOnDuplicatedName int            `json:"actionOnDuplicatedName" xmlrpc:"actiononduplicatedname"`
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
	Number            int    `json:"number" xmlrpc:"step_number"`
	Actions           string `json:"actions" xmlrpc:"actions"`
	ExpectedResults   string `json:"expectedResults" xmlrpc:"expected_results"`
	Active            bool   `json:"active"`
	ExecutionType     int    `json:"executionType" xmlrpc:"execution_type"`
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
