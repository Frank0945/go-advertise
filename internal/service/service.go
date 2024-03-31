package service

import (
	"context"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
	"go.uber.org/fx"
)

type AdvertiseHandler interface {
	Create(context.Context, *api.CreatePayload) error
	List(context.Context, *api.AdList) ([]*api.Ads, error)
}

type service struct {
	AdvertiseHandler
}

type Params struct {
	fx.In

	AdHndlr AdvertiseHandler
}

func New(p Params) api.Service {
	return &service{
		AdvertiseHandler: p.AdHndlr,
	}
}
