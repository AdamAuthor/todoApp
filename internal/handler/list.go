package handler

import (
	"net/http"
	"strconv"
	"todoApp/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	notFound    = "user id not found"
	invalidID   = "invalid id param"
	invalidBody = "invalid input body"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFound)
		return
	}

	var list models.List
	if err := c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidBody)
		return
	}

	id, err := h.service.TodoList.Create(userID, list)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		newErrorResponse(c, http.StatusInternalServerError, notFound)
		return
	}
	lists, err := h.service.TodoList.GetAll(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFound)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidID)
		return
	}

	list, err := h.service.TodoList.GetByID(userID, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFound)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidID)
		return
	}

	var list models.UpdateListInput
	if err = c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidBody)
		return
	}

	if err = h.service.TodoList.Update(userID, id, list); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, invalidID)
		return
	}

	if err = h.service.TodoList.Delete(userID, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
