package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Agent struct {
	Base
	ServiceID    uuid.UUID `gorm:"type:varchar(36);index"`
	Service      Service   `gorm:"foreignKey:ServiceID"`
	AgentStack   string
	AgentVersion string
	Jobs         []Job
	RegisteredAt time.Time
	LastSeenAt   time.Time
}

func (Agent) TableName() string {
	return "aur_agent"
}
