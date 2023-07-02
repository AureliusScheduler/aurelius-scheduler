package repository

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-storage"
	"context"
	"github.com/doug-martin/goqu/v9"
)

type MySQLJobRepository struct {
	dbContext db_context.DbContext
	dialect   *goqu.DialectWrapper
}

// NewJobRepository creates a new JobRepository
func NewMySQLJobRepository(dbContext db_context.DbContext) storage.JobRepository {
	dialect := goqu.Dialect("mysql")
	return &MySQLJobRepository{dbContext: dbContext, dialect: &dialect}
}

// GetAll Get all jobs
func (r *MySQLJobRepository) GetAll() ([]storage.Job, error) {
	var jobs []storage.Job

	ds := r.dialect.From("jobs").Select("id", "name", "cron_expression")
	q, _, _ := ds.ToSQL()

	rows, err := r.dbContext.Db.QueryContext(context.Background(), q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var job storage.Job
		err := rows.Scan(&job.ID, &job.Name, &job.CronExpression)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

// Create a new job
func (r *MySQLJobRepository) Create(job *storage.Job) error {
	ds := r.dialect.Insert("jobs").Rows(goqu.Record{"id": job.ID, "name": job.Name, "cron_expression": job.CronExpression})
	q, _, _ := ds.ToSQL()
	_, err := r.dbContext.Db.ExecContext(context.Background(), q)
	return err
}
