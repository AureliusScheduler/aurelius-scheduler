package controllers

import (
	"aurelius/apps/aurelius-rest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentController struct {
	Get func(c *gin.Context)
}

func NewAgentController(service *service.AgentService) AgentController {
	return AgentController{
		Get: getAllAgents(service),
	}
}

// GetAgents godoc
// @Summary Get all agents
// @Produce json
// @Success 200 {array} dto.AgentDto
// @Tags agents
// @Router /agents [get]
func getAllAgents(service *service.AgentService) func(c *gin.Context) {
	return func(c *gin.Context) {
		items, err := service.GetAll()

		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, items)
	}
}
