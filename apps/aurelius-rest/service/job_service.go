package service

import (
	"aurelius/apps/aurelius-rest/dto"
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"time"
)

type JobService struct {
	dbContext *db_context.DbContext
}

func NewJobService(ctx db_context.DbContext) *JobService {
	return &JobService{
		dbContext: &ctx,
	}
}

func (s *JobService) GetAll() ([]dto.JobDto, error) {
	var jobs []entity.Job
	result := s.dbContext.GormDb.Debug().Find(&jobs)

	if result.Error != nil {
		return nil, result.Error
	}

	var jobDtos []dto.JobDto

	for _, job := range jobs {
		jobDtos = append(jobDtos, dto.JobDto{
			ID:         job.ID.String(),
			AgentID:    job.AgentID.String(),
			Name:       job.Name,
			Version:    job.Version,
			StartedAt:  job.StartedAt.Format(time.RFC3339),
			FinishedAt: job.FinishedAt.Format(time.RFC3339),
		})
	}

	return jobDtos, nil
}
