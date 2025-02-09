package api

import (
	"container-monitoring/backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "container-monitoring/backend/docs"
)

func NewRouter(ps service.PingService) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(ps)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "If-Match"},
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		statusGroup := api.Group("/status")
		{
			statusGroup.GET("", handler.ListStatuses)
			statusGroup.POST("", handler.CreateStatus)
			statusGroup.GET("/:id", handler.GetStatus)
			statusGroup.PUT("/:id", handler.UpdateStatus)
			statusGroup.DELETE("/:id", handler.DeleteStatus)
		}
	}

	return router
}
