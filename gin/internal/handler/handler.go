package handler

import (
	"gin/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(srv *services.Service) *Handler {
	return &Handler{
		services: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	room := router.Group("/room")
	{
		room.GET("/:id", h.GetRoomById)
		room.POST("/", h.CreateRoom)
	}

	return router
}
