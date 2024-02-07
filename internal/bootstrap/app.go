package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"projects/content_service/internal/config"
	"projects/content_service/internal/dbstore"
	"projects/content_service/internal/gateways/rest"

	"projects/content_service/pkg/logger"
)

type Application struct {
	Logger   logger.Logger
	Config   config.Config
	RestAPI  *rest.Server
	Teardown []func()
}

func New(cfg config.Config) *Application {
	teardown := make([]func(), 0)
	app := Application{Config: cfg}

	appLogger := logger.NewApiLogger(&cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.AppVersion, cfg.Logger.LogLevel, cfg.Environment)
	app.Logger = appLogger

	db, c := initDB(context.TODO(), app.Logger, cfg.Postgres)
	teardown = append(teardown, c)

	storage := dbstore.New(app.Logger, db)
	useCases := buildUseCases(app.Config, app.Logger, storage)

	app.RestAPI = rest.New(cfg, useCases.Blog)

	app.Teardown = teardown
	return &app
}

// Run application and all its sub core
func (app *Application) Run(ctx context.Context) {
	go func() {
		if err := app.RestAPI.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed To Run REST Server: %s\n", err.Error())
		}
		app.Logger.Info("REST Server started at port " + app.Config.HTTPPort)
	}()
	app.Teardown = append(app.Teardown, func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := app.RestAPI.Shutdown(ctx); err != nil {
			app.Logger.Error(fmt.Sprintf("REST Server Graceful Shutdown Failed: %s\n", err))
		}
		app.Logger.Info("REST Server Graceful Shutdown")
	})

	<-ctx.Done()
	for i := 0; i < len(app.Teardown); i++ {
		app.Teardown[len(app.Teardown)-1-i]()
	}
	app.Logger.Info("Program Gracefully Shut Down")
}
