package entity

import (
	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
	"time"
)

type JobInstanceStatus string

const (
	Pending JobInstanceStatus = "pending"
	Running JobInstanceStatus = "running"
	Success JobInstanceStatus = "success"
	Failure JobInstanceStatus = "failure"
)

type JobInstance struct {
	Base
	JobID      uuid.UUID `gorm:"type:varchar(36);index"`
	JobVersion string
	Cron       string
	RunAt      time.Time
	Status     JobInstanceStatus `gorm:"type:enum('pending','running','success','failure');default:'pending'"`

	LockedAt null.Time

	Job Job `gorm:"foreignKey:JobID"`
}

func (JobInstance) TableName() string {
	return "aur_job_instance"
}
