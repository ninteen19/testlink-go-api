package outbounds

import (
	"bytes"
	"context"
	"github.com/divan/gorilla-xmlrpc/xml"
	"net/http"
)

type TestLinkOutbound struct {
	Url string
}

func NewTestLinkOutbound(url string) *TestLinkOutbound {
	return &TestLinkOutbound{Url: url}
}

//ignore ctx for now
func (o *TestLinkOutbound) TestLinkXmlRpcCall(ctx context.Context, method string, args map[string]interface{}) (*http.Response, error) {
	buf, _ := xml.EncodeClientRequest(method, &args)
	return http.Post(o.Url, "text/xml", bytes.NewBuffer(buf))
}
