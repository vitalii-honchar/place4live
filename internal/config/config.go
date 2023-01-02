package config

import (
	"fmt"
	"os"
)

const envDbConnectionStr = "DB_CONNECTION_URL"
const envDbMigrationsFolder = "DB_MIGRATIONS_FOLDER"
const envApiSecret = "API_SECRET"
const defaultMigrationsFolder = "migrations"

type Config struct {
	DbConnectionStr    string
	DbMigrationsFolder string
	ApiSecret          string
}

func NewConfig() (*Config, error) {
	connStr := os.Getenv(envDbConnectionStr)
	if connStr == "" {
		return nil, MissedConfigError(fmt.Sprintf("missed config: %s", envDbConnectionStr))
	}
	migrations := os.Getenv(envDbMigrationsFolder)
	if migrations == "" {
		migrations = defaultMigrationsFolder
	}
	apiSecret := os.Getenv(envApiSecret)
	if apiSecret == "" {
		return nil, MissedConfigError(fmt.Sprintf("missed config: %s", envApiSecret))
	}

	return &Config{DbConnectionStr: connStr, DbMigrationsFolder: migrations, ApiSecret: apiSecret}, nil
}

type MissedConfigError string

func (e MissedConfigError) Error() string {
	return string(e)
}
