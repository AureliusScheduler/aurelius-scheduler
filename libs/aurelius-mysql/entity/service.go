package entity

type Service struct {
	Base
	Name string
}

func (Service) TableName() string {
	return "aur_service"
}
