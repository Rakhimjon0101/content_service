package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"projects/content_service/pkg/logger"
)

type outbox interface {
	Run(ctx context.Context) error
	Clean(ctx context.Context) error
}

type Server struct {
	cfg        config.Config
	log        logger.Logger
	router     *gin.Engine
	httpServer *http.Server

	outbox outbox
}

func New(cfg config.Config, box outbox) *Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	l := logger.FromCtx(context.Background(), "gateways.rest")

	s := &Server{
		cfg:    cfg,
		log:    l,
		router: r,
		outbox: box,
	}
	s.endpoints()

	s.httpServer = &http.Server{
		Addr:              cfg.HTTPPort,
		Handler:           r,
		ReadHeaderTimeout: time.Second * 10, //nolint:gomnd
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Run() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
