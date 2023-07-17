package entity

type JobSchedule struct {
	Base
	JobName    string
	JobVersion string // nullable
	Cron       string
}

func (JobSchedule) TableName() string {
	return "aur_job_schedule"
}
