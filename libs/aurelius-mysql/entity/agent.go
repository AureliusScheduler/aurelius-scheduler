package entity

import (
	"time"
)

type Agent struct {
	Base
	Name         string
	Instance     string
	Stack        string
	Version      string
	Jobs         []Job
	Registration int `gorm:"type:int(1);default:1"`
	RegisteredAt time.Time
	LastSeenAt   time.Time
}

func (Agent) TableName() string {
	return "aur_agent"
}
