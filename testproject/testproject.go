package testproject

import (
	"github.com/ninteen19/testlink-go-api"
	"github.com/ninteen19/testlink-go-api/outbounds"
)

type ITestLinkOutbound interface {
	TestLinkXmlRpcCall(method string, request testlink.BaseRequest, resp interface{}) error
}

var outbound ITestLinkOutbound

func Get(
	request *GetRequest,
) (*Response, error) {
	resp := &Response{}
	err := outbound.TestLinkXmlRpcCall(testlink.TestLinkMethodGetTestProjectByName, &testlink.TestProjectRequest{
		TestProjectName: request.ProjectName,
	}, resp)

	return resp, err
}

func init() {
	outbound = &outbounds.TestLinkOutbound{
		Config: &testlink.Conf,
	}
}
