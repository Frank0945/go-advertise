// Code generated by goa v3.15.2, DO NOT EDIT.
//
// manager HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/Frank0945/go-advertise/api/design -o api

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	manager "github.com/Frank0945/go-advertise/api/gen/manager"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateAdResponse returns an encoder for responses returned by the
// manager create_ad endpoint.
func EncodeCreateAdResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*manager.CreateAdResult)
		enc := encoder(ctx, w)
		body := NewCreateAdResponseBody(res)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateAdRequest returns a decoder for requests sent to the manager
// create_ad endpoint.
func DecodeCreateAdRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateAdRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateAdRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateAdPayload(&body)

		return payload, nil
	}
}

// EncodeListAdsResponse returns an encoder for responses returned by the
// manager list_ads endpoint.
func EncodeListAdsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*manager.Ad)
		enc := encoder(ctx, w)
		body := NewListAdsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListAdsRequest returns a decoder for requests sent to the manager
// list_ads endpoint.
func DecodeListAdsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			offset   int
			limit    int
			ageStart *int
			ageEnd   *int
			gender   []string
			country  []string
			platform []string
			err      error
		)
		{
			offsetRaw := r.URL.Query().Get("offset")
			if offsetRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("offset", "query string"))
			}
			v, err2 := strconv.ParseInt(offsetRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("offset", offsetRaw, "integer"))
			}
			offset = int(v)
		}
		if offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
		}
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("limit", "query string"))
			} else {
				v, err2 := strconv.ParseInt(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
				}
				limit = int(v)
			}
		}
		if limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
		}
		{
			ageStartRaw := r.URL.Query().Get("age_start")
			if ageStartRaw != "" {
				v, err2 := strconv.ParseInt(ageStartRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("age_start", ageStartRaw, "integer"))
				}
				pv := int(v)
				ageStart = &pv
			}
		}
		if ageStart != nil {
			if *ageStart < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("age_start", *ageStart, 1, true))
			}
		}
		if ageStart != nil {
			if *ageStart > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("age_start", *ageStart, 100, false))
			}
		}
		{
			ageEndRaw := r.URL.Query().Get("age_end")
			if ageEndRaw != "" {
				v, err2 := strconv.ParseInt(ageEndRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("age_end", ageEndRaw, "integer"))
				}
				pv := int(v)
				ageEnd = &pv
			}
		}
		if ageEnd != nil {
			if *ageEnd < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("age_end", *ageEnd, 1, true))
			}
		}
		if ageEnd != nil {
			if *ageEnd > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("age_end", *ageEnd, 100, false))
			}
		}
		gender = r.URL.Query()["gender"]
		for _, e := range gender {
			err = goa.MergeErrors(err, goa.ValidatePattern("gender[*]", e, "M|F"))
		}
		country = r.URL.Query()["country"]
		for _, e := range country {
			err = goa.MergeErrors(err, goa.ValidatePattern("country[*]", e, "TW|JP"))
		}
		platform = r.URL.Query()["platform"]
		for _, e := range platform {
			err = goa.MergeErrors(err, goa.ValidatePattern("platform[*]", e, "ios|android|web"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListAdsAdQuery(offset, limit, ageStart, ageEnd, gender, country, platform)

		return payload, nil
	}
}

// marshalManagerAdToAdResponse builds a value of type *AdResponse from a value
// of type *manager.Ad.
func marshalManagerAdToAdResponse(v *manager.Ad) *AdResponse {
	res := &AdResponse{
		Title: v.Title,
		EndAt: v.EndAt,
	}

	return res
}