package main

import (
	"github.com/Frank0945/go-advertise/api/gen/http/manager/server"
	"github.com/Frank0945/go-advertise/api/gen/manager"
	"github.com/Frank0945/go-advertise/internal/config"
	"github.com/Frank0945/go-advertise/internal/entities/advertise"
	adResource "github.com/Frank0945/go-advertise/internal/resources/advertise"
	"github.com/Frank0945/go-advertise/internal/service"
	"github.com/Frank0945/go-advertise/pkg/databasefx"
	"github.com/Frank0945/go-advertise/pkg/goafx/muxfx"
	"github.com/Frank0945/go-advertise/pkg/httpfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		databasefx.Module,
		config.Module,

		// Features
		adResource.Module,

		// Endpoints
		fx.Provide(
			service.New,
			manager.NewEndpoints,
		),

		fx.Provide(
			fx.Annotate(
				advertise.NewAdMapper,
				fx.As(new(advertise.AdMapper)),
			),
		),

		// Muxer
		muxfx.Module,

		// Server
		fx.Provide(server.New),

		// HTTP Server
		httpfx.Module,
	).Run()
}
