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
	tx, err := a.db.BeginTxx(ctx, nil)
	defer tx.Rollback()

	if err != nil {
		return "", err
	}

	if err := a.increaseDailyAdCount(tx); err != nil {
		return "", err
	}

	insertedID, err := a.createAd(tx, ad)
	if err != nil {
		return "", err
	}

	return insertedID, tx.Commit()
}

func (a *adMapper) createAd(tx *sqlx.Tx, ad *AdOverview) (string, error) {
	query := `
		INSERT INTO advertisement (title, start_at, end_at, age_start, age_end, gender, country, platform)
		VALUES (:title, :start_at, :end_at, :age_start, :age_end, :gender, :country, :platform)
		RETURNING id
	`

	rows, err := tx.NamedQuery(query, ad)
	if err != nil {
		return "", err
	}

	var insertedID string
	if rows.Next() {
		if err := rows.Scan(&insertedID); err != nil {
			return "", err
		}
	}

	return insertedID, err
}

func (a *adMapper) increaseDailyAdCount(tx *sqlx.Tx) error {
	query := `
		INSERT INTO daily_ad_count (date, count)
		VALUES (CURRENT_DATE, 1)
		ON CONFLICT (date)
		DO UPDATE SET count = daily_ad_count.count + 1
	`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (a *adMapper) List(ctx context.Context, q *AdQuery) ([]*Ad, error) {
	whereQuery := `
		WHERE start_at < CURRENT_TIMESTAMP AND end_at > CURRENT_TIMESTAMP
	`
	if q.AgeStart.Valid {
		whereQuery += fmt.Sprintf(" AND age_start >= %d", q.AgeStart.Int64)
	}
	if q.AgeEnd.Valid {
		whereQuery += fmt.Sprintf(" AND age_end <= %d", q.AgeEnd.Int64)
	}
	if q.Gender.Valid {
		whereQuery += fmt.Sprintf(" AND gender @> '%s'", q.Gender.String)
	}
	if q.Country.Valid {
		whereQuery += fmt.Sprintf(" AND country @> '%s'", q.Country.String)
	}
	if q.Platform.Valid {
		whereQuery += fmt.Sprintf(" AND platform @> '%s'", q.Platform.String)
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
		return nil, err
	}

	return ads, nil
}
