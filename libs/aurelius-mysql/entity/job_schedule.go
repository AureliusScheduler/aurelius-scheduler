package entity

import "github.com/gofrs/uuid"

type JobSchedule struct {
	Base
	ServiceID  uuid.UUID `gorm:"type:varchar(36);index"`
	Service    Service   `gorm:"foreignKey:ServiceID"`
	JobName    string
	JobVersion string // nullable
	Cron       string
}

func (JobSchedule) TableName() string {
	return "aur_job_schedule"
}
