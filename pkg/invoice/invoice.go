package invoice

import (
	"github.com/sebasromero/go_db_postgre/pkg/invoiceheader"
	"github.com/sebasromero/go_db_postgre/pkg/invoiceitem"
)

// Invoice Model
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage is an interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Invoice Service
type Service struct {
	storage Storage
}

//NewService return a service pointer
func NewService(s Storage) *Service {
	return &Service{s}
}

//Create a new invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
