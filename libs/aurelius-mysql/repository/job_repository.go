package repository

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type JobRepository struct {
	dbContext *db_context.DbContext
}

// NewJobRepository creates a new JobRepository
func NewJobRepository(dbContext *db_context.DbContext) *JobRepository {
	return &JobRepository{dbContext: dbContext}
}

func (r *JobRepository) FindByUUID(uuid uuid.UUID) (*entity.Job, error) {
	var job entity.Job

	result := r.dbContext.GormDb.
		Where("id = ?", uuid).
		First(&job)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &job, nil
}
