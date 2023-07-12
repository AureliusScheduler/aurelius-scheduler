package entity

import "github.com/gofrs/uuid"

type JobSchedule struct {
	Base
	ServiceID  uuid.UUID `gorm:"type:uuid;index"`
	Service    Service   `gorm:"foreignKey:ServiceID"`
	JobName    string
	JobVersion string // nullable
	Cron       string
}
