package advertise

import (
	"context"
	"time"
)

type Ad struct {
	Title string    `db:"title"`
	EndAt time.Time `db:"end_at"`
}

type AdOverview struct {
	Title    string  `db:"title"`
	StartAt  string  `db:"start_at"`
	EndAt    string  `db:"end_at"`
	AgeStart *int    `db:"age_start"`
	AgeEnd   *int    `db:"age_end"`
	Gender   *string `db:"gender"`
	Country  *string `db:"country"`
	Platform *string `db:"platform"`
}

type AdQuery struct {
	Offset   int     `db:"offset"`
	Limit    int     `db:"limit"`
	AgeStart *int    `db:"age_start"`
	AgeEnd   *int    `db:"age_end"`
	Gender   *string `db:"gender"`
	Country  *string `db:"country"`
	Platform *string `db:"platform"`
}

type AdMapper interface {
	Create(context.Context, *AdOverview) (string, error)
	List(context.Context, *AdQuery) ([]*Ad, error)
}
