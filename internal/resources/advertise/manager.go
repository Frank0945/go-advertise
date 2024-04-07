package advertise

import (
	"context"
	"fmt"

	"github.com/Frank0945/go-advertise/internal/entities/advertise"
)

type Manager interface {
	CreateAd(context.Context, *advertise.AdOverview) (string, error)
	ListAds(context.Context, *advertise.AdQuery) ([]*advertise.Ad, error)
}

type manager struct {
	adMapper advertise.AdMapper
}

func NewManager(mapper advertise.AdMapper) *manager {
	return &manager{
		adMapper: mapper,
	}
}

func (m *manager) CreateAd(ctx context.Context, ad *advertise.AdOverview) (string, error) {
	if ad.AgeEnd.Valid && ad.AgeStart.Valid {
		// Check if `age_end` is greater than/equal to `age_start`
		if ad.AgeEnd.Int64 < ad.AgeStart.Int64 {
			return "", fmt.Errorf("`age_end` must be greater than/equal to `age_start`")
		}
	}

	id, err := m.adMapper.Create(ctx, ad)
	if err != nil {
		// Maybe log the error here
		return "", err
	}

	return id, err
}

func (m *manager) ListAds(ctx context.Context, query *advertise.AdQuery) ([]*advertise.Ad, error) {
	ads, err := m.adMapper.List(ctx, query)
	if err != nil {
		// Maybe log the error here
		return nil, err
	}

	return ads, nil
}
