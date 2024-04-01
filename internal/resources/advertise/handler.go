package advertise

import (
	"context"
	"fmt"
	"time"

	api "github.com/Frank0945/go-advertise/api/gen/manager"
	"github.com/Frank0945/go-advertise/internal/entities/advertise"
	. "github.com/Frank0945/go-advertise/internal/utils"
)

type Handler struct {
	manager Manager
}

func NewHttpHandler(mgr Manager) *Handler {
	return &Handler{
		manager: mgr,
	}
}

func (h *Handler) CreateAd(ctx context.Context, p *api.CreateAdPayload) (*api.CreateAdResult, error) {
	adOverviewEntity := &advertise.AdOverview{
		Title:   p.Title,
		StartAt: p.StartAt,
		EndAt:   p.EndAt,
	}
	if p.Conditions != nil {
		adOverviewEntity.AgeStart = p.Conditions.AgeStart
		adOverviewEntity.AgeEnd = p.Conditions.AgeEnd

		gender := ConvSlicToStr(RmDupStrSlic(p.Conditions.Gender))
		adOverviewEntity.Gender = &gender

		country := ConvSlicToStr(RmDupStrSlic(p.Conditions.Country))
		adOverviewEntity.Country = &country

		platform := ConvSlicToStr(RmDupStrSlic(p.Conditions.Platform))
		adOverviewEntity.Platform = &platform
	}

	id, err := h.manager.CreateAd(ctx, adOverviewEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to create AD: %w", err)
	}

	return &api.CreateAdResult{
		ID: id,
	}, nil
}

func (h *Handler) ListAds(ctx context.Context, q *api.AdQuery) ([]*api.Ad, error) {
	res := []*api.Ad{}

	gender := ConvSlicToStr(RmDupStrSlic(q.Gender))
	country := ConvSlicToStr(RmDupStrSlic(q.Country))
	platform := ConvSlicToStr(RmDupStrSlic(q.Platform))

	queryEntity := &advertise.AdQuery{
		Offset:   q.Offset,
		Limit:    q.Limit,
		AgeStart: q.AgeStart,
		AgeEnd:   q.AgeEnd,
		Gender:   &gender,
		Country:  &country,
		Platform: &platform,
	}

	ads, err := h.manager.ListAds(ctx, queryEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to list ADs: %w", err)
	}

	for _, ad := range ads {
		res = append(res, &api.Ad{
			Title: ad.Title,
			EndAt: ad.EndAt.Format(time.RFC3339),
		})
	}

	return res, nil
}
