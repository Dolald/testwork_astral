package handler

import (
	"github.com/Dolald/testwork_astral/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	logger  *logrus.Logger
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		documents := api.Group("/documents")
		{
			documents.POST("/upload", h.createDocument)
			documents.GET("/", h.getAllDocuments)
			documents.GET("/:id", h.getDocumentById)
			documents.DELETE("/:id", h.deleteDocument)
		}
	}
	return router
}
