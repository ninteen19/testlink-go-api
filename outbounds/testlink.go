package outbounds

import (
	"bytes"
	"context"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/ninteen19/testlink-go-api"
	"github.com/ninteen19/testlink-go-api/constants"
	"net/http"
)

type TestLinkOutbound struct {
	Config *testlink.Config
}

//ignore ctx for now
func (o *TestLinkOutbound) TestLinkXmlRpcCallWithContext(ctx context.Context, method string, args map[string]interface{}) (*http.Response, error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	if len(args) > 0 {
		args[constants.DevKey] = o.Config.Key
	}
	return http.Post(o.Config.Url, "text/xml", bytes.NewBuffer(buf))
}

func (o *TestLinkOutbound) TestLinkXmlRpcCall(method string, args map[string]interface{}) (*http.Response, error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	if len(args) > 0 {
		args[constants.DevKey] = o.Config.Key
	}
	return http.Post(o.Config.Url, "text/xml", bytes.NewBuffer(buf))
}
