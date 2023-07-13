package entity

type Service struct {
	Base
	Name   string
	Agents []Agent
}

func (Service) TableName() string {
	return "aur_service"
}

//func (s *Service) BeforeCreate() (err error) {
//	s.ID = NewID()
//	return
//}
