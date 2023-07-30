package aurelius_service

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"aurelius/libs/aurelius-mysql/repository"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/gorhill/cronexpr"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type JobService struct {
	dbContext             *db_context.DbContext
	jobInstanceRepository *repository.JobInstanceRepository
	jobRepository         *repository.JobRepository
	nowProvider           NowProvider
}

func NewJobService(
	ctx *db_context.DbContext,
	jobRepository *repository.JobRepository,
	jobInstanceRepository *repository.JobInstanceRepository,
	nowProvider NowProvider,
) *JobService {
	return &JobService{
		dbContext:             ctx,
		jobRepository:         jobRepository,
		jobInstanceRepository: jobInstanceRepository,
		nowProvider:           nowProvider,
	}
}

func (s *JobService) GetAll() ([]entity.Job, error) {
	jobs := make([]entity.Job, 0)
	result := s.dbContext.GormDb.Find(&jobs)

	if result.Error != nil {
		return nil, result.Error
	}

	return jobs, nil
}

func (s *JobService) UpdateCron(id uuid.UUID, cron string) error {
	cronExpr, err := cronexpr.Parse(cron)

	if err != nil {
		return err
	}

	job, err := s.jobRepository.FindByUUID(id)

	if err != nil {
		return err
	}

	if job == nil {
		return errors.New("job not found")
	}

	jobInstance, err := s.jobInstanceRepository.FindPendingByJobId(id)

	if err != nil {
		return err
	}

	job.Cron = null.StringFrom(cron)
	nextRun := cronExpr.Next(s.nowProvider.Now())

	if jobInstance == nil {
		jobInstance = &entity.JobInstance{
			JobID:      job.ID,
			JobVersion: job.Version,
			Cron:       job.Cron.ValueOrZero(),
			RunAt:      nextRun,
		}
	} else {
		jobInstance.RunAt = nextRun
	}

	err = s.dbContext.GormDb.Transaction(func(tx *gorm.DB) error {
		tx.Save(&job)
		tx.Save(&jobInstance)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
