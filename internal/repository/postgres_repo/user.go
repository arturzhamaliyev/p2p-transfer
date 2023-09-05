package postgres_repo

import (
	"context"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/knadh/koanf/v2"
)

type user struct {
	cfg *koanf.Koanf
	db  *pgx.Conn
}

func NewUserRepo(cfg *koanf.Koanf, db *pgx.Conn) *user {
	return &user{
		cfg: cfg,
		db:  db,
	}
}

func (u *user) CreateUser(ctx context.Context, userData dto.UserCreate) error {
	_, err := u.db.Exec(ctx, `INSERT INTO users(amount, currency) VALUES($1, $2)`, userData.Amount, userData.Currency)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) GetCount(ctx context.Context, accountID int64) (int8, error) {
	var count int8
	err := u.db.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE id = $1`, accountID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
