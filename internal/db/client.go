package db

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/config"

	_ "github.com/lib/pq"
)

type Client struct {
	client *ent.Client
}

type contextKey struct{}

func NewClient(config *config.Config) *Client {
	client, err := ent.Open("postgres", config.Database.Url)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to run auto migration")
	}

	return &Client{
		client: client,
	}
}

func (c *Client) Client() *ent.Client {
	return c.client
}

func (c *Client) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// Ensure existing transaction is reused
	if _, ok := ctx.Value(contextKey{}).(*ent.Tx); ok {
		return fn(ctx)
	}

	tx, err := c.client.Tx(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, contextKey{}, tx)
	if err := fn(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	return tx.Commit()
}

func (c *Client) GetConn(ctx context.Context) *ent.Client {
	if tx, ok := ctx.Value(contextKey{}).(*ent.Tx); ok {
		return tx.Client()
	}

	return c.client
}

func (c *Client) Close() error {
	return c.client.Close()
}
