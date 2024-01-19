package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/sethvargo/go-envconfig"
	"honnef.co/go/tools/config"
)

// @title Go Contect Service
// @version 1.0
// @description Golang Contect Service
// @contact.name Rakhimjon Shokirov
// @contact.url https://github.com/Rakhimjon0101
// @contact.email raximvarresult@gmail.com
// @BasePath /api/v1
func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		cancel()
	}()

	var cfg config.Config
	// if err := envconfig.ProcessWith(context.TODO(), &cfg, envconfig.OsLookuper()); err != nil {
	// 	log.Fatal(err)
	// }

	// app := bootstrap.New(cfg)
	defer app.Shutdown(ctx)

	app.Run(ctx)
}
