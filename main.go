package main

import (
	"github.com/Frank0945/go-advertise/api/gen/http/ad/server"
	"github.com/Frank0945/go-advertise/internal/config"
	"github.com/Frank0945/go-advertise/pkg/goafx/muxfx"
	"github.com/Frank0945/go-advertise/pkg/httpfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,

		// Muxer
		muxfx.Module,

		// Server
		fx.Provide(server.New),

		// HTTP Server & Execution
		httpfx.Module,
	).Run()
}
