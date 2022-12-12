package invoiceitem

import "time"

type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
}

//Service of invoiceitem
type Service struct {
	storage Storage
}

//NewService return a service pointer
func NewService(s Storage) *Service {
	return &Service{s}
}

//Migrate is used to migrate invoiceitem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
