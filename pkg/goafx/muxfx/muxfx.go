package muxfx

import (
	"go.uber.org/fx"
	goahttp "goa.design/goa/v3/http"
	goaHttpMw "goa.design/goa/v3/http/middleware"
)

// Module provides a muxer instance.
var Module = fx.Provide(
	fx.Annotate(
		NewSmartRedirectMuxer,
		fx.As(new(goahttp.Muxer)),
	),
)

// NewSmartRedirectMuxer returns a muxer with smart redirect slashes middleware.
func NewSmartRedirectMuxer() goahttp.ResolverMuxer {
	mux := goahttp.NewMuxer()
	mux.Use(goaHttpMw.SmartRedirectSlashes)
	return mux
}
