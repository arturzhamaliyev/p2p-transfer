package service

import (
	"context"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/internal/repository"
	"github.com/arturzhamaliyev/p2p-transfer/internal/service/services"
	"github.com/knadh/koanf/v2"
)

type Transfer interface {
	MakeTransfer(ctx context.Context, transfer dto.TransferRequest) error
}

type User interface {
	CreateUser(ctx context.Context, userData dto.UserCreate) error
}

type Service struct {
	Transfer Transfer
	User     User
}

func NewService(cfg *koanf.Koanf, repo repository.Repo) Service {
	return Service{
		Transfer: services.NewTransferService(cfg, repo),
		User:     services.NewUserService(cfg, repo),
	}
}
