// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetV1AppruntimesHandlerFunc turns a function with the right signature into a get v1 appruntimes handler
type GetV1AppruntimesHandlerFunc func(GetV1AppruntimesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1AppruntimesHandlerFunc) Handle(params GetV1AppruntimesParams) middleware.Responder {
	return fn(params)
}

// GetV1AppruntimesHandler interface for that can handle valid get v1 appruntimes params
type GetV1AppruntimesHandler interface {
	Handle(GetV1AppruntimesParams) middleware.Responder
}

// NewGetV1Appruntimes creates a new http.Handler for the get v1 appruntimes operation
func NewGetV1Appruntimes(ctx *middleware.Context, handler GetV1AppruntimesHandler) *GetV1Appruntimes {
	return &GetV1Appruntimes{Context: ctx, Handler: handler}
}

/*GetV1Appruntimes swagger:route GET /v1/appruntimes app-runtimes getV1Appruntimes

Gets some app runtimes

Returns a list containing all app runtimes.

*/
type GetV1Appruntimes struct {
	Context *middleware.Context
	Handler GetV1AppruntimesHandler
}

func (o *GetV1Appruntimes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetV1AppruntimesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
