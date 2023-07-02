package storage

type Job struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	CronExpression string `db:"cron_expression"`
}

type JobRepository interface {
	GetAll() ([]Job, error)
	Create(job *Job) error
}
