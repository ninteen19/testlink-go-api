package outbounds

import (
	"github.com/kolo/xmlrpc"
	"github.com/ninteen19/testlink-go-api"
)

type TestLinkOutbound struct {
	Config *testlink.Config
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, baseRequest testlink.BaseRequest, response interface{}) error {
	baseRequest.SetDevKey(o.Config.Key)
	client, _ := xmlrpc.NewClient(o.Config.Url, nil)
	return client.Call(method, baseRequest, response)
}
