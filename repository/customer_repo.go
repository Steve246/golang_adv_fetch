package repository

import (
	"errors"

	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/utils"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id string) error
	GetAll(page int, totalRow int) ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) Insert(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *model.Customer) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_CUSTOMER, id)
	return
}

//bagian select

func (c *customerRepository) GetAll(page int, totalRow int) ([]model.Customer, error) {
	//rumus pagination
	limit := totalRow
	offset := limit * (page - 1) //batesin per page ad brp

	//pengunaan fetch data: select($1), named querry {: column_name}, querry{}
	customer := []model.Customer{}

	if err := c.db.Select(&customer, utils.SELECT_ALL_CUSTOMER, limit, offset); err != nil {
		return nil, err
	}

	return customer, nil

}

func (c *customerRepository) GetById(id string) (model.Customer, error) {
	// panic("implement me")

	customer := []model.Customer{}

	if err := c.db.Select(&customer, utils.SELECT_CUSTOMER_BY_ID, id); err != nil {

		return customer[0], err

	}

	return customer[0], nil
}

func (c *customerRepository) GetByName(name string) ([]model.Customer, error) {
	// panic("implement me")

	customer := []model.Customer{}

	if err := c.db.Select(&customer, utils.SELECT_CUSTOMER_BY_NAME, "%"+name+"%"); err != nil {
		return nil, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
