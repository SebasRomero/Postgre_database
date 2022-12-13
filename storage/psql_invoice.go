package storage

import (
	"database/sql"
	"fmt"

	"github.com/sebasromero/go_db_postgre/pkg/invoice"
	"github.com/sebasromero/go_db_postgre/pkg/invoiceheader"
	"github.com/sebasromero/go_db_postgre/pkg/invoiceitem"
)

//PsqlInvoice is used to work with postgres - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

//NewPsqlInvoice returns a new psqlInvoice pointer
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

//Create implements the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}
	fmt.Printf("Invoice created with ID: %d\n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}
	fmt.Printf("Items created: %d\n", len(m.Items))
	return tx.Commit()
}
