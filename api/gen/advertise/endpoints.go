// Code generated by goa v3.15.2, DO NOT EDIT.
//
// advertise endpoints
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package advertise

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "advertise" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	List   goa.Endpoint
}

// NewEndpoints wraps the methods of the "advertise" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		List:   NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "advertise" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.List = m(e.List)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "advertise".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		return nil, s.Create(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "advertise".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AdList)
		return s.List(ctx, p)
	}
}