// Code generated by jrpc. DO NOT EDIT.

package dmxproxydef

import (
	context "context"

	rpc "github.com/jakewright/home-automation/libraries/go/rpc"
)

// Request builds an RPC request
func (m *SetRequest) Request() *rpc.Request {
	return &rpc.Request{
		Method: "PUT",
		URL:    "service.dmx-proxy/set",
		Body:   m,
	}
}

// Do performs the request
func (m *SetRequest) Do(ctx context.Context) (*SetResponse, error) {
	rsp := &SetResponse{}
	_, err := rpc.Do(ctx, m.Request(), rsp)
	return rsp, err
}
