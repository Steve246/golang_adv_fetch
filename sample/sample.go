package sample

import (
	"fmt"
	"log"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
	"enigmacamp.com/go-db-fundamnetal/usecase"
	"github.com/jmoiron/sqlx"
)

func CustomerRun(db *sqlx.DB) {

	cstrepo := repository.NewCustomerRepository(db)

	cstUse := usecase.NewCustomerUseCase(cstrepo)

	//Get All

	// var customers []model.Customer
	// customers, err := cstUse.GetAll(2, 2)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for _, customer := range customers {
	// 	fmt.Println(customer)

	//Get ID

	// var customers model.Customer
	// customers, err := cstUse.GetById("2611fb50-4fa0-4d8d-a923-dec7d10f1c40")

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customers)

	//Get Name

	var customers []model.Customer
	customers, err := cstUse.GetByName("Kunyil Unyil")

	if err != nil {
		log.Println(err.Error())
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}

}
