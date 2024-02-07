package bootstrap

import (
	"context"

	"projects/content_service/internal/config"
	"projects/content_service/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func initDB(ctx context.Context, l logger.Logger, cfg *config.DBConfig) (*pgxpool.Pool, func()) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.PostgresURL())
	if err != nil {
		l.Fatal(err.Error())
	}

	pgxCfg.MaxConnIdleTime = cfg.ConnectionMaxLifetime
	pgxCfg.MaxConns = cfg.MaxIdleConnections
	pgxCfg.MaxConnLifetime = cfg.ConnectionMaxLifetime

	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		l.Fatal("sqlx.Connect", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		l.Fatal("Failed to ping Database", zap.Error(err))
	}

	l.Info("Database connection established")

	return pool, pool.Close
}
