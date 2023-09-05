package postgres_repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/knadh/koanf/v2"
)

type transfer struct {
	cfg *koanf.Koanf
	db  *pgx.Conn
}

func NewTransferRepo(cfg *koanf.Koanf, db *pgx.Conn) *transfer {
	return &transfer{
		cfg: cfg,
		db:  db,
	}
}

func (t *transfer) Withdraw(ctx context.Context, accountID, amount int64) (pgx.Tx, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	tx.Exec(ctx, `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ`)

	query := `
	UPDATE users
	SET amount = (
		SELECT amount
		FROM users
		WHERE id = $1
	) - $2
	WHERE id = $3;
	`

	_, err = tx.Exec(ctx, query, accountID, amount, accountID)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (t *transfer) PutMoney(ctx context.Context, trx pgx.Tx, accountID, amount int64) error {
	query := `
	UPDATE users
	SET amount = (
		SELECT amount
		FROM users
		WHERE id = $1
	) + $2
	WHERE id = $3;
	`

	_, err := trx.Exec(ctx, query, accountID, amount, accountID)
	if err != nil {
		return err
	}
	return nil
}
