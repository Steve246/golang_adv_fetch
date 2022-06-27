package usecase

import (
	"enigmacamp.com/go-db-fundamnetal/model"
	"enigmacamp.com/go-db-fundamnetal/repository"
)

type ShopUseCase interface {
	InsertShop(customer *model.Shop) error
	UpdateShop(customer *model.Shop) error
	DeleteShop(id string) error
}

type shopUseCase struct {
	repo repository.ShopRepository
}

func (c *shopUseCase) InsertShop(shop *model.Shop) error {
	return c.repo.Insert(shop)
}

func (c *shopUseCase) UpdateShop(shop *model.Shop) error {
	return c.repo.Update(shop)
}

func (c *shopUseCase) DeleteShop(id string) error {
	return c.repo.Delete(id)
}

func NewShopUseCase(repo repository.ShopRepository) ShopUseCase {
	usc := new(shopUseCase)
	usc.repo = repo
	return usc
}
