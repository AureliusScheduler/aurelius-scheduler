package aurelius_service

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type JobInstanceService struct {
	dbContext   *db_context.DbContext
	nowProvider NowProvider
}

func NewJobInstanceService(
	ctx *db_context.DbContext,
	nowProvider NowProvider,
) *JobInstanceService {
	return &JobInstanceService{
		dbContext:   ctx,
		nowProvider: nowProvider,
	}
}

func (s *JobInstanceService) FindAndLockNext() ([]*entity.JobInstance, error) {
	items := make([]*entity.JobInstance, 0)

	err := s.dbContext.GormDb.Transaction(func(tx *gorm.DB) error {
		result := tx.
			Where("status = ?", entity.Pending).
			Where("run_at <= ?", s.nowProvider.Now()).
			Where("locked_at IS NULL").
			Find(&items)

		if result.Error != nil {
			return result.Error
		}

		if len(items) == 0 {
			return nil
		}

		ids := make([]uuid.UUID, 0)
		for _, item := range items {
			ids = append(ids, item.ID)
		}

		err := tx.Model(&entity.JobInstance{}).Where("id IN ?", ids).Updates(map[string]interface{}{
			"status":    entity.Pending,
			"locked_at": s.nowProvider.Now(),
		}).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}
