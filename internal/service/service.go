package service

import (
	"context"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
	"go.uber.org/fx"
)

type ManageHandler interface {
	Create(context.Context, *api.CreatePayload) error
	List(context.Context, *api.AdList) ([]*api.Ads, error)
}

type service struct {
	ManageHandler
}

type Params struct {
	fx.In

	MgHndlr ManageHandler
}

func New(p Params) api.Service {
	return &service{
		ManageHandler: p.MgHndlr,
	}
}
