package outbounds

import (
	"bytes"
	"context"
	"encoding/xml"
	"github.com/ninteen19/testlink-go-api"
	"net/http"
)

type TestLinkOutbound struct {
	Config *testlink.Config
}

//ignore ctx for now
func (o *TestLinkOutbound) TestLinkXmlRpcCallWithContext(ctx context.Context, method string, testCase *testlink.TestCase) (*http.Response, error) {
	testCase.DevKey = o.Config.Key
	buf, _ := xml.Marshal(testCase)
	//if len(args) > 0 {
	//	args[testlink.TestLinkParamDevKey] = o.Config.Key
	//}
	return http.Post(o.Config.Url, "text/xml", bytes.NewBuffer(buf))
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, testCase *testlink.TestCase) (*http.Response, error) {
	testCase.DevKey = o.Config.Key
	buf, _ := xml.Marshal(testCase)
	return http.Post(o.Config.Url, "text/xml", bytes.NewBuffer(buf))
}
