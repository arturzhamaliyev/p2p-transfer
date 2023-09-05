package repository

import (
	"context"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/internal/repository/postgres_repo"
	"github.com/jackc/pgx/v5"
	"github.com/knadh/koanf/v2"
)

type Transfer interface {
	Withdraw(ctx context.Context, accountID, amount int64) (pgx.Tx, error)
	PutMoney(ctx context.Context, trx pgx.Tx, accountID, amount int64) error
}

type User interface {
	CreateUser(ctx context.Context, userData dto.UserCreate) error
	GetCount(ctx context.Context, accountID int64) (int8, error)
}

type Repo struct {
	Transfer Transfer
	User     User
}

func NewRepository(cfg *koanf.Koanf, db *pgx.Conn) Repo {
	return Repo{
		Transfer: postgres_repo.NewTransferRepo(cfg, db),
		User:     postgres_repo.NewUserRepo(cfg, db),
	}
}
