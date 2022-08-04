package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	FindCartByID(int64) (*model.Cart, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: cartRepository}
}

func (c *CartDataService) AddCart(cart *model.Cart) (id int64, err error) {
	return c.CartRepository.CreateCart(cart)
}

func (c *CartDataService) FindCartByID(id int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(id)
}

func (c *CartDataService) DeleteCart(id int64) error {
	return c.CartRepository.DeleteCart(id)
}

func (c *CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}

func (c CartDataService) FindAllCart(userId int64) ([]model.Cart, error) {
	return c.CartRepository.FindAllCart(userId)
}

func (c *CartDataService) CleanCart(userId int64) error {
	return c.CartRepository.CleanCart(userId)
}

func (c *CartDataService) IncrNum(id int64, num int64) error {
	return c.CartRepository.IncrNum(id, num)
}

func (c *CartDataService) DecrNum(id int64, num int64) error {
	return c.CartRepository.IncrNum(id, num)
}
