package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/knadh/koanf/v2"
)

var ctx = context.Background()

func Connect(cfg *koanf.Koanf) (*pgx.Conn, error) {
	//  "postgres://postgres:admin@localhost:5432/postgres"
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.String("db.driver"),
		cfg.String("db.user"),
		cfg.String("db.password"),
		cfg.String("db.host"),
		cfg.Int("db.port"),
		cfg.String("db.name"))

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
