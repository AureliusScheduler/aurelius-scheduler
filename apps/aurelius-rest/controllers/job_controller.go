package controllers

import (
	storage "aurelius/libs/aurelius-storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobController struct {
	Get func(c *gin.Context)

	Post func(c *gin.Context)
}

func NewJobController(repository storage.JobRepository) JobController {
	return JobController{
		Get:  getScheduledJobs(repository),
		Post: createScheduledJob(),
	}
}

type scheduledJob struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cron string `json:"cron"`
}

var scheduledJobs = []scheduledJob{
	{ID: 1, Name: "hosp_daily_job", Cron: "0 0 0 * * *"},
	{ID: 2, Name: "remove_stale_consultants", Cron: "0 1 0 * * *"},
	{ID: 3, Name: "recalculate_invoices", Cron: "0 2 0 * * *"},
}

// GetScheduledJobs godoc
// @Summary Get all scheduled jobs
// @Router /jobs [get]
// @Produce json
// @Tags jobs
// @Success 200 {array} scheduledJob
func getScheduledJobs(repository storage.JobRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		items, err := repository.GetAll()

		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, items)
	}
}

func createScheduledJob() func(c *gin.Context) {
	return func(c *gin.Context) {
		var job scheduledJob
		if err := c.BindJSON(&job); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		scheduledJobs = append(scheduledJobs, job)
		c.IndentedJSON(http.StatusCreated, job)
	}
}
