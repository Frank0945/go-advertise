package manager

import (
	"context"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
)

type Handler struct {
}

func NewHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(context.Context, *api.CreatePayload) error {
	return nil
}

func (h *Handler) List(context.Context, *api.AdList) ([]*api.Ads, error) {
	return nil, nil
}
