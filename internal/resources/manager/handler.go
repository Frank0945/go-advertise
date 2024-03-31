package manager

import (
	"context"
	"fmt"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
)

type Handler struct {
}

func NewHttpHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(ctx context.Context, p *api.CreatePayload) error {
	if p.Conditions.AgeEnd != nil && p.Conditions.AgeStart != nil {
		if *p.Conditions.AgeEnd < *p.Conditions.AgeStart {
			return fmt.Errorf("`age_end` must be greater than/equal to `age_start`")
		}
	}

	return nil
}

func (h *Handler) List(ctx context.Context, list *api.AdList) ([]*api.Ads, error) {

	return nil, nil
}
