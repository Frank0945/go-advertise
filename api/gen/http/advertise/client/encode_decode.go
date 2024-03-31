// Code generated by goa v3.15.2, DO NOT EDIT.
//
// advertise HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	advertise "github.com/Frank0945/go-advertise/api/gen/advertise"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "advertise" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAdvertisePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("advertise", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the advertise
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*advertise.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("advertise", "create", "*advertise.CreatePayload", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("advertise", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// advertise create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("advertise", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "advertise" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAdvertisePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("advertise", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the advertise list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*advertise.AdList)
		if !ok {
			return goahttp.ErrInvalidType("advertise", "list", "*advertise.AdList", v)
		}
		body := NewListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("advertise", "list", err)
		}
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the advertise
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("advertise", "list", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateAdsResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("advertise", "list", err)
			}
			res := NewListAdsOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("advertise", "list", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAdsResponseToAdvertiseAds builds a value of type *advertise.Ads
// from a value of type *AdsResponse.
func unmarshalAdsResponseToAdvertiseAds(v *AdsResponse) *advertise.Ads {
	res := &advertise.Ads{
		Title: *v.Title,
		EndAt: *v.EndAt,
	}

	return res
}