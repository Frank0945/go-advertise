package manage

import (
	"context"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
)

type Handler struct {
}

func NewHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateAd(ctx context.Context, p *api.CreateAdPayload) error {

	return nil
}

func (h *Handler) ListAds(ctx context.Context, list *api.AdOverview) ([]*api.Ad, error) {

	return nil, nil
}
