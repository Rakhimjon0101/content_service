package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"projects/content_service/internal/config"
	"projects/content_service/internal/domains"
	"projects/content_service/pkg/logger"
)

type blog interface {
	CreateBlog(ctx context.Context, b *domains.Blog) error
	GetBlogByID(ctx context.Context, id string) (domains.Blog, error)
	GetBlogs(ctx context.Context) ([]domains.Blog, error)
}

type news interface {
	CreateNews(ctx context.Context) error
}

type Server struct {
	cfg        config.Config
	log        logger.Logger
	router     *gin.Engine
	httpServer *http.Server

	blog blog
	news news
}

func New(cfg config.Config, log logger.Logger, usc blog) *Server {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	s := &Server{
		cfg:    cfg,
		log:    log,
		router: r,
		blog:   usc,
	}

	s.endpoints()

	s.httpServer = &http.Server{
		Addr:              cfg.HTTPPort,
		Handler:           r,
		ReadHeaderTimeout: time.Second * 10, //nolint:gomnd
	}
	s.log.Info(fmt.Sprintf("HTTP server is initialized on port %s", cfg.HTTPPort))
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
