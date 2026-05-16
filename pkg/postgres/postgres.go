package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"task_queue/pkg/env"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSL      string
}

func NewConnection() (*sql.DB, error) {
	config := DBConfig{
		Host:     env.Get("POSTGRES_HOST", "localhost"),
		Port:     env.Get("POSTGRES_PORT", "5432"),
		User:     env.Get("POSTGRES_USER", "root"),
		Password: env.Get("POSTGRES_PASSWORD", "root"),
		DBName:   env.Get("POSTGRES_DB", "main"),
		SSL:      env.Get("POSTGRES_SSL_MODE", "disable"),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSL,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB: %w", err)
	}
	slog.Info("Successfully connected to DB")

	return db, nil
}
