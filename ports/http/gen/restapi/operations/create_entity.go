// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateEntityHandlerFunc turns a function with the right signature into a create entity handler
type CreateEntityHandlerFunc func(CreateEntityParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateEntityHandlerFunc) Handle(params CreateEntityParams) middleware.Responder {
	return fn(params)
}

// CreateEntityHandler interface for that can handle valid create entity params
type CreateEntityHandler interface {
	Handle(CreateEntityParams) middleware.Responder
}

// NewCreateEntity creates a new http.Handler for the create entity operation
func NewCreateEntity(ctx *middleware.Context, handler CreateEntityHandler) *CreateEntity {
	return &CreateEntity{Context: ctx, Handler: handler}
}

/* CreateEntity swagger:route POST /entities createEntity


Creates new Entity


*/
type CreateEntity struct {
	Context *middleware.Context
	Handler CreateEntityHandler
}

func (o *CreateEntity) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateEntityParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
