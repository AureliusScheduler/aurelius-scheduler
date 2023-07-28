package service

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"gorm.io/gorm"
	"time"
)

type AgentService struct {
	dbContext *db_context.DbContext
}

type RegisterAgentParams struct {
	Name     string
	Instance string
	Stack    string
	Version  string
}

type RegisterJobParams struct {
	Name    string
	Version string
}

type RegisterAgentResult struct {
	AgentId string
}

func NewAgentService(dbContext *db_context.DbContext) *AgentService {
	return &AgentService{
		dbContext: dbContext,
	}
}

func (s *AgentService) RegisterAgent(params *RegisterAgentParams, jobs []*RegisterJobParams) (*RegisterAgentResult, error) {
	// find agent if exists
	// update jobs
	// update last seen at

	// if not exists
	// create agent
	// create jobs
	// set RegisteredAt
	// set LastSeenAt

	var agent entity.Agent
	isNew := false

	result := s.dbContext.GormDb.Where(&entity.Agent{Name: params.Name, Instance: params.Instance}).First(&agent)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			agent = entity.Agent{
				Name:         params.Name,
				Instance:     params.Instance,
				Stack:        params.Stack,
				Version:      params.Version,
				Registration: 1,
				RegisteredAt: time.Now(),
			}

			isNew = true
		} else {
			return nil, result.Error
		}
	} else {
		agent.Registration += 1
	}

	agent.LastSeenAt = time.Now()

	var jobEntities []entity.Job

	// register jobs
	for _, job := range jobs {
		var jobEntity entity.Job

		// if the agent has existed before try finding a job
		// otherwise create a new job

		if !isNew {
			result := s.dbContext.GormDb.Where(&entity.Job{
				Name:    job.Name,
				Version: job.Version,
				AgentID: agent.ID,
			}).First(&jobEntity)

			if result.Error != nil {
				if result.Error == gorm.ErrRecordNotFound {
					jobEntity = entity.Job{
						Name:    job.Name,
						Version: job.Version,
						AgentID: agent.ID,
					}
				} else {
					return nil, result.Error
				}
			}

		} else {
			jobEntity = entity.Job{
				Name:    job.Name,
				Version: job.Version,
				AgentID: agent.ID,
			}
		}

		jobEntity.AgentRegistration = agent.Registration
		jobEntities = append(jobEntities, jobEntity)
	}

	agent.Jobs = jobEntities

	err := s.dbContext.GormDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&agent).Error; err != nil {
			return err
		}

		if !isNew {
			if err := tx.Save(&jobEntities).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	opResult := RegisterAgentResult{
		AgentId: agent.ID.String(),
	}

	return &opResult, nil
}
