package httpfx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Frank0945/advertise/internal/config"
	"github.com/Frank0945/advertise/pkg/goafx"

	apisvr "github.com/Frank0945/advertise/api/gen/http/ad/server"

	"go.uber.org/fx"
	goahttp "goa.design/goa/v3/http"
)

// Module creates a new http server and invokes it.
var Module = fx.Options(
	// goa http server dependencies
	fx.Provide(
		goafx.NewDefaultRequestDecoder,
		goafx.NewDefaultResponseEncoder,
		goafx.NewDefaultErrorHandler,
		goafx.NewDefaultFormatter,
	),
	fx.Provide(
		New,
		NewHTTPServer,
	),
	fx.Invoke(
		func(*http.Server) {},
	),
)

type server struct {
	mux goahttp.Muxer
	svr *apisvr.Server
	cfg *config.Config
}

type Params struct {
	fx.In

	Mux    goahttp.Muxer
	Server *apisvr.Server
	Config *config.Config
}

func New(p Params) *server {
	return &server{
		mux: p.Mux,
		svr: p.Server,
		cfg: p.Config,
	}
}

func NewHTTPServer(lc fx.Lifecycle, s *server) *http.Server {
	apisvr.Mount(s.mux, s.svr)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: s.mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
