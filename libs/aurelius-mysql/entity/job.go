package entity

import (
	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type Job struct {
	Base
	AgentID           uuid.UUID `gorm:"type:varchar(36);index"`
	Agent             *Agent    `gorm:"foreignKey:AgentID"`
	AgentRegistration int
	Name              string
	Version           string
	StartedAt         null.Time
	FinishedAt        null.Time
}

func (Job) TableName() string {
	return "aur_job"
}
