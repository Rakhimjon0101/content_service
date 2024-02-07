package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"projects/content_service/api/docs"
	"projects/content_service/internal/bootstrap"
	"projects/content_service/internal/config"

	"github.com/sethvargo/go-envconfig"
)

// @title Go Content Service
// @version 1.0
// @description Golang Content Service
// @contact.name Rakhimjon Shokirov
// @contact.url https://github.com/Rakhimjon0101
// @contact.email raximvarresult@gmail.com
// @BasePath /api/v1
func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	docs.SwaggerInfo.Title = config.AppName
	docs.SwaggerInfo.Description = "Retrieves content service, creates new application based on it"

	var cfg config.Config
	if err := envconfig.ProcessWith(context.TODO(), &cfg, envconfig.OsLookuper()); err != nil {
		log.Fatal(err)
	}

	app := bootstrap.New(cfg)
	app.Logger.Info("[main] Application bootstrapped successfully")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		OSCall := <-quit
		app.Logger.Info(fmt.Sprintf("\nSystem Call: %+v", OSCall))
		cancel()
	}()

	app.Run(ctx)
}
