package outbounds

import (
	"alexejk.io/go-xmlrpc"
	"context"
	"github.com/ninteen19/testlink-go-api"
)

type TestLinkOutbound struct {
	Config *testlink.Config
}

//ignore ctx for now
func (o *TestLinkOutbound) TestLinkXmlRpcCallWithContext(ctx context.Context, method string, testCase *testlink.TestCaseRequest) (*testlink.TestCaseResponse, error) {
	testCase.DevKey = o.Config.Key
	client, _ := xmlrpc.NewClient(o.Config.Url)
	result := &testlink.TestCaseResponse{}
	err := client.Call(method, testCase, result)
	return result, err
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, testCase *testlink.TestCaseRequest) (*testlink.TestCaseResponse, error) {
	testCase.DevKey = o.Config.Key
	client, _ := xmlrpc.NewClient(o.Config.Url)
	result := &testlink.TestCaseResponse{}
	err := client.Call(method, testCase, result)
	return result, err
}
