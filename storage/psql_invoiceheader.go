package storage

import (
	"database/sql"
	"fmt"

	"github.com/sebasromero/go_db_postgre/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP, 
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES
	($1) RETURNING id, created_at`
)

//PsqlInvoiceHeader used to work with Postgre - invoiceHeader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

//NewPsqlInvoiceHeader returns a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db: db}
}

//Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	statement, err := p.db.Prepare(psqlMigrateInvoiceHeader)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	fmt.Println("InvoiceHeader migration created succesfully")
	return nil
}

//CreateTx implements the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	statement, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer statement.Close()

	return statement.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
