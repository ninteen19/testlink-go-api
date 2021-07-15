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

type TestCaseStep struct {
	Id                int    `json:"id"`
	TestCaseVersionId int    `json:"testCaseVersionId"`
	Number            int    `json:"number" xmlrpc:"step_number"`
	Actions           string `json:"actions" xmlrpc:"actions"`
	ExpectedResults   string `json:"expectedResults" xmlrpc:"expected_results"`
	Active            bool   `json:"active"`
	ExecutionType     int    `json:"executionType" xmlrpc:"execution_type"`
}
