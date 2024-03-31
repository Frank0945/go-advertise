// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ad HTTP client types
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package client

import (
	ad "github.com/Frank0945/go-advertise/api/gen/ad"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "ad" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title of AD
	Title string `form:"title" json:"title" xml:"title"`
	// Start time of AD
	StartAt string `form:"start_at" json:"start_at" xml:"start_at"`
	// End time of AD
	EndAt string `form:"end_at" json:"end_at" xml:"end_at"`
	// Start age of target
	AgeStart *int `form:"age_start,omitempty" json:"age_start,omitempty" xml:"age_start,omitempty"`
	// End age of target
	AgeEnd *int `form:"age_end,omitempty" json:"age_end,omitempty" xml:"age_end,omitempty"`
	// Gender of target
	Gender *string `form:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// Country of target
	Country *string `form:"Country,omitempty" json:"Country,omitempty" xml:"Country,omitempty"`
	// Platform of target
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
}

// ListRequestBody is the type of the "ad" service "list" endpoint HTTP request
// body.
type ListRequestBody struct {
	// Offset of AD
	Offset int `form:"offset" json:"offset" xml:"offset"`
	// Limit of AD
	Limit int `form:"limit" json:"limit" xml:"limit"`
	// Start age of target
	AgeStart *int `form:"age_start,omitempty" json:"age_start,omitempty" xml:"age_start,omitempty"`
	// End age of target
	AgeEnd *int `form:"age_end,omitempty" json:"age_end,omitempty" xml:"age_end,omitempty"`
	// Gender of target
	Gender *string `form:"gender,omitempty" json:"gender,omitempty" xml:"gender,omitempty"`
	// Country of target
	Country *string `form:"Country,omitempty" json:"Country,omitempty" xml:"Country,omitempty"`
	// Platform of target
	Platform *string `form:"platform,omitempty" json:"platform,omitempty" xml:"platform,omitempty"`
}

// ListResponseBody is the type of the "ad" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Title of AD
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// End time of AD
	EndAt *string `form:"end_at,omitempty" json:"end_at,omitempty" xml:"end_at,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "ad" service.
func NewCreateRequestBody(p *ad.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Title:    p.Title,
		StartAt:  p.StartAt,
		EndAt:    p.EndAt,
		AgeStart: p.AgeStart,
		AgeEnd:   p.AgeEnd,
		Gender:   p.Gender,
		Country:  p.Country,
		Platform: p.Platform,
	}
	return body
}

// NewListRequestBody builds the HTTP request body from the payload of the
// "list" endpoint of the "ad" service.
func NewListRequestBody(p *ad.AdSearchPayload) *ListRequestBody {
	body := &ListRequestBody{
		Offset:   p.Offset,
		Limit:    p.Limit,
		AgeStart: p.AgeStart,
		AgeEnd:   p.AgeEnd,
		Gender:   p.Gender,
		Country:  p.Country,
		Platform: p.Platform,
	}
	return body
}

// NewListAdsOK builds a "ad" service "list" endpoint result from a HTTP "OK"
// response.
func NewListAdsOK(body *ListResponseBody) *ad.Ads {
	v := &ad.Ads{
		Title: *body.Title,
		EndAt: *body.EndAt,
	}

	return v
}

// ValidateListResponseBody runs the validations defined on ListResponseBody
func ValidateListResponseBody(body *ListResponseBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.EndAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("end_at", "body"))
	}
	return
}
