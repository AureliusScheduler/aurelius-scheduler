package seeder

import (
	"aurelius/libs/aurelius-mysql/entity"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Seeder struct {
	Services []entity.Service
}

func NewSeeder() *Seeder {
	return &Seeder{}
}

func (s *Seeder) Seed(db *gorm.DB) {

	// create 5 services with faker
	for i := 0; i < 5; i++ {
		s.Services = append(s.Services, entity.Service{
			Name: faker.Name(),
		})
	}

	// create 2-5 agents for each service
	for i := 0; i < len(s.Services); i++ {
		num, _ := faker.RandomInt(2, 5, 1)
		for j := 0; j < num[0]; j++ {
			s.Services[i].Agents = append(s.Services[i].Agents, entity.Agent{
				AgentStack:   faker.Word(),
				AgentVersion: faker.Word(),
				RegisteredAt: getRandomDateBefore(30),
				LastSeenAt:   getRandomDateBefore(15),
			})
		}
	}

	// for each agent create 1-5 jobs
	for i := 0; i < len(s.Services); i++ {
		for j := 0; j < len(s.Services[i].Agents); j++ {
			num, _ := faker.RandomInt(1, 5, 1)
			version, _ := faker.RandomInt(1, 5, 1)

			for k := 0; k < num[0]; k++ {
				s.Services[i].Agents[j].Jobs = append(s.Services[i].Agents[j].Jobs, entity.Job{
					JobName:    faker.Word(),
					JobVersion: strconv.Itoa(version[0]),
					StartedAt:  getRandomDateBefore(15),
					FinishedAt: getRandomDateBefore(10),
				})
			}
		}
	}

	// get a unique list of job names concatenated with job versions for each job
	var jobNames []string
	for i := 0; i < len(s.Services); i++ {
		for j := 0; j < len(s.Services[i].Agents); j++ {
			for k := 0; k < len(s.Services[i].Agents[j].Jobs); k++ {
				jobNames = append(jobNames, s.Services[i].Agents[j].Jobs[k].JobName+"___"+s.Services[i].Agents[j].Jobs[k].JobVersion)
			}
		}
	}

	result := db.Create(s.Services)

	if result.Error != nil {
		panic(result.Error)
	}
}

func getRandomDateBefore(days int32) time.Time {
	num, _ := faker.RandomInt(0, int(days), 1)
	return time.Now().AddDate(0, 0, -int(num[0]))
}
