package advertise

import (
	"context"
	"database/sql"
	"time"
)

type Ad struct {
	Title string    `db:"title"`
	EndAt time.Time `db:"end_at"`
}

type AdOverview struct {
	Title    string         `db:"title"`
	StartAt  string         `db:"start_at"`
	EndAt    string         `db:"end_at"`
	AgeStart sql.NullInt64  `db:"age_start"`
	AgeEnd   sql.NullInt64  `db:"age_end"`
	Gender   sql.NullString `db:"gender"`
	Country  sql.NullString `db:"country"`
	Platform sql.NullString `db:"platform"`
}

type AdQuery struct {
	Offset   int            `db:"offset"`
	Limit    int            `db:"limit"`
	AgeStart sql.NullInt64  `db:"age_start"`
	AgeEnd   sql.NullInt64  `db:"age_end"`
	Gender   sql.NullString `db:"gender"`
	Country  sql.NullString `db:"country"`
	Platform sql.NullString `db:"platform"`
}

type AdMapper interface {
	Create(context.Context, *AdOverview) (string, error)
	List(context.Context, *AdQuery) ([]*Ad, error)
}
