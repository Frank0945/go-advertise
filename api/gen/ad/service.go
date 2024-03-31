// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ad service
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package ad

import (
	"context"
)

// Service is the ad service interface.
type Service interface {
	// Create a new edge
	Create(context.Context, *CreatePayload) (err error)
	// List all ADs by filter
	List(context.Context, *AdSearchPayload) (res *Ads, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ad"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "ad"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"create", "list"}

// AdSearchPayload is the payload type of the ad service list method.
type AdSearchPayload struct {
	// Offset of AD
	Offset int
	// Limit of AD
	Limit int
	// Start age of target
	AgeStart *int
	// End age of target
	AgeEnd *int
	// Gender of target
	Gender *string
	// Country of target
	Country *string
	// Platform of target
	Platform *string
}

// Ads is the result type of the ad service list method.
type Ads struct {
	// Title of AD
	Title string
	// End time of AD
	EndAt string
}

// CreatePayload is the payload type of the ad service create method.
type CreatePayload struct {
	// Title of AD
	Title string
	// Start time of AD
	StartAt string
	// End time of AD
	EndAt string
	// Start age of target
	AgeStart *int
	// End age of target
	AgeEnd *int
	// Gender of target
	Gender *string
	// Country of target
	Country *string
	// Platform of target
	Platform *string
}
