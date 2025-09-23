package database

import (
	"go-api-arch-clean-template/pkg"
)

type Config struct {
	Host     string
	Database string
	Port     string
	Driver   string
	User     string
	Password string
}

// MySQL用の設定を作成
func NewConfigMySQL() *Config {
	return &Config{
		Host:     pkg.GetEnvDefault("DB_HOST", "localhost"),
		Database: pkg.GetEnvDefault("DB_NAME", "api_database"),
		Port:     pkg.GetEnvDefault("DB_PORT", "3306"),
		Driver:   pkg.GetEnvDefault("DB_DRIVER", "mysql"),
		User:     pkg.GetEnvDefault("DB_USER", "app"),
		Password: pkg.GetEnvDefault("DB_PASSWORD", "password"),
	}
}

// SQLite用の設定を作成
func NewConfigSQLite() *Config {
	return &Config{
		Database: pkg.GetEnvDefault("DB_NAME", "api_database.sqlite"),
	}
}
