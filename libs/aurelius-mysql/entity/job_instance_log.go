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
	JobInstanceID uuid.UUID   `gorm:"type:uuid;index"`
	JobInstance   JobInstance `gorm:"foreignKey:JobInstanceID"`
	Message       string
	Level         LogLevel `gorm:"type:enum('debug','info','warn','error');default:'info'"`
}
