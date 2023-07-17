package seeder

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/entity"
	"github.com/go-faker/faker/v4"
	"strconv"
	"time"
)

type Seeder struct {
	Agents []entity.Agent
}

func NewSeeder() *Seeder {
	return &Seeder{}
}

func (s *Seeder) Seed(ctx db_context.DbContext) {
	serviceNames := []string{"php-service", "nodejs-service", "go-service"}

	// create 2-5 agents for each service name
	for i := 0; i < len(serviceNames); i++ {
		num, _ := faker.RandomInt(2, 5, 1)

		for j := 0; j < num[0]; j++ {
			s.Agents = append(s.Agents, entity.Agent{
				ServiceName:  serviceNames[i],
				AgentStack:   faker.Word(),
				AgentVersion: faker.Word(),
				RegisteredAt: getRandomDateBefore(15),
				LastSeenAt:   getRandomDateBefore(10),
			})
		}
	}

	// for each agent create 1-5 jobs
	for i := 0; i < len(s.Agents); i++ {
		num, _ := faker.RandomInt(1, 5, 1)
		version, _ := faker.RandomInt(1, 5, 1)

		for k := 0; k < num[0]; k++ {
			s.Agents[i].Jobs = append(s.Agents[i].Jobs, entity.Job{
				JobName:    faker.Word(),
				JobVersion: strconv.Itoa(version[0]),
				StartedAt:  getRandomDateBefore(15),
				FinishedAt: getRandomDateBefore(10),
			})
		}
	}

	// get a unique list of job names concatenated with job versions for each job
	var jobNames []string
	for j := 0; j < len(s.Agents); j++ {
		for k := 0; k < len(s.Agents[j].Jobs); k++ {
			jobNames = append(jobNames, s.Agents[j].Jobs[k].JobName+"___"+s.Agents[j].Jobs[k].JobVersion)
		}
	}

	result := ctx.GormDb.Create(s.Agents)

	if result.Error != nil {
		panic(result.Error)
	}
}

func getRandomDateBefore(days int32) time.Time {
	num, _ := faker.RandomInt(0, int(days), 1)
	return time.Now().AddDate(0, 0, -int(num[0]))
}
