package testcase

import (
	"github.com/ninteen19/testlink-go-api"
	"github.com/ninteen19/testlink-go-api/outbounds"
)

type ITestLinkOutbound interface {
	TestLinkXmlRpcCall(method string, request testlink.BaseRequest, resp interface{}) error
}

var outbound ITestLinkOutbound

func Create(
	request *CreateRequest,
) error {
	err := request.Validate()

	if err != nil {
		return err
	}

	testCase := &testlink.TestCaseRequest{
		Name:                   request.TestCaseName,
		TestSuiteId:            request.TestSuiteId,
		TestProjectId:          request.TestProjectId,
		AuthorLogin:            request.AuthorLogin,
		Summary:                request.Summary,
		Steps:                  request.Steps,
		Preconditions:          request.Preconditions,
		TestCaseStatus:         request.Status.Value(),
		TestImportance:         request.Importance.Value(),
		ExecutionType:          request.Execution.Value(),
		Order:                  request.Order,
		InternalId:             request.InternalId,
		ExternalId:             -1,
		CheckDuplicatedName:    request.CheckDuplicatedName,
		ActionOnDuplicatedName: request.ActionOnDuplicatedName.Value(),
	}

	return outbound.TestLinkXmlRpcCall(testlink.TestLinkMethodCreateTestCase, testCase, &[]struct{}{})
}

func init() {
	outbound = &outbounds.TestLinkOutbound{
		Config: &testlink.Conf,
	}
}
