package entity

import "github.com/gofrs/uuid"

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

type JobInstanceLog struct {
	Base
	JobInstanceID uuid.UUID   `gorm:"type:varchar(36);index"`
	JobInstance   JobInstance `gorm:"foreignKey:JobInstanceID"`
	Message       string
	Level         LogLevel `gorm:"type:enum('debug','info','warn','error');default:'info'"`
}

func (JobInstanceLog) TableName() string {
	return "aur_job_instance_log"
}
