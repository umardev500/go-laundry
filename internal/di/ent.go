package di

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/migrate"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to postgres")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func NewEntClient(ctx context.Context, config *config.AppConfig) *ent.Client {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.SSLMode,
	)

	client := Open(dsn)
	if err := client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	log.Info().
		Str("dsn", dsn).
		Msg("database connected successfully")

	return client
}

func ProvideEntClient(ctx context.Context, cfg *config.AppConfig) *ent.Client {
	return NewEntClient(ctx, cfg)
}
