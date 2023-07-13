package entity

import (
	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primarykey;default:(uuid())"`
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID, err = uuid.NewV4()

	if err != nil {
		return err
	}

	return
}
