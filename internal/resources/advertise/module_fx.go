package advertise

import (
	"github.com/Frank0945/go-advertise/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewHttpHandler,
			fx.As(new(service.ManageHandler)),
		),

		fx.Annotate(
			NewManager,
			fx.As(new(Manager)),
		),
	),
)
