package rest

import "github.com/gin-gonic/gin"

func (s *Server) ping(c *gin.Context) {
	s.handleSuccessResponse(c, "Pong!")
}
