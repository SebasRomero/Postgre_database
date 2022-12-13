package storage

import (
	"database/sql"
	"fmt"

	"github.com/sebasromero/go_db_postgre/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP, 
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products(name, 
		observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct  = `SELECT * FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

//PsqlProduct used to work with Postgre - Product
type PsqlProduct struct {
	db *sql.DB
}

//NewPsqlProduct returns a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db: db}
}

//Migrate implements the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	statement, err := p.db.Prepare(psqlMigrateProduct)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration created succesfully")
	return nil
}

//Create implements the interface product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	statement, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer statement.Close()

	err = statement.QueryRow(
		m.Name,
		stringtoNull(m.Observations),
		m.Price,
		m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("Row created succesfully")
	return nil
}

//Update implements the interface product.Storage
func (p *PsqlProduct) Update(m *product.Model) error {
	statement, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer statement.Close()
	res, err := statement.Exec(
		m.Name,
		stringtoNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("The product with the ID %d doesn't exist", m.ID)
	}

	fmt.Println("The row was updated succesfully")
	return nil
}

//Delete implements the interface product.Storage
func (p *PsqlProduct) Delete(id uint) error {
	statement, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer statement.Close()

	res, err := statement.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("The product with the ID %d doesn't exist", id)
	}
	fmt.Println("The row was deleted succesfully")
	return nil
}

//GetAll implements the interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	statement, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

//GetById implements the interface product.Storage
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	statement, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer statement.Close()
	return scanRowProduct(statement.QueryRow(id))
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
