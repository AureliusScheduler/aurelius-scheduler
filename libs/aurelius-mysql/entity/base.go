package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (base *Base) BeforeCreate() (err error) {
	base.ID, err = uuid.NewV4()

	if err != nil {
		return err
	}

	return
}
