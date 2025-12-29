package handler

import (
	"asteriskAPI/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	authorized := router.Group("/")

	authorized.Use(ValidateToken)
	{
		authorized.GET("/callInfo", h.getCallInfo)
		authorized.POST("/originate", h.originate)
	}

	return router
}
