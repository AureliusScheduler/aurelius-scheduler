package controllers

import (
	"aurelius/apps/aurelius-rest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobController struct {
	Get  func(c *gin.Context)
	Post func(c *gin.Context)
}

func NewJobController(jobService *service.JobService) JobController {
	return JobController{
		Get:  listJobs(jobService),
		Post: createScheduledJob(),
	}
}

// ListJobs godoc
// @Summary Get all scheduled jobs
// @Router /jobs [get]
// @Produce json
// @Tags jobs
// @Success 200 {array} dto.JobDto
func listJobs(service *service.JobService) func(c *gin.Context) {
	return func(c *gin.Context) {
		items, err := service.GetAll()

		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, items)
	}
}

func createScheduledJob() func(c *gin.Context) {
	return func(c *gin.Context) {
		panic("not implemented")
		//var job scheduledJob
		//if err := c.BindJSON(&job); err != nil {
		//	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//	return
		//}
		//
		//scheduledJobs = append(scheduledJobs, job)
		//c.IndentedJSON(http.StatusCreated, job)
	}
}
