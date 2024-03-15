package handler

import (
	"net/http"
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	notFound = "user id not found"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	var list models.List
	if err := c.BindJSON(&list); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := h.service.TodoList.Create(userID, list)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
