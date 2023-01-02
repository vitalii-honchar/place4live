package config

import (
	"database/sql"
	"place4live/internal/lib/postgres"
	"place4live/internal/lib/web"
)

type AppContext struct {
	Db        *sql.DB
	ApiEngine *web.ApiEngine
	Config
}

func NewAppContext(cfg *Config) (*AppContext, error) {
	db, err := postgres.OpenConnection(cfg.DbConnectionStr, cfg.DbMigrationsFolder)
	if err != nil {
		return nil, err
	}
	return &AppContext{Db: db, ApiEngine: web.NewApiEngine(), Config: *cfg}, nil
}
