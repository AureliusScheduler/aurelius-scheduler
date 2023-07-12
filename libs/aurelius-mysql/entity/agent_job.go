package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type AgentJob struct {
	Base
	AgentID    uuid.UUID `gorm:"type:uuid;index"`
	JobName    string
	JobVersion string
	StartedAt  time.Time
	FinishedAt time.Time
}
