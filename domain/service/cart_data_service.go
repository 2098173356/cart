package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

//创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartDataService repository.ICartRepository
}

func (c CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartDataService.CreateCart(cart)
}

func (c CartDataService) DeleteCart(i int64) error {
	return c.CartDataService.DeleteCart(i)
}

func (c CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartDataService.UpdateCart(cart)
}

func (c CartDataService) FindCartByID(i int64) (*model.Cart, error) {
	return c.CartDataService.FindCartByID(i)
}

func (c CartDataService) FindAllCart(i int64) ([]model.Cart, error) {
	return c.CartDataService.FindAll(i)
}

func (c CartDataService) CleanCart(i int64) error {
	return c.CartDataService.CleanCart(i)
}

func (c CartDataService) DecrNum(i int64, i2 int64) error {
	return c.CartDataService.DecrNum(i, i2)
}

func (c CartDataService) IncrNum(i int64, i2 int64) error {
	return c.CartDataService.IncrNum(i, i2)
}
