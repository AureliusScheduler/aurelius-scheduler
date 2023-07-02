package routes

import (
	"aurelius/apps/aurelius-rest/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Invoke(registerRoutes),
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() Handler {
	return Handler{
		Gin: gin.Default(),
	}
}

func registerRoutes(controller controllers.JobController, handler Handler) {

	jobRoutes := handler.Gin.Group("/jobs")
	{
		jobRoutes.GET("/", controller.Get)
		jobRoutes.POST("/", controller.Post)
	}

	handler.Gin.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
	})

}
