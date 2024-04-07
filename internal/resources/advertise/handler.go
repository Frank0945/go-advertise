package advertise

import (
	"context"
	"fmt"
	"time"

	api "github.com/Frank0945/go-advertise/api/gen/manager"
	"github.com/Frank0945/go-advertise/internal/entities/advertise"
	"github.com/Frank0945/go-advertise/internal/utils/slice"
	"github.com/Frank0945/go-advertise/internal/utils/sqlconv"
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
	// Convert the conditions to required format for the entity
	if p.Conditions != nil {
		adOverviewEntity.AgeStart = sqlconv.ConvToNullInt64(p.Conditions.AgeStart)
		adOverviewEntity.AgeEnd = sqlconv.ConvToNullInt64(p.Conditions.AgeEnd)

		gender := slice.ConvSlicToStr(slice.RmDupStrSlic(p.Conditions.Gender))
		adOverviewEntity.Gender = sqlconv.ConvToNullStr(&gender)

		country := slice.ConvSlicToStr(slice.RmDupStrSlic(p.Conditions.Country))
		adOverviewEntity.Country = sqlconv.ConvToNullStr(&country)

		platform := slice.ConvSlicToStr(slice.RmDupStrSlic(p.Conditions.Platform))
		adOverviewEntity.Platform = sqlconv.ConvToNullStr(&platform)
	}

	// Call the manager to create the AD
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

	gender := slice.ConvSlicToStr(slice.RmDupStrSlic(q.Gender))
	country := slice.ConvSlicToStr(slice.RmDupStrSlic(q.Country))
	platform := slice.ConvSlicToStr(slice.RmDupStrSlic(q.Platform))

	queryEntity := &advertise.AdQuery{
		Offset:   q.Offset,
		Limit:    q.Limit,
		AgeStart: sqlconv.ConvToNullInt64(q.AgeStart),
		AgeEnd:   sqlconv.ConvToNullInt64(q.AgeEnd),
		Gender:   sqlconv.ConvToNullStr(&gender),
		Country:  sqlconv.ConvToNullStr(&country),
		Platform: sqlconv.ConvToNullStr(&platform),
	}

	ads, err := h.manager.ListAds(ctx, queryEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to list ADs: %w", err)
	}

	// Convert the entity to the required format
	for _, ad := range ads {
		res = append(res, &api.Ad{
			Title: ad.Title,
			EndAt: ad.EndAt.Format(time.RFC3339),
		})
	}

	return res, nil
}
