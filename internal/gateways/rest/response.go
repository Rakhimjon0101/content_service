package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = "Success"
	Failure = "Failure"

	InternalError = -50
)

type R struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note,omitempty"`
	Data      interface{} `json:"data"`
}

func (s *Server) handleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, R{
		Status:    Success,
		ErrorCode: 0,
		ErrorNote: "",
		Data:      data,
	})
}

func (s *Server) handleErrorResponse(c *gin.Context, httpCode, errCode int, err error) {
	c.JSON(httpCode, R{
		Status:    Failure,
		ErrorCode: errCode,
		ErrorNote: err.Error(),
		Data:      nil,
	})
}
