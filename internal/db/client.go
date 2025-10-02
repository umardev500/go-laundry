package db

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/config"

	_ "github.com/lib/pq"
)

type Client struct {
	Client *ent.Client
}

type contextKeyTx struct{}

func NewEntClient(cfg *config.Config) *Client {
	client, err := ent.Open(cfg.Database.Driver, cfg.Database.Dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema resources")
	}

	return &Client{
		Client: client,
	}
}

// GetConn returns the current client, or a transactional client if present in ctx
func (c *Client) GetConn(ctx context.Context) *ent.Client {
	if tx, ok := ctx.Value(contextKeyTx{}).(*ent.Tx); ok {
		return tx.Client()
	}
	return c.Client
}

// WithTransaction runs fn inside a transaction and injects it into ctx
func (e *Client) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// Skip reassign if ctx alredy has tx
	if _, ok := ctx.Value(contextKeyTx{}).(*ent.Tx); ok {
		return fn(ctx)
	}

	tx, err := e.Client.Tx(ctx)
	if err != nil {
		return err
	}

	ctxTx := context.WithValue(ctx, contextKeyTx{}, tx)
	if err := fn(ctxTx); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Error().Err(err).Msg("Failed to rollback transaction")
		}
		return err
	}

	return tx.Commit()
}
