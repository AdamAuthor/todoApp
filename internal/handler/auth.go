package handler

import (
	"net/http"
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), h.log)
		return
	}
	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error(), h.log)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var signIn signInInput

	if err := c.BindJSON(&signIn); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	token, err := h.service.Authorization.GenerateToken(signIn.Username, signIn.Password)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
