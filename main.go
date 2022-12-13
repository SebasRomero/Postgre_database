package main

import (
	"github.com/sebasromero/go_db_postgre/storage"
)

func main() {
	//Conection to the database
	storage.NewPostgresDB()

	//Invoice creation
	/*
	   storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	   	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	   	storageInvoice := storage.NewPsqlInvoice(
	   		storage.Pool(),
	   		storageHeader,
	   		storageItems,
	   	)

	   	m := &invoice.Model{
	   		Header: &invoiceheader.Model{
	   			Client: "Pacho",
	   		},
	   		Items: invoiceitem.Models{
	   			&invoiceitem.Model{ProductID: 4},
	   		},
	   	}

	   	invoiceService := invoice.NewService(storageInvoice)
	   	if err := invoiceService.Create(m); err != nil {
	   		log.Fatalf("invoice.Create: %v", err)
	   	}
	*/

	//Define the service
	/*
		productStorage := storage.NewPsqlProduct(storage.Pool())
		productService := product.NewService(productStorage)
	*/

	//Here we delete the item ID specified
	/*
		err := productService.Delete(3)
		if err != nil {
			log.Fatalf("product.Delete: %v", err)
		}
	*/

	//Here we update a row, creating the new model and putting in the Update method
	/*
	   m := &product.Model{
	   		ID:           1,
	   		Name:         "The best course of Go",
	   		Observations: "Updated with the new versions",
	   		Price:        99,
	   	}
	   	err := productService.Update(m)
	   	if err != nil {
	   		log.Fatalf("product.Update: %v", err)
	   	}
	*/

	//Example to how to create an item
	/*
	   m := &product.Model{
	   		Name:  "Java course",
	   		Price: 67,
	   	}
	   	if err := productService.Create(m); err != nil {
	   		log.Fatalf("product.Create: %v", err)
	   	}
	   	fmt.Printf("%+v\n", m)
	*/

	//To get all the rows
	/*
		ms, err := productService.GetAll()
		if err != nil {
			log.Fatalf("product.GetAll: %v", err)
		}
		fmt.Println(ms)
	*/

	//Get the exact row by id
	/*
		m, err := productService.GetByID(1)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			fmt.Println("Row couldn't be found, check the ID")
		case err != nil:
			log.Fatalf("product.GetBydID: %v", err)
		default:
			fmt.Println(m)
		}
	*/

	//With this, the database inicialization is created with the schemas
	/*
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
		   	}
	*/
}
