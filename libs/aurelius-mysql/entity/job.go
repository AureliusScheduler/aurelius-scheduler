package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Job struct {
	Base
	AgentID    uuid.UUID `gorm:"type:varchar(36);index"`
	Agent      Agent     `gorm:"foreignKey:AgentID"`
	JobName    string
	JobVersion string
	StartedAt  time.Time
	FinishedAt time.Time
}

func (Job) TableName() string {
	return "aur_job"
}
