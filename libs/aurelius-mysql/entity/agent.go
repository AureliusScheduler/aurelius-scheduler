package entity

import (
	"time"
)

type Agent struct {
	Base
	ServiceName  string
	AgentStack   string
	AgentVersion string
	Jobs         []Job
	RegisteredAt time.Time
	LastSeenAt   time.Time
}

func (Agent) TableName() string {
	return "aur_agent"
}
