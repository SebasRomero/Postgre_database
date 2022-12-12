package invoiceheader

import "time"

//Inoiceheader model
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Storage interface {
	Migrate() error
}

//Service of invoiceheader
type Service struct {
	storage Storage
}

//NewService return a service pointer
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is used to migrate invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
