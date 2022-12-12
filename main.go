package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/sebasromero/go_db_postgre/pkg/product"
	"github.com/sebasromero/go_db_postgre/storage"
)

func main() {
	//Conection to the database
	storage.NewPostgresDB()
	productStorage := storage.NewPsqlProduct(storage.Pool())
	productService := product.NewService(productStorage)

	//Example to how to give format to a row
	/*
		m := &product.Model{
			Name:  "Go course",
			Price: 50,
		}
	*/

	//To get all the rowss
	/* 	ms, err := productService.GetAll()
	   	if err != nil {
	   		log.Fatalf("product.GetAll: %v", err)
	   	}
	   	fmt.Println(ms) */

	//Get the exact row by id
	m, err := productService.GetByID(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("Row couldn't be found, check the ID")
	case err != nil:
		log.Fatalf("product.GetBydID: %v", err)
	default:
		fmt.Println(m)
	}

	/* 	With this, the database inicialization is do it
	with the schemas
	storage.NewPostgresDB()

	   	productStorage := storage.NewPsqlProduct(storage.Pool())
	   	productService := product.NewService(productStorage)
	   	if err := productService.Migrate(); err != nil {
	   		log.Fatalf("product.Migrate: %v", err)
	   	}

	   	invoiceHeaderStorage := storage.NewPsqlInvoiceHeader(storage.Pool())
	   	invoiceHeaderService := invoiceheader.NewService(invoiceHeaderStorage)

	   	if err := invoiceHeaderService.Migrate(); err != nil {
	   		log.Fatalf("invoiceHeader.Migrate: %v", err)
	   	}

	   	invoiceItemStorage := storage.NewPsqlInvoiceItem(storage.Pool())
	   	invoiceItemService := invoiceheader.NewService(invoiceItemStorage)

	   	if err := invoiceItemService.Migrate(); err != nil {
	   		log.Fatalf("invoiceItem.Migrate: %v", err)
	   	} */
}
