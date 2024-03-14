package handler

import (
	"todoApp/internal/service"
	"todoApp/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	log     logger.Logger
	service *service.Service
}

func NewHandler(log logger.Logger, srv *service.Service) *Handler {
	return &Handler{
		log: log,
		service: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := router.Group("/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
