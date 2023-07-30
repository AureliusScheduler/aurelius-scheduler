package crontab

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"aurelius/libs/aurelius_service"
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CrontabRunner struct {
	dbContext          *db_context.DbContext
	jobInstanceService *aurelius_service.JobInstanceService
	nowProvider        aurelius_service.NowProvider
}

func NewCrontabRunner(
	dbContext *db_context.DbContext,
	jobInstanceService *aurelius_service.JobInstanceService,
	nowProvider aurelius_service.NowProvider,
) *CrontabRunner {
	return &CrontabRunner{
		dbContext:          dbContext,
		jobInstanceService: jobInstanceService,
		nowProvider:        nowProvider,
	}
}

func (r *CrontabRunner) Run() {
	//now := time.Now()
	//nextMinute := now.Truncate(time.Minute).Add(time.Minute)
	//duration := nextMinute.Sub(now)
	//
	//time.Sleep(duration)

	ticker := time.NewTicker(time.Second * 10)

	go func() {
		for range ticker.C {
			fmt.Println("Tick")
			err := r.tick()

			if err != nil {
				fmt.Println(err)
			}

		}
	}()
}

func (r *CrontabRunner) tick() error {
	nextJobs, err := r.jobInstanceService.FindAndLockNext()

	if err != nil {
		return err
	}

	fmt.Println("Found", len(nextJobs), "jobs")

	if len(nextJobs) == 0 {
		return nil
	}

	// create next instance
	go func() {
		nextInstances := make([]*entity.JobInstance, 0)

		for _, job := range nextJobs {
			expr := cronexpr.MustParse(job.Cron)
			nextRun := expr.Next(r.nowProvider.Now())

			nextInstance := &entity.JobInstance{
				JobID:      job.JobID,
				JobVersion: job.JobVersion,
				Cron:       job.Cron,
				RunAt:      nextRun,
				Status:     entity.Pending,
			}

			nextInstances = append(nextInstances, nextInstance)
		}

		err := r.dbContext.GormDb.Save(nextInstances)

		if err != nil {
			fmt.Println(err)
		}
	}()

	// run jobs
	go func() {
		for _, job := range nextJobs {
			fmt.Println("Running job", job.ID)
		}
	}()

	return nil
}
