package service

import (
	"context"

	api "github.com/Frank0945/go-advertise/api/gen/manager"
	"go.uber.org/fx"
)

type ManageHandler interface {
	CreateAd(context.Context, *api.CreateAdPayload) (*api.CreateAdResult, error)
	ListAds(context.Context, *api.AdQuery) ([]*api.Ad, error)
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
