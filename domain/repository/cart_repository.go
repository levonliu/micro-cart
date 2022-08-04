package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	CreateCart(*model.Cart) (int64, error)
	FindCartByID(int64) (*model.Cart, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{
		mysqlDb: db,
	}
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) (id int64, err error) {
	db := c.mysqlDb.FirstOrCreate(cart, &model.Cart{
		ProductID: cart.ProductID,
		UserID:    cart.UserID,
		SizeID:    cart.SizeID,
	})
	if db.Error != nil {
		return 0, db.Error
	}

	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}

	return cart.ID, nil
}

func (c *CartRepository) FindCartByID(id int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, c.mysqlDb.First(cart, id).Error
}

func (c *CartRepository) DeleteCart(id int64) error {
	return c.mysqlDb.Where("id = ?", id).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

func (c *CartRepository) FindAllCart(userId int64) (carts []model.Cart, err error) {
	return carts, c.mysqlDb.Where("user_id = ?", userId).Find(&carts).Error
}

func (c *CartRepository) CleanCart(userId int64) error {
	return c.mysqlDb.Where("user_id = ?", userId).Delete(&model.Cart{}).Error

}

func (c *CartRepository) IncrNum(id int64, changeNum int64) error {
	cart := &model.Cart{ID: id}
	return c.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", changeNum)).Error
}

func (c *CartRepository) DecrNum(id int64, changeNum int64) error {
	cart := &model.Cart{ID: id}
	db := c.mysqlDb.Model(cart).Where("num >= ?", changeNum).UpdateColumn("num", gorm.Expr("num - ?", changeNum))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}

	return nil
}
