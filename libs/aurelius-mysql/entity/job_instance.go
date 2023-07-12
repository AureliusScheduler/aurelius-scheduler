package entity

import "github.com/gofrs/uuid"

type JobInstanceStatus string

const (
	Pending JobInstanceStatus = "pending"
	Running JobInstanceStatus = "running"
	Success JobInstanceStatus = "success"
	Failure JobInstanceStatus = "failure"
)

type JobInstance struct {
	Base
	ServiceID  uuid.UUID `gorm:"type:uuid;index"`
	Service    Service   `gorm:"foreignKey:ServiceID"`
	AgentJobID uuid.UUID `gorm:"type:uuid;index"`
	AgentJob   AgentJob  `gorm:"foreignKey:AgentJobID"`
	AgentStack string
	Status     JobInstanceStatus `gorm:"type:enum('pending','running','success','failure');default:'pending'"`
}
