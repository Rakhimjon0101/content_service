package bootstrap

import (
	"context"

	"projects/content_service/config"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"projects/content_service/pkg/logger"
)

type App struct {
	restAPI  *rest.Server
	teardown []func()
}

func New(cfg config.Config) *App {
	teardown := make([]func(), 0)
	app := new(App)

	gLog := logger.New(cfg.LogLevel, cfg.App, zap.AddCallerSkip(1))
	teardown = append(teardown, func() { _ = logger.Cleanup() })
	gLog.Info("Logger init")

	app.log = logger.FromCtx(context.Background(), "bootstrap.New")

	db, repo, c := initRepository(cfg, app.log)
	teardown = append(teardown, c)

	app.outbox = initOutbox(cfg, app.log, db, amqp)
	cases := buildUseCases(app.log, repo, app.outbox)

	app.restAPI = rest.New(cfg, app.outbox)

	app.teardown = teardown

	return app
}

func initRepository(cfg config.Config, l logger.Logger) (*sqlx.DB, repository.Container, func()) {
	client, err := sqlx.Connect("postgres", cfg.PostgresURL())
	if err != nil {
		panic(err)
	}

	r := repository.New(client)
	l.Debug("repository connected")

	return client, r, func() {
		if err := client.Close(); err != nil {
			l.Error("db close error", zap.Error(err))
		}
	}
}
