package handler

import (
	"todoApp/pkg/logger"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, log logger.Logger) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
