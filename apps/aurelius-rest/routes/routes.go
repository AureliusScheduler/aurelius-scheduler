package routes

import (
	"aurelius/apps/aurelius-rest/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func registerRoutes(
	jobController controllers.JobController,
	agentController controllers.AgentController,
	handler Handler,
) {

	handler.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jobRoutes := handler.Gin.Group("/jobs")
	{
		jobRoutes.GET("/", jobController.Get)
		jobRoutes.PATCH("/:id", jobController.Patch)
	}

	agentRoutes := handler.Gin.Group("/agents")
	{
		agentRoutes.GET("/", agentController.Get)
	}

	handler.Gin.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
	})

}
