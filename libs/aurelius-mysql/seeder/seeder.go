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

func (s *Seeder) Seed(ctx *db_context.DbContext) {
	serviceNames := []string{"php-service", "nodejs-service", "go-service"}
	versions := []string{"1.0.0", "1.1.0"}

	// create 2-5 agents for each service name
	for i := 0; i < len(serviceNames); i++ {
		num, _ := faker.RandomInt(2, 5, 1)

		for j := 0; j < num[0]; j++ {
			ver, _ := faker.RandomInt(0, len(versions)-1, 1)
			s.Agents = append(s.Agents, entity.Agent{
				Name:         faker.Word(),
				Stack:        faker.Word(),
				Version:      versions[ver[0]],
				Registration: 1,
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
				Name:    faker.Word(),
				Version: strconv.Itoa(version[0]),
				//StartedAt:  null.TimeFrom(getRandomDateBefore(15)),
				//FinishedAt: null.TimeFrom(getRandomDateBefore(10)),
			})
		}
	}

	// get a unique list of job names concatenated with job versions for each job
	var jobNames []string
	for j := 0; j < len(s.Agents); j++ {
		for k := 0; k < len(s.Agents[j].Jobs); k++ {
			jobNames = append(jobNames, s.Agents[j].Jobs[k].Name+"___"+s.Agents[j].Jobs[k].Version)
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
