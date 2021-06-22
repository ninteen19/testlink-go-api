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
func (o *TestLinkOutbound) TestLinkXmlRpcCallWithContext(ctx context.Context, method string, testCase *testlink.TestCase) (*testlink.TestCase, error) {
	testCase.DevKey = o.Config.Key
	client, _ := xmlrpc.NewClient(o.Config.Url)
	result := &testlink.TestCase{}
	_ = client.Call(method, testCase, result)
	return result, nil
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, testCase *testlink.TestCase) (*testlink.TestCase, error) {
	testCase.DevKey = o.Config.Key
	client, _ := xmlrpc.NewClient(o.Config.Url)
	result := &testlink.TestCase{}
	_ = client.Call(method, testCase, result)
	return result, nil
}
