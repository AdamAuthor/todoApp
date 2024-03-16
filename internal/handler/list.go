package handler

import (
	"net/http"
	"strconv"
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	notFound  = "user id not found"
	invalidID = "invalid id param"
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

type getAllListsResponse struct {
	Data []models.List `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}
	lists, err := h.service.TodoList.GetAll(userID)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: invalidID})
		return
	}

	list, err := h.service.TodoList.GetByID(userID, id)
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
