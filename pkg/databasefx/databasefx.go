package databasefx

import (
	"fmt"

	"github.com/Frank0945/go-advertise/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

// Module provides a *sqlx.DB with a connection to a PostgreSQL database.
var Module = fx.Provide(New)

type Params struct {
	fx.In

	Cfg *config.Config
}

func New(p Params) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		p.Cfg.DBCfg.User,
		p.Cfg.DBCfg.Password,
		p.Cfg.DBCfg.Database,
		p.Cfg.DBCfg.Host,
		p.Cfg.DBCfg.Port,
	))
	if err != nil {
		return nil, err
	}

	return db, nil
}
