package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/internal/repository"
	"github.com/knadh/koanf/v2"
)

var accountExistenseErr = "account doesn't exist: "

type transfer struct {
	cfg  *koanf.Koanf
	repo repository.Repo
}

func NewTransferService(cfg *koanf.Koanf, repo repository.Repo) *transfer {
	return &transfer{
		cfg:  cfg,
		repo: repo,
	}
}

func (t *transfer) MakeTransfer(ctx context.Context, transfer dto.TransferRequest) error {
	count, err := t.repo.User.GetCount(ctx, transfer.FromAccountID)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(accountExistenseErr + strconv.Itoa(int(transfer.FromAccountID)))
	}

	count, err = t.repo.User.GetCount(ctx, transfer.ToAccountID)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(accountExistenseErr + strconv.Itoa(int(transfer.FromAccountID)))
	}

	transaction, err := t.repo.Transfer.Withdraw(ctx, transfer.FromAccountID, transfer.Amount)
	if err != nil {
		transaction.Rollback(ctx)
		return err
	}

	err = t.repo.Transfer.PutMoney(ctx, transaction, transfer.ToAccountID, transfer.Amount)
	if err != nil {
		transaction.Rollback(ctx)
		return err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
