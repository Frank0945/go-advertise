package advertise_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Frank0945/go-advertise/internal/entities/advertise"
	"github.com/jmoiron/sqlx"
)

func TestAdMapper_Create(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create a sqlx.DB object using the mock database
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new instance of adMapper with the mock database
	adMapper := advertise.NewAdMapper(sqlxDB)

	// Create a context
	ctx := context.Background()

	// Mock the expectations for the transaction
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO daily_ad_count`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(`INSERT INTO advertisement`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	// Create a sample AdOverview
	ad := &advertise.AdOverview{
		Title:    "Test Ad",
		StartAt:  time.Now().Format(time.RFC3339),
		EndAt:    time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		AgeStart: sql.NullInt64{Int64: 18, Valid: true},
		AgeEnd:   sql.NullInt64{Int64: 60, Valid: true},
		Gender:   sql.NullString{String: "M", Valid: true},
		Country:  sql.NullString{String: "TW", Valid: true},
		Platform: sql.NullString{String: "web", Valid: true},
	}

	// Call the Create method
	_, err = adMapper.Create(ctx, ad)
	if err != nil {
		t.Fatalf("failed to create ad: %v", err)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAdMapper_List(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create a sqlx.DB object using the mock database
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new instance of adMapper with the mock database
	adMapper := advertise.NewAdMapper(sqlxDB)

	// Create a context
	ctx := context.Background()

	// Sample query parameters
	q := &advertise.AdQuery{
		AgeStart: sql.NullInt64{Int64: 18, Valid: true},
		AgeEnd:   sql.NullInt64{Int64: 60, Valid: true},
		Gender:   sql.NullString{String: "M", Valid: true},
		Country:  sql.NullString{String: "TW", Valid: true},
		Platform: sql.NullString{String: "web", Valid: true},
		Limit:    10,
		Offset:   0,
	}

	// Mock the expectations for the query
	rows := sqlmock.NewRows([]string{"title", "end_at"}).AddRow("Test Ad", time.Now().Add(24*time.Hour))
	mock.ExpectQuery(`SELECT title, end_at FROM advertisement`).WillReturnRows(rows)

	// Call the List method
	_, err = adMapper.List(ctx, q)
	if err != nil {
		t.Fatalf("failed to list ads: %v", err)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
