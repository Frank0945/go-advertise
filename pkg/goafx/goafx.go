package goafx

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

type (
	RequestDecoder  = func(*http.Request) goahttp.Decoder
	ResponseEncoder = func(context.Context, http.ResponseWriter) goahttp.Encoder
	ErrorHandler    = func(context.Context, http.ResponseWriter, error)
	Formatter       = func(context.Context, error) goahttp.Statuser
)

func NewDefaultRequestDecoder() RequestDecoder {
	return goahttp.RequestDecoder
}

func NewDefaultResponseEncoder() ResponseEncoder {
	return goahttp.ResponseEncoder
}

func NewDefaultErrorHandler() ErrorHandler {
	return nil
}

func NewDefaultFormatter() Formatter {
	return nil
}
