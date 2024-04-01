package advertise

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type adMapper struct {
	db *sqlx.DB
}

func NewAdMapper(db *sqlx.DB) *adMapper {
	return &adMapper{
		db: db,
	}
}

func (a *adMapper) Create(ctx context.Context, ad *AdOverview) (string, error) {
	query := `
		INSERT INTO advertisement (title, start_at, end_at, age_start, age_end, gender, country, platform)
		VALUES (:title, :start_at, :end_at, :age_start, :age_end, :gender, :country, :platform)
		RETURNING id
	`
	rows, err := a.db.NamedQueryContext(ctx, query, ad)
	if err != nil {
		return "", err
	}

	var insertedID string
	if rows.Next() {
		if err := rows.Scan(&insertedID); err != nil {
			return "", fmt.Errorf("%s", err)
		}
	}

	return insertedID, err
}

func (a *adMapper) List(ctx context.Context, q *AdQuery) ([]*Ad, error) {
	whereQuery := `
		WHERE start_at < CURRENT_TIMESTAMP AND end_at > CURRENT_TIMESTAMP
	`
	if q.AgeStart != nil {
		whereQuery += fmt.Sprintf(" AND age_start >= %d", *q.AgeStart)
	}
	if q.AgeEnd != nil {
		whereQuery += fmt.Sprintf(" AND age_end <= %d", *q.AgeEnd)
	}
	if q.Gender != nil {
		whereQuery += fmt.Sprintf(" AND gender @> '%s'", *q.Gender)
	}
	if q.Country != nil {
		whereQuery += fmt.Sprintf(" AND country @> '%s'", *q.Country)
	}
	if q.Platform != nil {
		whereQuery += fmt.Sprintf(" AND platform @> '%s'", *q.Platform)
	}

	query := fmt.Sprintf(`
		SELECT title, end_at
		FROM advertisement
		%s
		ORDER BY end_at ASC
		LIMIT %d OFFSET %d
	`, whereQuery, q.Limit, q.Offset)

	ads := []*Ad{}
	err := a.db.SelectContext(ctx, &ads, query)
	if err != nil {
		return nil, fmt.Errorf(query)
	}

	return ads, nil
}
