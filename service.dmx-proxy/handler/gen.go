// Code generated by jrpc. DO NOT EDIT.

package handler

import (
	context "context"
	http "net/http"

	network "github.com/jakewright/home-automation/libraries/go/network"
	oops "github.com/jakewright/home-automation/libraries/go/oops"
	router "github.com/jakewright/home-automation/libraries/go/router"
	slog "github.com/jakewright/home-automation/libraries/go/slog"
	def "github.com/jakewright/home-automation/service.dmx-proxy/def"
)

type controller interface {
	Set(*request, *def.SetRequest) (*def.SetResponse, error)
}

type request struct {
	context.Context
	*http.Request
}

// NewRouter returns a router with appropriate handlers set
func NewRouter(c controller) *router.Router {
	r := router.New()

	r.Handle("PUT", "/set", func(w http.ResponseWriter, r *http.Request) {
		body := &def.SetRequest{}
		if err := network.DecodeRequest(r, body); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to decode request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		if err := body.Validate(); err != nil {
			err = oops.Wrap(err, oops.ErrBadRequest, "failed to validate request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		req := &request{
			Context: r.Context(),
			Request: r,
		}

		rsp, err := c.Set(req, body)
		if err != nil {
			err = oops.WithMessage(err, "failed to handle request")
			slog.Error(err)
			network.WriteJSONResponse(w, err)
			return
		}

		network.WriteJSONResponse(w, rsp)
	})

	return r
}
