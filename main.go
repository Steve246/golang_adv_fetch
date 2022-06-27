package main

import (
	"enigmacamp.com/go-db-fundamnetal/config"
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			utils.IsError(err)
		}
	}(db)
	// cstRepo := repository.NewCustomerRepository(db)
	// cstUse := usecase.NewCustomerUseCase(cstRepo)

	// INSERT
	// cstId := utils.GenerateId()
	// customer := model.Customer{
	// 	Id:      cstId,
	// 	Name:    "Jution Kirana",
	// 	Address: "Ragunan",
	// 	Phone:   "08292929",
	// 	Email:   "jutionck@gmail.com",
	// 	Balance: 150000,
	// }
	// cstUse.InsertCustomer(&customer)

	//INSERT PAKE SHOP

	cstRepoShop := repository.NewShopRepository(db)
	// cstUseShop := usecase.NewShopUseCase(cstRepoShop)

	// cstIdShop := utils.GenerateId()
	shop := model.Shop{
		Id:      "6efd1a98-e554-497f-8ad8-f2086004f70b",
		No_siup: "57858558",
		Name:    "Diganti Pen Pen",
		Address: "Diganti Bandung",
		Phone:   "0865544344",
	}
	// cstUseShop.InsertShop(&shop)

	// DELETE
	// customerId := "C004"
	// err := cstUse.DeleteCustomer(customerId)
	// if err != nil {
	// 	fmt.Println("error test")
	// 	fmt.Println(err.Error())
	// }

	//DELETE PAKE SHOP

	// shopperId := "b4b7d8f2-b0f0-43ff-b425-8af101dec20c"
	// err := cstUseShop.DeleteShop(shopperId)
	// if err != nil {
	// 	fmt.Println("error test")
	// 	fmt.Println(err.Error())
	// }

	// UPDATE
	// customerUpdate := model.Customer{
	// 	Name:    "Jution Aja",
	// 	Address: "Ragunan",
	// 	Phone:   "08292929",
	// 	Email:   "jutionck@gmail.com",
	// 	Id:      "0fa70a67-52b4-4e17-9e6d-e6733b2c127b",
	// }
	// cstRepo.Update(&customerUpdate)

	//UPDATE PAKE SHOP

	cstRepoShop.Update(&shop)
}
