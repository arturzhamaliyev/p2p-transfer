package services

import (
	"context"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/internal/repository"
	"github.com/knadh/koanf/v2"
)

type user struct {
	cfg  *koanf.Koanf
	repo repository.Repo
}

func NewUserService(cfg *koanf.Koanf, repo repository.Repo) *user {
	return &user{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *user) CreateUser(ctx context.Context, userData dto.UserCreate) error {
	err := u.repo.User.CreateUser(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}
