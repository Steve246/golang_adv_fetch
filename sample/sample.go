package sample

import (
	"fmt"
	"log"

	"enigmacamp.com/go-db-fundamnetal/dto"
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

	// customerId01 := "2d8212ae-78ef-47c7-835f-1bcbafdd493d"
	// customerId02 := "1711fb50-4fa0-4d8d-a923-dec7d10f1c40"

	// var customers model.Customer
	// customers, err := cstUse.GetById(customerId01)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customers)

	// customers2, err := cstUse.GetById(customerId02)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customers2)

	//Get Name

	// var customers []model.Customer
	// customers, err := cstUse.GetByName("Kunyil Unyil")

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for _, customer := range customers {
	// 	fmt.Println(customer)
	// }

	//Get count

	// var customers int
	// customers, err := cstUse.GetTotalCustomerActive()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println(customers)

	//Get sum

	var customers int
	customers, err := cstUse.GetTotalBalanceActive()

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(customers)

	//get kota - total

	// var customersCount []dto.CustomerCount
	// customersCount, _ = cstUse.GetKotaTotalByAddress()

	// for _, row := range customersCount {
	// 	fmt.Println(row)
	// }

}

// func ProductRun(db *sqlx.DB) {
// 	repo := repository.GetShopProduct(db)
// 	useCase := usecase.FindShopWithProduct(repo)

// 	product := model.Product{
// 		Id:          utils.GenerateId(),
// 		Name:        "Mouse",
// 		Price:       200000,
// 		Description: "Mouse Benda",
// 		Stock:       20,
// 		Shop: model.Shop{
// 			Id: "7dc92174-82ad-41a0-98d8-44aa45552x13",
// 		},
// 	}

// 	err := useCase.InsertNewProduct(&product)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// }

func ShopRun(db *sqlx.DB) {
	repo := repository.NewShopRepository(db)
	useCase := usecase.NewShopUseCase(repo)

	var shopWithproducts []dto.ShopProductDto
	shopWithproducts, err := useCase.FindShopWithProduct(1, 10)
	if err != nil {
		log.Println(err.Error())
	}

	for _, sp := range shopWithproducts {
		fmt.Printf("%#v\n", sp)
	}
}
