package main

import (
	"enigmacamp.com/go-db-fundamnetal/config"
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

	// cstRepoShop := repository.NewShopRepository(db)
	// cstUseShop := usecase.NewShopUseCase(cstRepoShop)

	// cstIdShop := utils.GenerateId()
	// shop := model.Shop{
	// 	Id:      cstIdShop,
	// 	No_siup: "57858558",
	// 	Name:    "Diganti Pen Pen",
	// 	Address: "Diganti Bandung",
	// 	Phone:   "0865544344",
	// }
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

	// cstRepoShop.Update(&shop)

	//Panggil Get

	// sample.CustomerRun(db)

	//tambain transaction

	// tx := db.MustBegin()
	// cstId := 10001
	// cstId2 := 10002
	// tx.MustExec(`insert into customer (id, name, saldo) values ($1, $2, $3)`, cstId, "Bulan", 5000)

	// tx.MustExec(`insert into customer (id, name, saldo) values ($1, $2, $3)`, cstId2, "Bintang", 5000)

	// rslt := tx.MustExec(`update customer set saldo=(saldo+200) where id=$1`, cstId)

	// r, _ := rslt.RowsAffected()

	// if r == 0 {
	// 	tx.Rollback()
	// }

	// rslt2 := tx.MustExec(`update customer set saldo=(saldo-200) where id=$1`, cstId2)

	// r2, _ := rslt2.RowsAffected()

	// if r2 == 0 {
	// 	tx.Rollback()
	// }

	// tx.Commit()

	//beli roti bakar, dikurang dari harganya untuk pembeli dan kurangin quantitynya

	tx := db.MustBegin()
	cstId := 10001

	rslt := tx.MustExec(`update customer set saldo=(saldo-10000) where id=$1`, cstId)

	r, _ := rslt.RowsAffected()

	if r == 0 {
		tx.Rollback()
	}

	productId := "E0001"

	pslt := tx.MustExec(`update product set stock=(stock-1) where id=$1`, productId)

	p, _ := pslt.RowsAffected()

	if p == 0 {
		tx.Rollback()
	}

	//masukin ke dalam transaksi

	tslt := tx.MustExec(`insert into transaction (id, customer_id, product_id, purchase_date, quantity) values ($1, $2, $3, $4, $5)`, 001, cstId, productId, 2022-06-28, 1)

	pt, _ := tslt.RowsAffected()

	if pt == 0 {
		tx.Rollback()
	}

	tx.Commit()

}
