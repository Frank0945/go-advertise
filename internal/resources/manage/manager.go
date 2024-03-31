package manage

import (
	"context"
	"fmt"

	api "github.com/Frank0945/go-advertise/api/gen/advertise"
)

type Manager interface {
	Create(ctx context.Context, p *api.CreateAdPayload) (*api.CreateAdResult, error)
	List(ctx context.Context, list *api.AdOverview) ([]*api.Ad, error)
}

type manager struct {
}

func NewManager() *manager {
	return &manager{}
}

func (m *manager) Create(ctx context.Context, p *api.CreateAdPayload) (*api.CreateAdResult, error) {
	if p.Conditions.AgeEnd != nil && p.Conditions.AgeStart != nil {
		if *p.Conditions.AgeEnd < *p.Conditions.AgeStart {
			return nil, fmt.Errorf("`age_end` must be greater than/equal to `age_start`")
		}
	}

	return nil, nil
}
