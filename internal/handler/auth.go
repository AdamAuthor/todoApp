package handler

import (
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
