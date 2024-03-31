package main

import (
	"github.com/Frank0945/go-advertise/api/gen/advertise"
	"github.com/Frank0945/go-advertise/api/gen/http/advertise/server"
	"github.com/Frank0945/go-advertise/internal/config"
	"github.com/Frank0945/go-advertise/internal/resources/manager"
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
		manager.Module,

		// Endpoints
		fx.Provide(
			service.New,
			advertise.NewEndpoints,
		),

		// Muxer
		muxfx.Module,

		// Server
		fx.Provide(server.New),

		// HTTP Server
		httpfx.Module,
	).Run()
}
