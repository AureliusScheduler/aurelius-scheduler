package service

import (
	"aurelius/apps/aurelius-rest/dto"
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"time"
)

type AgentService struct {
	dbContext *db_context.DbContext
}

func NewAgentService(ctx db_context.DbContext) *AgentService {
	return &AgentService{
		dbContext: &ctx,
	}
}

func (s *AgentService) GetAll() ([]dto.AgentDto, error) {
	var agents []entity.Agent
	result := s.dbContext.GormDb.Debug().Find(&agents)

	if result.Error != nil {
		return nil, result.Error
	}

	var agentDtos []dto.AgentDto

	for _, agent := range agents {
		agentDtos = append(agentDtos, dto.AgentDto{
			ID:           agent.ID.String(),
			Service:      agent.ServiceName,
			AgentStack:   agent.AgentStack,
			AgentVersion: agent.AgentVersion,
			RegisteredAt: agent.RegisteredAt.Format(time.RFC3339),
			LastSeenAt:   agent.LastSeenAt.Format(time.RFC3339),
		})
	}

	return agentDtos, nil
}
