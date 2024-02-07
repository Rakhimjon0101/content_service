package rest

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) endpoints() {
	url := ginSwagger.URL("/swagger/doc.json")
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// back
	back := s.router.Group("/")
	back.GET("ping", s.ping)
}