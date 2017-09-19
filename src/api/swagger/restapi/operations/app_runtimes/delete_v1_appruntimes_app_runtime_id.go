// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteV1AppruntimesAppRuntimeIDHandlerFunc turns a function with the right signature into a delete v1 appruntimes app runtime ID handler
type DeleteV1AppruntimesAppRuntimeIDHandlerFunc func(DeleteV1AppruntimesAppRuntimeIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteV1AppruntimesAppRuntimeIDHandlerFunc) Handle(params DeleteV1AppruntimesAppRuntimeIDParams) middleware.Responder {
	return fn(params)
}

// DeleteV1AppruntimesAppRuntimeIDHandler interface for that can handle valid delete v1 appruntimes app runtime ID params
type DeleteV1AppruntimesAppRuntimeIDHandler interface {
	Handle(DeleteV1AppruntimesAppRuntimeIDParams) middleware.Responder
}

// NewDeleteV1AppruntimesAppRuntimeID creates a new http.Handler for the delete v1 appruntimes app runtime ID operation
func NewDeleteV1AppruntimesAppRuntimeID(ctx *middleware.Context, handler DeleteV1AppruntimesAppRuntimeIDHandler) *DeleteV1AppruntimesAppRuntimeID {
	return &DeleteV1AppruntimesAppRuntimeID{Context: ctx, Handler: handler}
}

/*DeleteV1AppruntimesAppRuntimeID swagger:route DELETE /v1/appruntimes/{appRuntimeId} app-runtimes deleteV1AppruntimesAppRuntimeId

Deletes an app runtime

Delete a single app runtime identified via its id

*/
type DeleteV1AppruntimesAppRuntimeID struct {
	Context *middleware.Context
	Handler DeleteV1AppruntimesAppRuntimeIDHandler
}

func (o *DeleteV1AppruntimesAppRuntimeID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteV1AppruntimesAppRuntimeIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
