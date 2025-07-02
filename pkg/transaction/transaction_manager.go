package transaction

import (
	"context"

	"github.com/umardev500/go-laundry/internal/ent"
)

type txKey struct{}

type TransactionManager struct {
	Client *ent.Client
}

func NewTransactionManager(client *ent.Client) *TransactionManager {
	return &TransactionManager{
		Client: client,
	}
}

func (tm *TransactionManager) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.Client.Tx(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (tm *TransactionManager) GetClient(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(txKey{}).(*ent.Tx)
	if ok {
		return tx.Client()
	}
	return tm.Client
}
