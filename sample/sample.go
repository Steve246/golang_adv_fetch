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

	var customers []model.Customer
	customers, err := cstUse.GetAll(2, 2)

	if err != nil {
		log.Println(err.Error())
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}

}
