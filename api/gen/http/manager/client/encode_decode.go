// Code generated by goa v3.15.2, DO NOT EDIT.
//
// manager HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	manager "github.com/Frank0945/go-advertise/api/gen/manager"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateAdRequest instantiates a HTTP request object with method and path
// set to call the "manager" service "create_ad" endpoint
func (c *Client) BuildCreateAdRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAdManagerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("manager", "create_ad", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateAdRequest returns an encoder for requests sent to the manager
// create_ad server.
func EncodeCreateAdRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*manager.CreateAdPayload)
		if !ok {
			return goahttp.ErrInvalidType("manager", "create_ad", "*manager.CreateAdPayload", v)
		}
		body := NewCreateAdRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("manager", "create_ad", err)
		}
		return nil
	}
}

// DecodeCreateAdResponse returns a decoder for responses returned by the
// manager create_ad endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateAdResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateAdResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("manager", "create_ad", err)
			}
			err = ValidateCreateAdResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("manager", "create_ad", err)
			}
			res := NewCreateAdResultCreated(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("manager", "create_ad", resp.StatusCode, string(body))
		}
	}
}

// BuildListAdsRequest instantiates a HTTP request object with method and path
// set to call the "manager" service "list_ads" endpoint
func (c *Client) BuildListAdsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAdsManagerPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("manager", "list_ads", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListAdsRequest returns an encoder for requests sent to the manager
// list_ads server.
func EncodeListAdsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*manager.AdQuery)
		if !ok {
			return goahttp.ErrInvalidType("manager", "list_ads", "*manager.AdQuery", v)
		}
		values := req.URL.Query()
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.AgeStart != nil {
			values.Add("age_start", fmt.Sprintf("%v", *p.AgeStart))
		}
		if p.AgeEnd != nil {
			values.Add("age_end", fmt.Sprintf("%v", *p.AgeEnd))
		}
		for _, value := range p.Gender {
			values.Add("gender", value)
		}
		for _, value := range p.Country {
			values.Add("country", value)
		}
		for _, value := range p.Platform {
			values.Add("platform", value)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListAdsResponse returns a decoder for responses returned by the
// manager list_ads endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeListAdsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListAdsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("manager", "list_ads", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateAdResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("manager", "list_ads", err)
			}
			res := NewListAdsAdOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("manager", "list_ads", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAdResponseToManagerAd builds a value of type *manager.Ad from a
// value of type *AdResponse.
func unmarshalAdResponseToManagerAd(v *AdResponse) *manager.Ad {
	res := &manager.Ad{
		Title: *v.Title,
		EndAt: *v.EndAt,
	}

	return res
}
