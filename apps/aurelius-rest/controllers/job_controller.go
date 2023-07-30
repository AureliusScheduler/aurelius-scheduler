package controllers

import (
	"aurelius/apps/aurelius-rest/dto"
	"aurelius/libs/aurelius_service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
)

type JobController struct {
	Get   func(c *gin.Context)
	Patch func(c *gin.Context)
}

func NewJobController(
	jobService *aurelius_service.JobService,
) JobController {
	return JobController{
		Get:   listJobs(jobService),
		Patch: updateJob(jobService),
	}
}

// ListJobs godoc
// @Summary Get all scheduled jobs
// @Router /jobs [get]
// @Produce json
// @Tags jobs
// @Success 200 {array} dto.JobDto
func listJobs(service *aurelius_service.JobService) func(c *gin.Context) {
	return func(c *gin.Context) {
		items, err := service.GetAll()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jobDtos := make([]dto.JobDto, 0)

		for _, job := range items {
			jobDtos = append(jobDtos, dto.JobDto{
				ID:      job.ID.String(),
				AgentID: job.AgentID.String(),
				Name:    job.Name,
				Version: job.Version,
				Cron:    job.Cron.ValueOrZero(),
			})
		}

		c.IndentedJSON(http.StatusOK, jobDtos)
	}
}

// UpdateJob godoc
// @Summary Update a scheduled job
// @Router /jobs/{id} [patch]
// @Produce json
// @Tags jobs
// @Param id path string true "Job ID"
// @Param input body dto.JobInput true "Job input"
// @Success 200 {string} string "Job updated successfully"
func updateJob(service *aurelius_service.JobService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		uuid, err := uuid.FromString(id)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var input dto.JobInput

		if err := c.BindJSON(&input); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.UpdateCron(uuid, input.Cron)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
	}
}
