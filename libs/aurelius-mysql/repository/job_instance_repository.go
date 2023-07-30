package repository

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type JobInstanceRepository struct {
	dbContext *db_context.DbContext
}

func NewJobInstanceRepository(ctx *db_context.DbContext) *JobInstanceRepository {
	return &JobInstanceRepository{
		dbContext: ctx,
	}
}

func (r *JobInstanceRepository) FindPendingByJobId(uuid uuid.UUID) (*entity.JobInstance, error) {
	var jobInstance entity.JobInstance
	result := r.dbContext.GormDb.
		Where("job_id = ?", uuid).
		Where("status = ?", entity.Pending).
		First(&jobInstance)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &jobInstance, nil
}
