package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Agent struct {
	Base
	ServiceID    uuid.UUID
	Service      Service `gorm:"foreignKey:ServiceID"`
	AgentStack   string
	AgentVersion string
	RegisteredAt time.Time
	LastSeenAt   time.Time
}
