package config

import (
	"fmt"
	"os"
)

const envDbConnectionStr = "DB_CONNECTION_URL"
const envDbMigrationsFolder = "DB_MIGRATIONS_FOLDER"

const defaultMigrationsFolder = "migrations"

type Config struct {
	DbConnectionStr    string
	DbMigrationsFolder string
}

func NewConfig() (*Config, error) {
	connStr := os.Getenv(envDbConnectionStr)
	if connStr == "" {
		return nil, MissedConfigError(fmt.Sprintf("Missed config: %s", envDbConnectionStr))
	}
	migrations := os.Getenv(envDbMigrationsFolder)
	if migrations == "" {
		migrations = defaultMigrationsFolder
	}
	return &Config{DbConnectionStr: connStr, DbMigrationsFolder: migrations}, nil
}

type MissedConfigError string

func (e MissedConfigError) Error() string {
	return string(e)
}
