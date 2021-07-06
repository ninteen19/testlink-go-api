package testcase

import (
	"context"
	"github.com/ninteen19/testlink-go-api"
	"github.com/ninteen19/testlink-go-api/outbounds"
)

type ITestLinkOutbound interface {
	TestLinkXmlRpcCallWithContext(ctx context.Context, method string, testCase *testlink.TestCaseRequest) (*testlink.TestCaseResponse, error)
	TestLinkXmlRpcCall(method string, testCase *testlink.TestCaseRequest) error
}

var outbound ITestLinkOutbound

func Create(
	testCaseName string,
	testSuiteId int,
	testProjectId int,
	authorLogin string,
	summary string,
	steps []testlink.TestCaseStep,
	preconditions string,
	status testlink.TestCaseStatus,
	importance testlink.TestImportance,
	execution testlink.ExecutionType,
	order int,
	internalId int,
	checkDuplicatedName bool,
	actionOnDuplicatedName testlink.ActionOnDuplicate,
) error {
	var id int

	testCase := &testlink.TestCaseRequest{
		Id:             id,
		Name:           testCaseName,
		TestSuiteId:    testSuiteId,
		TestProjectId:  testProjectId,
		AuthorLogin:    authorLogin,
		Summary:        summary,
		Steps:          steps,
		Preconditions:  preconditions,
		TestCaseStatus: status.Value(),
		TestImportance: importance.Value(),
		ExecutionType:  execution.Value(),
		//ExecutionOrder:         -1,
		Order:      order,
		InternalId: internalId,
		ExternalId: -1,
		//FullExternalId:         "",
		CheckDuplicatedName:    checkDuplicatedName,
		ActionOnDuplicatedName: actionOnDuplicatedName.Value(),
		//VersionId:              -1,
		//Version:                -1,
		//ParentId:               -1,
		//CustomFields:           []testlink.CustomField{},
		//ExecutionStatus:        testlink.ExecutionStatusNotRun,
		//Platform:               testlink.Platform{},
		//FeatureId:              -1,
		//Keywords:               []string{},
	}

	//executionData := testCase.ToMap()
	return outbound.TestLinkXmlRpcCall(testlink.TestLinkMethodCreateTestCase, testCase)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer resp.Body.Close()
	//var respBodyArrays []interface{}
	//err = xml.DecodeClientResponse(resp.Body, &respBodyArrays)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//var respMap = respBodyArrays[0].(map[string]interface{})
	//
	//testCase.Id = respMap["id"].(int)
	//testCase.Version = respMap["additionalInfo"].(map[string]int)["version_number"]
	//
	//return testCase, nil
}

func CreateWithContext(
	ctx context.Context,
	testCaseName string,
	testSuiteId int,
	testProjectId int,
	authorLogin string,
	summary string,
	steps []testlink.TestCaseStep,
	preconditions string,
	status testlink.TestCaseStatus,
	importance testlink.TestImportance,
	execution testlink.ExecutionType,
	order int,
	internalId int,
	checkDuplicatedName bool,
	actionOnDuplicatedName testlink.ActionOnDuplicate,
) (*testlink.TestCaseResponse, error) {
	var id int

	testCase := &testlink.TestCaseRequest{
		Id:             id,
		Name:           testCaseName,
		TestSuiteId:    testSuiteId,
		TestProjectId:  testProjectId,
		AuthorLogin:    authorLogin,
		Summary:        summary,
		Steps:          steps,
		Preconditions:  preconditions,
		TestCaseStatus: status.Value(),
		TestImportance: importance.Value(),
		ExecutionType:  execution.Value(),
		//ExecutionOrder:         -1,
		Order:      order,
		InternalId: internalId,
		ExternalId: -1,
		//FullExternalId:         "",
		CheckDuplicatedName:    checkDuplicatedName,
		ActionOnDuplicatedName: actionOnDuplicatedName.Value(),
		//VersionId:              -1,
		Version: -1,
		//ParentId:               -1,
		//CustomFields:           []testlink.CustomField{},
		//ExecutionStatus:        testlink.ExecutionStatusNotRun,
		//Platform:               testlink.Platform{},
		//FeatureId:              -1,
		//Keywords:               []string{},
	}

	//executionData := testCase.ToMap()
	return outbound.TestLinkXmlRpcCallWithContext(ctx, testlink.TestLinkMethodCreateTestCase, testCase)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer resp.Body.Close()
	//var respBodyArrays []interface{}
	//err = xml.DecodeClientResponse(resp.Body, &respBodyArrays)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//var respMap = respBodyArrays[0].(map[string]interface{})
	//
	//testCase.Id = respMap["id"].(int)
	//testCase.Version = respMap["additionalInfo"].(map[string]int)["version_number"]
	//
	//return testCase, nil
}

func init() {
	outbound = &outbounds.TestLinkOutbound{
		Config: &testlink.Conf,
	}
}
