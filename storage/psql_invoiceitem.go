package storage

import (
	"database/sql"
	"fmt"

	"github.com/sebasromero/go_db_postgre/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP, 
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
		(invoice_header_id) REFERENCES invoice_headers(id) ON UPDATE 
		RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY
		(product_id) REFERENCES products (id) ON UPDATE 
		RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id,
		product_id) VALUES($1, $2) RETURNING id, created_at`
)

//PsqlInvoiceItem used to work with Postgre - invoiceItem
type PsqlInvoiceItem struct {
	db *sql.DB
}

//NewPsqlInvoiceItem returns a new pointer of PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db: db}
}

//Migrate implement the interface invoiceItem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	statement, err := p.db.Prepare(psqlMigrateInvoiceItem)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	fmt.Println("InvoiceItem migration created succesfully")
	return nil
}

//CreateTx implements the interface invoiceitem.Storage
func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint,
	ms invoiceitem.Models) error {
	statement, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer statement.Close()

	for _, item := range ms {
		err = statement.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
