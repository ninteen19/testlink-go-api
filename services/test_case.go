package services

import (
	"context"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/ninteen19/testlink-go-api/constants"
	"github.com/ninteen19/testlink-go-api/enums"
	"github.com/ninteen19/testlink-go-api/models"
	"net/http"
)

type ITestLinkOutbound interface {
	TestLinkXmlRpcCall(ctx context.Context, method string, args map[string]interface{}) (*http.Response, error)
}

type TestCaseService struct {
	TestLinkOutbound ITestLinkOutbound
}

func NewTestCaseService(testLinkOutbound ITestLinkOutbound) *TestCaseService {
	return &TestCaseService{TestLinkOutbound: testLinkOutbound}
}

func (s *TestCaseService) CreateTestCase(
	ctx context.Context,
	testCaseName string,
	testSuiteId int,
	testProjectId int,
	authorLogin string,
	summary string,
	steps []models.TestCaseStep,
	preconditions string,
	status enums.TestCaseStatus,
	importance enums.TestImportance,
	execution enums.ExecutionType,
	order int,
	internalId int,
	checkDuplicatedName bool,
	actionOnDuplicatedName enums.ActionOnDuplicate,
) (interface{}, error) {
	var id int

	testCase := &models.TestCase{
		Id:                     id,
		Name:                   testCaseName,
		TestSuiteId:            testSuiteId,
		TestProjectId:          testProjectId,
		AuthorLogin:            authorLogin,
		Summary:                summary,
		Steps:                  steps,
		Preconditions:          preconditions,
		TestCaseStatus:         status,
		TestImportance:         importance,
		ExecutionType:          execution,
		ExecutionOrder:         nil,
		Order:                  order,
		InternalId:             internalId,
		ExternalId:             nil,
		FullExternalId:         nil,
		CheckDuplicatedName:    checkDuplicatedName,
		ActionOnDuplicatedName: actionOnDuplicatedName,
		VersionId:              nil,
		Version:                nil,
		ParentId:               nil,
		CustomFields:           nil,
		ExecutionStatus:        nil,
		Platform:               nil,
		FeatureId:              nil,
		Keywords:               nil,
	}

	executionData := testCase.ToMap()
	resp, err := s.TestLinkOutbound.TestLinkXmlRpcCall(ctx, constants.CreateTestCase, executionData)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var respBodyArrays []interface{}
	err = xml.DecodeClientResponse(resp.Body, &respBodyArrays)

	if err != nil {
		return nil, err
	}

	var respMap = respBodyArrays[0].(map[string]interface{})

	testCase.Id = respMap["id"].(int)
	testCase.Version = respMap["additionalInfo"].(map[string]int)["version_number"]

	return testCase, nil
}
