package outbounds

import (
	"github.com/kolo/xmlrpc"
	"github.com/ninteen19/testlink-go-api"
)

type TestLinkOutbound struct {
	Config *testlink.Config
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, testCase *testlink.TestCaseRequest) error {
	testCase.DevKey = o.Config.Key
	client, _ := xmlrpc.NewClient(o.Config.Url, nil)
	return client.Call(method, testCase, &[]struct{}{})
}
